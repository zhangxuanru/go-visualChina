package persist

import (
	"hash/crc32"
	"encoding/hex"
	"visualchina/Model"
	"time"
	"strconv"
	"libary/query"
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


