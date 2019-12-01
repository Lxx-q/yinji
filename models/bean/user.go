package bean

import (
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type User struct {

	//Id string `orm:"column(id);pk"`
	Name string
	Image string `orm:column(image)`
	EntityBase
}

func (u *User) TableName() string {
	return getUserTableName();
}

func getUserTableName() string{
	return "user";
}
func( self *User) NewToDbFunc() func(o orm.Ormer) (interface{}, error){

	return func(o orm.Ormer) (interface{}, error) {
		self.New();

		self.Code = "u_" +  strconv.FormatInt(self.Id,10)
		return o.Insert( self );
	};
}

func( self *User) NewToDb(){
	var ormService = db.GetOrmServiceInstance();
	ormService.Transaction( self.NewToDbFunc());
}

func ( self *User) RefreshToDbFunc() func(o orm.Ormer) (interface{}, error){


	return func(o orm.Ormer) (interface{}, error) {
		self.Refresh();
		return o.Update( self );
	};

}

func (self *User) RefreshToDb(){
	var ormService = db.GetOrmServiceInstance();
	ormService.Transaction(self.RefreshToDbFunc());
}