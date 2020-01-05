package service

import (
	"yinji/service/db"
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
)

type ResourceService struct {
	db.OrmService
}

//获取对应的信息
func ( self *ResourceService ) FindById( id int64 ) *bean.Resource{
	var resource = bean.Resource{}
	resource.Id = id

	self.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var err = o.Read(&resource , "id" )
		return nil , err
	})

	return &resource
}

//根据对应的 Resource  ( 程序 意义 ， 根据对应的 Resource 转化成 对应的信息)
func ( self *ResourceService ) TransferTo( resource *bean.Resource) string{
	return ""
}