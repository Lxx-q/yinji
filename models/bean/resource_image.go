package bean

import "yinji/models/base"

type ResourceImage struct {
	//还差一个对应的id值
	base.IdStruct

	Name string `orm:"column(name)" json:"name"`
	OriginFile int64 `orm:"column(origin_file)" json:"origin_file"`
	CompressFile30 int64 `orm:"column(compress_file_30)" json:"compress_file_30"`
	CompressFile60 int64 `orm:"column(compress_file_60)" json:"compress_file_60"`
	CompressFile80 int64 `orm:"column(compress_file_80)" json:"compress_file_80"`
	//对应的信息
	base.CreateTimeAndModifyTimeStruct
}

func ( self *ResourceImage) TableName() string {
	return GetResourceImageTableName()
}

func GetResourceImageTableName() string{
	return "resource_image"
}
