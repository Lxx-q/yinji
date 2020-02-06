package controllers

import (
	"yinji/models/bean"
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
	"yinji/service"
)

type UserDashboardController struct {
	BeegoController
}

/**

 */
func ( self *UserDashboardController ) FindById(){
	var id , getIdErr = self.GetInt64("id")

	if getIdErr != nil {
		self.FailJson( getIdErr )
		return
	}

	var userDashboard  * bean.UserDashboard

	var ormService = db.GetOrmServiceInstance()
	var userDashboardService = service.GetUserDashboardService()

	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		_userDashboard , readErr := userDashboardService.FindById( o , id)

		if readErr == nil {
			userDashboard = _userDashboard
			return _userDashboard , readErr
		}

		_userDashboard , newErr := userDashboardService.New( o , _userDashboard )

		userDashboard = _userDashboard
		return _userDashboard , newErr
	})

	if jdbcErr != nil {
		self.FailJson( jdbcErr )
		return
	}

	self.Json( userDashboard )

}
