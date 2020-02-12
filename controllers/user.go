package controllers

import (
	"yinji/models/bean"
	"yinji/service/db"
	"yinji/service"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	BeegoController
}

func( self *UserController ) NewToDb(){
	//来对相对应的 属性进行更改
	//进行获取对应的 参数
	var name = self.GetString("name");

	var user = new(bean.User);

	user.Name = name;

	user.NewToDb();

	self.Json( user );
}

/*
*	利用 相对应的 属性 ， 来创建对应的 信息
*/
func ( self *UserController ) UploadTest(){
	var new_user = bean.User{};

	new_user.Code="dqw" + "dw";
	new_user.Name=":dqw";
	new_user.Image="dqwdqw";

	new_user.NewToDb();

	//之后返回对应的 属性

	self.Json( new_user );

}

/**
	根据对应的 userId [ 用户 ] 来进行搜索
*/
func ( self *UserController ) FindUserById(){
	var id , getIdErr = self.GetInt64("id")

	if getIdErr != nil {
		self.FailJson( getIdErr )
		return
	}

	var ormService = db.GetOrmServiceInstance()
	var userService = service.GetUserServiceInstance()

	var result , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		return userService.FindUserById( o , id )
	})

	if jdbcErr != nil {
		self.FailJson( jdbcErr )
		return
	}

	self.Json( result )

}

/**
	更新user 和 userDetails 的信息
*/

func ( self *UserController ) UpdateUserAndDetails(){

	//先收集对应的信息
	/*
	var id , getIdErr = self.GetInt("id")

	if getIdErr != nil{
		self.FailJson( getIdErr )
		return
	}

	var name  = self.GetString("name")

	var introduction = self.GetString("introduction")

	var sex , getSexErr = self.GetInt("sex")

	if getSexErr != nil {
		self.FailJson(getSexErr)
		return
	}

	var address = self.GetString("address")
	*/

	self.String("helloworld")
}