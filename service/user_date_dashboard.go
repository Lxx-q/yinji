package service

import "yinji/models/bean"

type UserDateDashboardService struct {

}

func ( self *UserDateDashboardService ) NewByTemp( tempDashboard *bean.UserTempDashboard) {
	var userDateDashboard = &bean.UserDateDashboard{}
	userDateDashboard.New()
	userDateDashboard.DashboardBase = tempDashboard.DashboardBase
}

var USER_DATE_DASHBORAD_SERVICE_INSTANCE = &UserDateDashboardService{}


func GetUserDateDashboardServiceInstance() *UserDateDashboardService{
	return USER_DATE_DASHBORAD_SERVICE_INSTANCE
}
