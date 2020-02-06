package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
)

type UserTempDashboardService struct {

}
/**
	根据id ， 来生成对应的信息
 */
func ( self *UserTempDashboardService ) FindById ( o orm.Ormer , id int64 ) ( *bean.UserTempDashboard , error ){
	var dashboard = &bean.UserTempDashboard{}
	dashboard.Id = id
	var readErr = o.Read( dashboard )
	return dashboard , readErr
}

func ( self *UserTempDashboardService ) New( o orm.Ormer , dashboard *bean.UserTempDashboard ) ( *bean.UserTempDashboard , error ) {
	dashboard.DashboardBase.New()
	var _ , insertErr = o.Insert( dashboard )
	return dashboard , insertErr
}

func ( self *UserTempDashboardService ) NewById( o orm.Ormer , id int64 ) ( *bean.UserTempDashboard , error ) {
	var dashboard = &bean.UserTempDashboard{}
	dashboard.Id = id
	return self.New( o ,dashboard )
}

var USER_TEMP_DASHBOARD_SERVICE_INSTANCE = &UserTempDashboardService{}

func GetUserTempDashboardServiceInstance() *UserTempDashboardService {
	return USER_TEMP_DASHBOARD_SERVICE_INSTANCE
}