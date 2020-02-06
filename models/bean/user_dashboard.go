package bean

import (
	"yinji/models/base"
)

type UserDashboard struct {
	base.IdStruct
	base.DashboardBase
}

func ( self *UserDashboard) TableName() string {
	return GetUserDashboardTableName()
}

func GetUserDashboardTableName() string{
	return "user_dashboard"
}