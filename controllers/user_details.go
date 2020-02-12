package controllers

import (
	"yinji/service/db"
	"yinji/service"
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
)

type UserDetailsController struct {
	BeegoController
}

func ( self *UserDetailsController ) FindById(){
	var id , getIdErr = self.GetInt64("id")

	if getIdErr != nil {
		self.FailJson( getIdErr )
		return
	}

	var ormService = db.GetOrmServiceInstance()
	var userDetails = service.GetUserDetailsServiceInstance()

	var details  *bean.UserDetails

	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var findErr error
		details , findErr = userDetails.FindById( o , id )
		return details , findErr
	})

	if jdbcErr != nil {
		self.FailJson( jdbcErr )
		return
	}

	self.Json( details )

}
