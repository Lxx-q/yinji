package service

import "github.com/astaxie/beego/orm"

func Limit( seter orm.QuerySeter ) orm.QuerySeter {
	return seter.Limit( LIMIT_COUNT )
}
