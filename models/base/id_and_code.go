package base

import (
	"yinji/service/db"
	"reflect"
	"time"
	"strconv"
)

//开始设置对应的 id struct 体
type IdStruct struct {
	Id int64 `orm:"column(id);pk" json:"id"`
}

//对应的信息
func ( self *IdStruct ) NewEntity( current time.Time){
	var code_service = db.GetCodeServiceInstance();

	var code_int_64 = code_service.BuildCode( reflect.TypeOf(self) , current );

	//设置对应的信息
	self.Id = code_int_64;
}

func ( self *IdStruct ) New(){
	var current = time.Now()
	self.NewEntity( current )
}

//开始对应的信息
type CodeStruct struct {
	Code string `orm:"column(code)" json:"code"`
}

func ( self *CodeStruct ) NewEntity( current time.Time ){

}

func ( self *CodeStruct ) NewEntityById( id int64 ){
	self.Code = strconv.FormatInt( id , 10 )
}

//对应的信息的绑定
type IdAndCodeStruct struct {
	IdStruct
	CodeStruct
}


func ( self *IdAndCodeStruct ) NewEntity( current time.Time ){
	/*
	var code_service = db.GetCodeServiceInstance();

	var code_int_64 = code_service.BuildCode( reflect.TypeOf(self) , current );

	//设置对应的信息
	self.Id = code_int_64;

	//注意 ， 这个生成id 的策略 ， 只是暂时生成的
	self.Code = strconv.FormatInt( self.Id , 10 )
	*/
	self.IdStruct.NewEntity( current )
	self.CodeStruct.NewEntityById(self.Id)
}

func ( self *IdAndCodeStruct ) New(){
	var current = time.Now()
	self.NewEntity( current )
}


