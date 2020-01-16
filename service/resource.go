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

// 根据对应的 url 信息 来转化为对应的 url 路径
func ( self *ResourceService ) toUrl( resource *bean.Resource ) string{
	var server = resource.Server
	if server != "" {

	}
	var path = resource.Path
	return path
}
