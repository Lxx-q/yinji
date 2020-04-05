package bean

import (
	"yinji/models/base"
	"time"
)

type ResourceAudio struct {
	base.IdStruct
	base.CreateTimeAndModifyTimeStruct
	OriginResourceId int64 `orm:"column(origin_resource_id)" json:"origin_resource_id"`
}

func ( self *ResourceAudio) NewEntity( t time.Time ){
	self.IdStruct.NewEntity( t )
	self.CreateTimeAndModifyTimeStruct.NewEntity( t )
}

func ( self *ResourceAudio ) New(){
	var t = time.Now()
	self.NewEntity( t )
}


//对应的信息
func ( self *ResourceAudio) TableName() string {
	return GetResourceAudioTableName()
}

func GetResourceAudioTableName() string{
	return "resource_audio"
}



