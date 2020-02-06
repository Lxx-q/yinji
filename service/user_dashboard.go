package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
)

type UserDashboardService struct {

}

/**
	根据id ， 来生成对应的信息
 */
func ( self *UserDashboardService ) FindById ( o orm.Ormer , id int64 ) ( *bean.UserDashboard , error ){
	var dashboard = &bean.UserDashboard{}
	dashboard.Id = id
	var readErr = o.Read( dashboard )
	return dashboard , readErr
}

func ( self *UserDashboardService ) New( o orm.Ormer , dashboard *bean.UserDashboard ) ( *bean.UserDashboard , error ) {
	dashboard.DashboardBase.New()
	var _ , insertErr = o.Insert( dashboard )
	return dashboard , insertErr
}

func ( self *UserDashboardService ) NewById( o orm.Ormer , id int64 ) ( *bean.UserDashboard , error ) {
	var dashboard = &bean.UserDashboard{}
	dashboard.Id = id
	return self.New( o ,dashboard )
}

var USER_DASHBOARD_SERVICE_INSTANCE = &UserDashboardService{}

func GetUserDashboardService() *UserDashboardService {
	return USER_DASHBOARD_SERVICE_INSTANCE
}
