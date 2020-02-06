package bean

import (
	"yinji/models/base"
	"time"
)


type AudioDashboard struct {
	base.IdStruct
	base.CreateTimeAndModifyTimeStruct
	base.DashboardBase
}

func ( self *AudioDashboard) NewEntity( t time.Time){
	self.IdStruct.NewEntity( t )
	self.CreateTimeAndModifyTimeStruct.NewEntity( t )
	self.LoveCount = 0
	self.ForwardCount = 0
	self.CollectionCount = 0
}

func ( self *AudioDashboard) New(){
	var current = time.Now()
	self.NewEntity( current )
}

func ( self *AudioDashboard) TableName() string {
	return GetAudioUserDashboardTableName()
}

func GetAudioUserDashboardTableName() string{
	return "audio_dashboard"
}

