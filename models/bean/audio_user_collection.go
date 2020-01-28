package bean

import (
	"yinji/models/base"
	"time"
)

//对应的收藏列表

type AudioUserCollection struct {
	AudioId int64 `orm:"column(audio_id)" json:"audioId"`//对应 audioId
	UserId int64 `orm:"column(user_id)" json:"userId"`
	base.IdStruct //对应的id主键
	base.CreateTimeStruct //对应的创建时间结构体
}

func ( self *AudioUserCollection ) NewEntity( time time.Time ){
	self.IdStruct.NewEntity( time )
	self.CreateTimeStruct.NewEntity( time )
}

func ( self *AudioUserCollection ) New(){
	var current = time.Now()
	self.NewEntity(current)
}


func ( self *AudioUserCollection) TableName() string {
	return GetAudioUserCollectionTableName()
}

func GetAudioUserCollectionTableName() string{
	return "audio_user_collection"
}
