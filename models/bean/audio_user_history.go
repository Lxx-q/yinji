package bean

import (
	"yinji/models/base"
	"time"
)

type AudioUserHistroy struct {
	base.IdStruct
	base.CreateTimeAndModifyTimeStruct
	UserId int64 `orm:"column(user_id)" json:"userId"`
	AudioId int64 `orm:"column(audio_id)" json:"audioId"`
	BrowseAllCount int64 `orm:"column(browse_all_count)" json:"browseAllCount"`
	LastTime int `orm:"column(last_time)" json:"lastTime"`
}

func ( self *AudioUserHistroy ) NewEntity( t time.Time ) {
	self.IdStruct.NewEntity( t )
	self.CreateTimeAndModifyTimeStruct.NewEntity( t )
}

func ( self *AudioUserHistroy ) New(){
	var t = time.Now()
	self.NewEntity( t )
}

func ( self *AudioUserHistroy) TableName() string {
	return  GetAudioUserHistroyTableName()
}

func GetAudioUserHistroyTableName() string{
	return "audio_user_history"
}