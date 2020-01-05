package bean

import "yinji/models/base"

type ResourceAudio struct {
	base.IdStruct
	OriginFile int64 `orm:"column(origin_file)" json:"origin_file"`
	base.CreateTimeAndModifyTimeStruct
}

func ( self *ResourceAudio) TableName() string {
	return GetResourceAudioTableName()
}

func GetResourceAudioTableName() string{
	return "resource_audio"
}