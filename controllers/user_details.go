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

/**
	刷新user details 的信息
*/
func ( self *UserDetailsController ) UpdateUserDetails(){

	var id , getIdErr = self.GetInt64("id")

	if getIdErr != nil{
		self.FailJson( getIdErr )
		return
	}

	var introduction = self.GetString("introduction")

	var sex , getSexErr = self.GetInt("sex")

	if getSexErr != nil {
		self.FailJson(getSexErr)
		return
	}

	var address = self.GetString("address")

	//获取对应的 service intsance
	var ormService = db.GetOrmServiceInstance()
	var userDetails = service.GetUserDetailsServiceInstance()

	var details *bean.UserDetails

	var _ , transacErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		var findErr error
		details , findErr = userDetails.FindById( o , id)

		if findErr != nil {
			return nil , findErr
		}

		details.Introduction = introduction
		details.Sex = sex
		details.Address = address
		details.Refresh()

		var _ , updateErr = o.Update( details )
		return details , updateErr
	})

	if transacErr != nil {
		self.FailJson( transacErr )
		return
	}

	self.Json( details )
}