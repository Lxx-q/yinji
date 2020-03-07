package bean

import (
	"time"
	"yinji/models/base"
)

type AudioTempDashboard struct {
	base.IdStruct
	AudioId int64 `orm:"column(audio_id)" json:"audioId"`
	WriteDate time.Time `orm:"column(write_date)" json:"writeDate"`
	base.DashboardBase
}

func ( self *AudioTempDashboard) NewEntity( t time.Time ){
	self.IdStruct.NewEntity( t )
	self.WriteDate = t
	self.DashboardBase.New()
}

func ( self *AudioTempDashboard) New(){
	var t = time.Now()
	self.NewEntity( t )
}


func ( self *AudioTempDashboard) TableName() string {
	return GetAudioTempTempDashboard()
}

func GetAudioTempTempDashboard() string{
	return "audio_temp_dashboard"
}