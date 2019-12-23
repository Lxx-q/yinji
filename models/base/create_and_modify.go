package base

import (
	"time"
)

/**
	创造时间的 struct
 */
type CreateTimeStruct struct {
	CreateTime time.Time `orm:"column(create_time)" json:"createTime"`
	CreateTimeStruct TimeStruct `orm:"-" json:createTimeStruct`
}

func(self *CreateTimeStruct) Parse( ){
	self.CreateTimeStruct.ParseTime( self.CreateTime )
}

func ( self *CreateTimeStruct ) NewEntity( t time.Time ){
	self.CreateTime = t
	self.Parse()
}

func (self *CreateTimeStruct ) New(){
	var t = time.Now();
	self.NewEntity(t)
}

/**
	修改时间的 struct
 */
type ModifyTimeStruct struct {
	ModifyTime time.Time `orm:"column(modify_time)" json:"createTime"`
	ModifyTimeStruct TimeStruct `orm:"-" json:"modifyTimeStruct"`
}

func(self *ModifyTimeStruct) Parse( ){
	self.ModifyTimeStruct.ParseTime( self.ModifyTime )
}

func ( self *ModifyTimeStruct ) NewEntity( t time.Time ){
	self.ModifyTime = t
	self.Parse()
}

func (self *ModifyTimeStruct ) New(){
	var t = time.Now();
	self.NewEntity(t)
}

func (self *ModifyTimeStruct ) RefreshEntity( t time.Time ){
	self.ModifyTime = t;
	self.Parse()
}

func (self *ModifyTimeStruct ) Refresh(){
	var t = time.Now()
	self.RefreshEntity( t )
}

type CreateTimeAndModifyTimeStruct struct{
	CreateTimeStruct
	ModifyTimeStruct
}

func (self *CreateTimeAndModifyTimeStruct) Parse(){
	self.CreateTimeStruct.Parse()
	self.ModifyTimeStruct.Parse()
}

func ( self *CreateTimeAndModifyTimeStruct ) NewEntity( t time.Time ){
	self.CreateTimeStruct.NewEntity(t)
	self.ModifyTimeStruct.NewEntity(t)
}

func (self *CreateTimeAndModifyTimeStruct ) New(){
	var t = time.Now();
	self.NewEntity(t)
}
