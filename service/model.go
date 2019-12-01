package service

import "github.com/astaxie/beego/orm"

type ModelService struct {

}

func( self *ModelService ) doJdbcFunction(  function func( ormer orm.Ormer) (interface{} , error)) interface{}{
	var ormer = orm.NewOrm();

	ormer.Begin();

	var result , err = function(ormer);

	//倘若 error 中存在值， 那么 便 重置
	if( err != nil ){
		ormer.Rollback();
		return nil;
	}

	ormer.Commit();

	return result;

}

