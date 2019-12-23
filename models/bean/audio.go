package bean

import (
	"github.com/astaxie/beego/orm"
	"yinji/service/db"
	"strconv"
)

type Audio struct {

	//Id string `orm:"column(id);pk"`
	Name string `orm:"column(name)"  json:"name"`
	Url string `orm:"column(url)" json:"url"`
	Image string `orm:"column(image)" json:"image"`
	UserId int64 `orm:"column(user_id)" json:"userId"`
	TimeLength int `orm:"column(time_length)" json:"timeLength"`
	//商业介绍
	Introduction string `orm:"column(introduction)" json:"introduction"`
	FolderId int64 `orm:"column(folder_id)" json:"folderId"`

	EntityBase
}


func ( self *Audio) TableName() string {
	return GetAduioTableName()
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