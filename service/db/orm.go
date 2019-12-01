package db

import (
	"github.com/astaxie/beego/orm"
	"errors"
)

type OrmService struct {

}

func ( self *OrmService) Transaction(function func( o orm.Ormer) (interface{} , error) )  (interface{} , error){
	var o =orm.NewOrm()

	var err = o.Begin()

	if err != nil {
		return nil , err
	}

	result , jdbcErr := self.ExecJdbc( o , function);

	if( jdbcErr != nil){
		//倘若不等于 nil , 那么 说明出现错误 ， 那么便开始进行相对应的 回滚
		o.Rollback()
	}else{
		//提交对应的 信息
		o.Commit()
	}

	//最后将二者一同返回
	return result , jdbcErr
}

func ( self *OrmService) Jdbc(  function func( o orm.Ormer) (interface{} , error) ) (interface{} , error){
	var o = orm.NewOrm();
	return self.ExecJdbc( o,function);
}

//下面开始对应的caozuo
func ( self *OrmService) ExecJdbc( o orm.Ormer , function func( o orm.Ormer) (interface{} , error) )  (interface{} , error){
	result , err := function( o );
	return result , err;
}


//无论进行上面样的操作 ， 都会进行对应的 回滚 ， 我们 使用对应的 方法 ， 来进行 对应的 操作
func( self *OrmService) RollBack( function func( o orm.Ormer) (interface{} , error) )(interface{} , error){
	return self.Transaction(func(o orm.Ormer) (interface{}, error) {
		//在这里的操作 ， 我们便可以将他修改为此，无论，最后的结果是如何 ，我们都会返回对应的错误
		/**
			当真的有错误的时候 ， 便会返回对应的错误的类型， 如果真的出现了错误 ， 那么就返回对应的错误的操作
		 */
		var result , error = function( o );

		if( error == nil){
			error = errors.New("Now ,let me start rollback");
		}

		return result , error;
	});
}


var ORM_SERVICE_INSTANCE = OrmService{}

func GetOrmServiceInstance() *OrmService{
	return &ORM_SERVICE_INSTANCE;
}