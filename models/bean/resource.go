package bean

import "yinji/models/base"

type Resource struct{
	base.IdStruct
	HostResourceId int64 `orm:"column(host_resource_id)" json:"resourceId"`
}

const (
	RESOURCE_TYPE_HOST = 1;
)

//对应的信息
func ( self *Resource) TableName() string {
	return GetResourceTableName()
}

func GetResourceTableName() string{
	return "resource"
}
