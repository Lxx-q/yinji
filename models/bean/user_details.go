package bean

import (
	"yinji/models/base"
	"time"
)

type UserDetails struct {
	base.IdStruct
	base.ModifyTimeStruct

	Introduction string `orm:"column(introduction)" json:"introduction"`
	Address      string `orm:"column(address)" json:"address"`
	Sex          int    `orm:"column(sex)" json:"sex"`
	Birthday     time.Time `orm:"column(birthday)" json:"birthday"`
}

func( self *UserDetails ) NewEntity( t time.Time ){
	self.Birthday = t
	self.IdStruct.NewEntity( t )
}

func( self *UserDetails ) New(){
	var current = time.Now()
	self.NewEntity( current )
}

func ( self *UserDetails) TableName() string {
	return GetUserDetailsTableName()
}

func GetUserDetailsTableName() string{
	return "user_details"
}

