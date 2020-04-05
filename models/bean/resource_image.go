package bean

import (
	"yinji/models/base"
	"time"
)

const (
	RESOURCE_IMAGE_ORIGIN_TYPE int = 0
	RESOURCE_IMAGE_THUMB_TYPE  int = 1
)

type ResourceImage struct {
	base.IdStruct
	base.CreateTimeAndModifyTimeStruct
	OriginResourceId int64 `orm:"column(origin_resource_id)" json:"originResourceId"`
	ThumbResourceId int64 `orm:"column(thumb_resource_id)" json:"thumbResourceId"`
}

func ( self *ResourceImage) NewEntity( t time.Time ){
	self.IdStruct.NewEntity( t )
	self.CreateTimeAndModifyTimeStruct.NewEntity( t )
}

func ( self *ResourceImage ) New(){
	var t = time.Now()
	self.NewEntity( t )
}


//对应的信息
func ( self *ResourceImage) TableName() string {
	return GetResourceImageTableName()
}

func GetResourceImageTableName() string{
	return "resource_image"
}

