package editorialPersist

import (
	"hash/crc32"
	"encoding/hex"
	"visualchina/Model"
	"time"
	"strconv"
	"libary/query"
	"encoding/json"
	"libary/upload"
	"strings"
	"regexp"
)

type Editorial  struct{
	NavDb  chan Model.NavDb
	status chan bool
}

func (s *Editorial) NavSave(r Model.NavDb) Model.NavDb {
	url := hex.EncodeToString([]byte(r.GrabUrl))
	crcStr := crc32.ChecksumIEEE([]byte(url))
	r.Crc32  = crcStr
	r.AddDate = time.Now().Unix()
    row := r.GetNavDataByCrc32(crcStr)
	if len(row) >0 {
		id, _ := strconv.ParseInt(row["id"], 10, 64)
		r.Id = id
		r.UpdateNavExecTimeById(id,r.ExecDate)
		return r
	}
	id, _ := r.NavSave()
	r.Id = id
	return  r
}

func (s *Editorial) SaveTag(category query.TagModels) (bool) {
	   var insertId  int64
       cate := Model.CategoryDb{
           	Code:category.Code,
           	Type: category.Type,
           	CategoryId: category.Id,
           	CategoryName:category.Name,
		    CategoryPid:category.Pid,
           	CategoryUrl:category.Url,
           	TotalCount:0,
           	AddDate:time.Now().Unix(),
	   }
	   row := cate.GetTagDataByCateId()
	   if len(row) == 0 {
		   insertId, _ = cate.Save()
	   }
 	  for _,v:=range category.SubTags{
		  cate  = Model.CategoryDb{
			  Code:v.Code,
			  Type:category.Type,
			  CategoryId: v.Id,
			  CategoryPid: v.Pid,
			  CategoryName:v.Name,
			  Pid: strconv.FormatInt(insertId,10),
			  CategoryUrl:v.Url,
			  TotalCount:0,
			  AddDate:time.Now().Unix(),
		  }
		  row = cate.GetTagDataByCateId()
		  if len(row) == 0{
		       cate.Save()
		  }
	  }
	  return  true
}

func (s *Editorial) SaveGenera(genera query.Generalize) (bool) {
	gen := Model.GeneralizeDb{
            CategoryId:genera.CategoryId,
		    ImageId:genera.ImageId,
		    TopicId:genera.TopicId,
		    Gtype:genera.Gtype,
		    Title:genera.Title,
		    Src:genera.Src,
		    Link:genera.Link,
		    AddDate:time.Now().Unix(),
	}
	row := gen.GetListByCateIdANDLink()
	if len(row) == 0{
	     gen.Save()
	}
	return true
}

func (s *Editorial) SaveRecommend(recommend query.LevelRecommend) (bool)  {
	saveRecommend := Model.RecommendDb{
		GroupId:    recommend.GroupId,
		TopicId:    recommend.TopicId,
		CategoryId: recommend.CategoryId,
		AddDate:    time.Now().Unix(),
	}
	row := saveRecommend.GetRecommendDataById()
	if len(row) == 0{
    	 saveRecommend.Save()
	}
	return true
}

func (s *Editorial) SaveGroup(group List) (b bool) {
	loc, _ := time.LoadLocation("Local")
	parse, _ := time.ParseInLocation("2006-01-02 15:04:05", group.ImgDate,loc)
	groupId,_ := strconv.ParseInt(group.GroupId,10,64)
    groupData := Model.GroupDb{
           GroupId:groupId,
           OneCategory:group.OneCategory,
           Title:group.Title,
           Caption:group.Caption,
           FirstPicId:strconv.FormatInt(group.FirstResId,10),
           GroupPicsNum:group.GroupPicsNum,
           Keywords:group.Keywords,
           ImgDate:parse.Unix(),
           AddDate:time.Now().Unix(),
     }
     fields :=[]string{"group_id"}
	 ret := Model.GetGroupDataByGroupId(groupId, fields)
	 var id int64
	 if len(ret) == 0{
		 id, _= groupData.Save()
	 }
	 if id == 0{
	 	  return  false
	 }
	 GroupPics,_:= json.Marshal(group.GroupPics)
	 groupDetail := Model.GroupDetailDb{
	 	 GroupId:groupId,
		 EqualwUrl:group.EqualwUrl,
		 EqualwImageId:upload.UploadToQiniu(group.EqualwUrl),
		 EqualhUrl:group.EqualhUrl,
		 EqualhImageId:upload.UploadToQiniu(group.EqualhUrl),
		 Width:group.Width,
		 Height:group.Height,
		 Url800:group.Url800,
		 Url800ImageId:upload.UploadToQiniu(group.Url800),
		 GroupPics: string(GroupPics),
		 AddDate:time.Now().Unix(),
	 }
	r := Model.GetGroupDetailByGroupId(groupId, fields)
	if len(r) == 0{
		id, _= groupDetail.Save()
	}
	if id == 0{
		 return  false
	}
	return  true
}

