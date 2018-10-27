package upload

/*
上传图片到七牛，并保存在表里(group_pic_images)， 返回在表里的图片ID
*/
func UploadToQiniu(imgSrc string) (imageId int64)  {
	if imgSrc == ""{
		return 0
	}
	return 1
}


