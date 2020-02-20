package bean

import (
	"yinji/models/base"
	"time"
)

type AudioFast struct {
	base.IdStruct
	base.CreateTimeAndModifyTimeStruct
	StartTime int `orm:"column(start_time)" json:"startTime"`
	EndTime int `orm:"column(end_time)" json:"endTime"`
	Introduction string `orm:"column(introduction)" json:"introduction"`
	UserId int64 `orm:"column(user_id)" json:"userId"`
	AudioId int64 `orm:"column(audio_id)" json:"audioId"`
}

func ( self *AudioFast ) NewEntity( t time.Time ){
	self.IdStruct.NewEntity(t)
	self.CreateTimeAndModifyTimeStruct.NewEntity(t)
}

func ( self *AudioFast ) New(){
	var t = time.Now()
	self.NewEntity( t )
}

func ( self *AudioFast ) TableName() string {
	return  GetAudioFastTableName()
}

func GetAudioFastTableName() string {
	return "audio_fast"
}
