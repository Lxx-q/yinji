package bean

import (
	"yinji/models/base"
	"time"
)

type UserTempDashboard struct {
	base.IdStruct
	base.DashboardBase
	UserId int64 `orm:"column(user_id)" json:"userId"`
	WriteDate time.Time `orm:"column(write_date)" json:"writeDate"`
}

func ( self *UserTempDashboard ) NewEntity( t time.Time ){
	self.IdStruct.NewEntity( t )
	self.WriteDate = t
	self.DashboardBase.New()
}

func ( self *UserTempDashboard ) New(){
	var t = time.Now()
	self.NewEntity( t )
}



func ( self *UserTempDashboard) TableName() string {
	return GetUserTempDashboardTableName()
}

func GetUserTempDashboardTableName() string{
	return "user_temp_dashboard"
}