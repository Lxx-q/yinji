package controllers

import (
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
	"yinji/service"
	"yinji/models/bean"
)

type UserTempDashoardController struct {
	BeegoController
}

func ( self *UserTempDashoardController ) FindById() {
	var userId, getUserIdErr = self.GetInt64( "userId" )

	if getUserIdErr != nil {
		self.FailJson(getUserIdErr)
		return
	}

	var ormService = db.GetOrmServiceInstance()
	var userTempDashboardService = service.GetUserTempDashboardServiceInstance()

	var dashboard *bean.UserTempDashboard
	var _ , transactionErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		_dashboard , findErr := userTempDashboardService.FindById( o , userId)

		if findErr == nil {
			dashboard = _dashboard
			//说明查询出错
			return dashboard , findErr
		}

		_dashboard , newErr := userTempDashboardService.New(o , _dashboard , userId)
		dashboard = _dashboard
		return _dashboard , newErr

	})

	if transactionErr != nil {
		self.FailJson( transactionErr )
		return
	}

	self.Json( dashboard )
}
