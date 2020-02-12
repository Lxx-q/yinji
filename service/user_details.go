package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
)

type UserDetailsService struct {

}

/**
	根据，目标的id 来进行搜索
 */
func ( self *UserDetailsService ) FindById( o orm.Ormer , id int64 ) ( *bean.UserDetails , error ){
	var details = &bean.UserDetails{}
	details.Id = id
	var readErr = o.Read(details)
	details.Parse()
	return details , readErr
}

var USER_DETAILS_SERVICE_INSTANCE = &UserDetailsService{}

func GetUserDetailsServiceInstance() *UserDetailsService{
	return USER_DETAILS_SERVICE_INSTANCE
}