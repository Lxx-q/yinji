package bean

import "yinji/models/base"

type UserTempDashboard struct {
	base.IdStruct
	base.DashboardBase
}


func ( self *UserTempDashboard) TableName() string {
	return GetUserTempDashboardTableName()
}

func GetUserTempDashboardTableName() string{
	return "user_temp_dashboard"
}