package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
)

type LoginService struct {

}

func ( self *LoginService ) ReadLogin( o orm.Ormer , id int64 ) ( *bean.Login , error ){
	var login = &bean.Login{}
	login.Id = id
	var readErr = o.Read( login )

	return login , readErr
}

func ( self *LoginService ) UpdateLogin( o orm.Ormer , login *bean.Login ) ( *bean.Login , error ){
	var _ , updateErr = o.Update(login)
	return login , updateErr
}

var LOGIN_SERVICE_INSTANCE = &LoginService{}

func GetLoginServiceInstance() *LoginService{
	return LOGIN_SERVICE_INSTANCE
}

