package editorial

import (
	"engine"
	"visualchina/persist/editorialPersist"
	"libary/upload"
	"logger"
)

//图片页
func ParseEditorialPic(contents []byte,url string,args engine.RequestArgs) (ret engine.ParseResult) {
	 //在ES中先判断一下是否已经存在

	  pic, err := editorialPersist.ParsePicJson(contents)
	  if err != nil || pic.Data.PicInfo.Id ==0 {
	     	return
	  }
	 list,ok := args.Data.(editorialPersist.PicList)
	 if !ok{
	     return ret
	 }
	  pic.Data.PicInfo.ImgId =  upload.UploadToQiniu(pic.Data.PicInfo.Url)
	  pic.Data.PicInfo.ImageDate = list.ImgDate
	  id, _ := editorial.SavePic(args.CategoryId, pic.Data)
	  if id > 0{
	  	  //保存PIC 的详情信息到 group_pic_detail
		  editorial.SavePicDetail(pic.Data)
		  logger.Info.Println(pic.Data.PicInfo.Id,"保存成功")
	  }
	  return ret
}
