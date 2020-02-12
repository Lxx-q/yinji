package bean

import "yinji/models/base"

type UserDetails struct {
	base.IdStruct
	base.ModifyTimeStruct

	Introduction string `orm:"column(introduction)" json:"introduction"`
	Address      string `orm:"column(address)" json:"address"`
	Sex          int    `orm:"column(sex)" json:"sex"`
	Birthday     string `orm:"column(birthday)" json:"birthday"`
}

func ( self *UserDetails) TableName() string {
	return GetUserDetailsTableName()
}

func GetUserDetailsTableName() string{
	return "user_details"
}