//save visual_category_group
func (s *Editorial) SaveCategoryGroup(item List,mp map[string]bool)  {
	GroupId, _ := strconv.ParseInt(item.GroupId, 10, 64)
	CategoryIdList := strings.Split(item.Keywords,",")
	for _,CateId := range CategoryIdList{
		key := item.GroupId+"_"+CateId
		if _,ok := mp[key];ok{
			continue
		}
		CateId,_ := strconv.ParseInt(CateId,10,64)
		cateGroup := Model.CategoryGroupDb{
			CategoryId:CateId,
			GroupId:GroupId,
			Status:1,
			AddDate:time.Now().Unix(),
		}
		id, _ := cateGroup.Save()
		if id > 0{
			mp[key] = true
		}
	}
}



func (s *Editorial) SavePic(CategoryId int64,item Info) (id int64, err error){
    ImageDate := item.PicInfo.ImageDate
	loc, _ := time.LoadLocation("Local")
	imgTime,_ := time.ParseInLocation("2006-01-02 15:04:05", ImageDate, loc)
	GroupId,_ := strconv.ParseInt(item.GroupInfo.GroupId,10,64)
    var (
    	TopicIdList []string
    	RelateGroup []string
	)
	for _,topic := range  item.Topic{
		TopicIdList  = append(TopicIdList,strconv.FormatInt(topic.TopicId,10))
	}
	for _,group := range item.RelateGroup{
		RelateGroup = append(RelateGroup,group.GroupId)
	}
	pic := Model.PicDb{
		ResId:item.PicInfo.ResId,
		PicId:item.PicInfo.Id,
		ImgId:item.PicInfo.ImgId,
		TopicId: strings.Join(TopicIdList,","),
		GroupId: GroupId,
		RelateGroupId:strings.Join(RelateGroup,","),
		CategoryId:CategoryId,
		ProviderId:item.PicInfo.ProviderId,
		NextPicId:item.PicInfo.NextPicId,
		Type:item.PicInfo.Type,
		Url: item.PicInfo.Url,
		Title:item.PicInfo.Title,
		ImageDate:imgTime.Unix(),
	}
	return  pic.Save()
}

func (s *Editorial) SavePicDetail(item Info)  (id int64, err error) {
	meta := regexp.QuoteMeta(item.PicInfo.Caption)
	picDetail := Model.PicDetailDb{
		PicId:item.PicInfo.Id,
		Caption:meta,
		FileType:item.PicInfo.FileType,
		Size:item.PicInfo.Size,
		StoreSize:item.PicInfo.StoreSize,
		Specification:item.PicInfo.Specification,
		Cameraman:item.PicInfo.Cameraman,
		Brand:item.PicInfo.Brand,
		CopyRight:item.PicInfo.CopyRight,
		Category:item.PicInfo.Category,
	}
	return picDetail.Save()
}


//写group_topic_relation表
func (s *Editorial) SaveTopicByGroupId(topic int64,groupId int64){}

func (s *Editorial) UpdateCateGoryTotalNum(id int64,group GroupJsonData) (bool) {
	   total := group.Data.TotalCount
	   Model.UpdateCateGoryTotalNum(id, total)
	   return  true
}

