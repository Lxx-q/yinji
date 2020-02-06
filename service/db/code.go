package db

import (
	"reflect"
	"time"
	"yinji/utils"
)

/*
*
  相对应的 CodeService  ， 来生成相对应的信息
*/

type CodeService struct {

}

/**
	生成对应的 信息
 */
func ( self *CodeService) BuildCode( targetType reflect.Type , t time.Time ) int64{
	//之后 我们 生成 对应的 信息
	var right int64 = 1000;
	var number int64 = t.Unix() ;

	var random = utils.Rann( 3 )
	//将后面
	return  right * number + int64( random );
}


var CODE_SERVICE_INSTANCE = &CodeService{};

func GetCodeServiceInstance() *CodeService{
	return CODE_SERVICE_INSTANCE;
}