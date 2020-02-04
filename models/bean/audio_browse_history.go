package bean

import (
	"yinji/models/base"
	"time"
)

// 目标 音频 浏览其实记录
type AudioBrowseHistroy struct {
	base.IdStruct
	BrowseAllCount int64
	base.CreateTimeAndModifyTimeStruct
}

func ( self *AudioBrowseHistroy) NewEntity( t time.Time ){
	self.IdStruct.NewEntity( t )
	self.CreateTimeAndModifyTimeStruct.NewEntity( t )
}

func ( self *AudioBrowseHistroy) New(){
	var t = time.Now()
	self.NewEntity( t )
}

func ( self *AudioBrowseHistroy) TableName() string {
	return GetAudioBrowseHistroyTableName()
}

func GetAudioBrowseHistroyTableName() string {
	return "audio_browse_history"
}