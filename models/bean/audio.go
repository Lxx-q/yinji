package bean

import (
	"github.com/astaxie/beego/orm"
	"yinji/service/db"
	"strconv"
)

type Audio struct {

	//Id string `orm:"column(id);pk"`
	Name string `orm:column(name)`
	Url string `orm:column(url)"`
	Image string `orm:column(image)`
	UserId int64 `orm:column(user_id)`
	TimeLength int `orm:column(time_length)`

	EntityBase
}


func ( self *Audio) TableName() string {

	return GetAduioTableName();
}

func GetAduioTableName() string{
	return "audio";
}


func( self *Audio) NewToDbFunc() func(o orm.Ormer) (interface{}, error){

	return func(o orm.Ormer) (interface{}, error) {
		self.New();
		self.Code = "m_" +  strconv.FormatInt(self.Id,10)
		return o.Insert( self );
	};
}

func( self *Audio) NewToDb(){
	var ormService = db.GetOrmServiceInstance();
	ormService.Transaction( self.NewToDbFunc());
}

func ( self *Audio) RefreshToDbFunc() func(o orm.Ormer) (interface{}, error){


	return func(o orm.Ormer) (interface{}, error) {
		self.Refresh();
		return o.Update( self );
	};

}

func (self *Audio) RefreshToDb(){
	var ormService = db.GetOrmServiceInstance();
	ormService.Transaction(self.RefreshToDbFunc());
}