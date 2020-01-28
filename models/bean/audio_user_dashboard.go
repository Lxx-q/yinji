package bean

import (
	"yinji/models/base"
	"time"
)

type AudioUserDashboard struct {
	base.IdStruct
	base.CreateTimeAndModifyTimeStruct
	AudioId int64 `orm:"column(audio_id)" json:"audioId"`//对应 audioId
	DashboardType int  `orm:column("dashboard_type)" json:"dashboardType"`//对应的仪表盘数据
	Count int64  `orm:"column(count)" json:"count"`//数据整合
}

func ( self *AudioUserDashboard ) NewEntity( t time.Time){
	self.IdStruct.NewEntity( t )
	self.CreateTimeAndModifyTimeStruct.NewEntity( t )
}

func ( self *AudioUserDashboard ) New(){
	var current = time.Now()
	self.NewEntity( current )
}

func ( self *AudioUserDashboard) TableName() string {
	return GetAudioUserDashboardTableName()
}

func GetAudioUserDashboardTableName() string{
	return "audio_user_dashboard"
}

