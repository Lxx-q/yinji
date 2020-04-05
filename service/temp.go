package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
)

func GetHostResource( o  orm.Ormer ,resourceId int64) (*bean.HostResource , error){
	var resource = bean.Resource{}
	resource.Id = resourceId
	var readResourceErr = o.Read(&resource)
	if readResourceErr != nil {
		return nil , readResourceErr
	}
	var hostResourceId = resource.HostResourceId
	var hostResource=bean.HostResource{}
	hostResource.Id = hostResourceId
	var readHostResourceErr = o.Read(&hostResource)

	return &hostResource , readHostResourceErr

}