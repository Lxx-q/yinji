package service

import (
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
)

type ResourceService struct {
	db.OrmService
}


func ( self *ResourceService) Insert( o orm.Ormer , resource *bean.Resource) error {
	var _ , insertErr = o.Insert( resource)
	return insertErr
}

var RESOURCE_SERVICE_INSTANCE = &ResourceService{}

func GetResourceServiceInstance() *ResourceService{
	return RESOURCE_SERVICE_INSTANCE
}

