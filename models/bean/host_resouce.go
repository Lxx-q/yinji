package bean

import (
	"yinji/models/base"
	"time"
)

type HostResource struct {

	base.IdStruct

	base.CreateTimeAndModifyTimeStruct
	Suffix string `orm:"column(suffix)" json:"suffix"`
	Path string `orm:"column(path)" json:"path"`
	Host string `orm:"column(host)" json:"host"`
	Port int `orm:"column(port)" json:"port"`

}

func ( self *HostResource) NewEntity( t time.Time ){
	self.IdStruct.NewEntity( t )
	self.CreateTimeAndModifyTimeStruct.NewEntity( t )
}

func ( self *HostResource ) New(){
	var t = time.Now()
	self.NewEntity( t )
}

func ( self *HostResource) TableName() string {
	return GetHostResourceTableName()
}

func GetHostResourceTableName() string{
	return "host_resource"
}
