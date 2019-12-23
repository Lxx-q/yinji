package base

import (
	"yinji/service/db"
	"reflect"
	"time"
	"strconv"
)

type IdStruct struct {
	Id int64 `orm:"column(id)" json:"id"`
	Code string `orm:"column(code)" json:"code"`
}

func ( self *IdStruct ) NewEntity( current time.Time){
	var code_service = db.GetCodeServiceInstance();

	var code_int_64 = code_service.BuildCode( reflect.TypeOf(self) , current );

	//设置对应的信息
	self.Id = code_int_64;

	//注意 ， 这个生成id 的策略 ， 只是暂时生成的
	self.Code = strconv.FormatInt( self.Id , 10 )

}

func ( self *IdStruct ) New(){
	var current = time.Now()
	self.NewEntity( current )
}


