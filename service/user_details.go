package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
	"errors"
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

func ( self *UserDetailsService ) New( o orm.Ormer , details *bean.UserDetails)( *bean.UserDetails , error ){
	var number , insertErr = o.Insert( details )
	if insertErr != nil {
		return nil , insertErr
	}
	if number == 0 {
		return nil, errors.New("the number is 0 ")
	}

	return details , nil

}

var USER_DETAILS_SERVICE_INSTANCE = &UserDetailsService{}

func GetUserDetailsServiceInstance() *UserDetailsService{
	return USER_DETAILS_SERVICE_INSTANCE
}