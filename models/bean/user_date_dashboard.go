package bean

import (
	"yinji/models/base"
	"time"
)

type UserDateDashboard struct {
	base.IdStruct
	UserId int64 `orm:"column(user_id)" json:"userId"`
	WriteDate time.Time `orm:"column(write_date)" json:"writeDate"`
	//BrowseAllCount int64 `orm:column"browse_all_count"`
	base.DashboardBase
}

func ( self *UserDateDashboard ) NewEntity( t time.Time ){
	self.IdStruct.NewEntity( t )
	self.WriteDate = t
	self.DashboardBase.New()
}

func ( self *UserDateDashboard ) New(){
	var t = time.Now()
	self.NewEntity( t )
}


func ( self *UserDateDashboard) TableName() string {
	return GetUserDateDashboardTableName()
}

func GetUserDateDashboardTableName() string{
	return "user_date_dashboard"
}