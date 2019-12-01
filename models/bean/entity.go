package bean

import (
	"time"
	"reflect"
	"yinji/service/db"
	"strconv"
	"yinji/utils"
)

type BaseEntity struct {

	Id int64 `orm:"column(id);pk"`

	CreateTime time.Time
	//上次修改时间
	ModifyTime time.Time
}

func (self *BaseEntity) NewEntity( t time.Time){

	self.CreateTime = t;
	self.ModifyTime = t;
}

func (self *BaseEntity) New(){
	var current = time.Now();
	self.NewEntity(current);
}

func( self *BaseEntity) Refresh(){
	var current = time.Now();

	self.RefreshEntity( current );

}

func( self *BaseEntity) RefreshEntity( t time.Time ){
	self.ModifyTime = t;
}


type EntityBase struct {
	// 简化id
	Code string
	//名称 name
	// 创建时间

	BaseEntity
}

func (self *EntityBase) Typeof() reflect.Type{
	return reflect.TypeOf( self );
}

func ( self *EntityBase) New(){
	//新建力所能及的属性

	//获取当前的时间
	var current = utils.Now()

	self.BaseEntity.NewEntity( current )
	//获取当前的时间

	var code_service = db.GetCodeServiceInstance();

	var code_int_64 = code_service.BuildCode( reflect.TypeOf(self) , current );

	//设置对应的信息
	self.Id = code_int_64;

	//注意 ， 这个生成id 的策略 ， 只是暂时生成的
	self.Code = strconv.FormatInt( self.Id , 10 )


}



/**
	倘若对应的 Entity 需要更新 ， 那么 需要进行那些操作
 */
func ( self *EntityBase) Refresh(){
	self.BaseEntity.Refresh();
}

/*

func( self *EntityBase) NewToDb(){
	self.New();
	//进行操作 之后我们 再将对应的 操作 保存进入 其中
	var orm_service = db.GetOrmServiceInstance();

	orm_service.RunJdbcFunction(func(o orm.Ormer) (interface{}, error) {
		return o.Insert( self );
	});
}

func ( self *EntityBase) RefreshToDb(){
	self.Refresh();
	var orm_service = db.GetOrmServiceInstance();

	orm_service.RunJdbcFunction(func(o orm.Ormer) (interface{}, error) {
		return o.Update( self );
	})
}

*/

/**
	根据对应的 方法 ， 来进行获取对应的 code
 */
