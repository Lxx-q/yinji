package bean

import "yinji/models/base"

type Resource struct{

	Path string `orm:"column(path)" json:"json"`
	Server string `orm:"column(server)" json:"server"`
	Suffix string `orm:"column(suffix)" json:"suffix"`

	base.IdAndCodeStruct
	//设置对应的 创造时间
	base.CreateTimeStruct
}

//对应的信息
func ( self *Resource) TableName() string {
	return GetResourceTableName()
}

func GetResourceTableName() string{
	return "resource"
}
