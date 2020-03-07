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

func ( self *UserController ) UpdateUser(){

	//先收集对应的信息
	var id , getIdErr = self.GetInt64("id")

	if getIdErr != nil{
		self.FailJson( getIdErr )
		return
	}

	var name  = self.GetString("name")

	//获取对应的 service intsance
	var ormService = db.GetOrmServiceInstance()
	var userService = service.GetUserServiceInstance()

	var user *bean.User

	var _ , jdbcErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		var findErr error
		user , findErr = userService.FindUserById( o , id)

		if findErr != nil {
			return nil , findErr
		}

		user.Name = name
		user.Refresh()
		var _ , updateErr = o.Update(user)
		return user , updateErr
	})

	if jdbcErr != nil {
		self.FailJson(jdbcErr)
		return
	}

	self.Json( user )
}

/**
	设置用户名与头像
*/
func( self  *UserController ) UpdateUserNameAndImage(){

	var id , getIdErr  = self.GetInt64("id")

	if getIdErr != nil {
		self.FailJson( getIdErr )
		return
	}

	var httpFileService = service.HttpFileService{}

	var userName = self.GetString("name")

	var userImage , userImageHeader , userImageErr = self.GetFile("image")

	if userImageErr != nil {
		self.FailJson(userImageErr)
		return
	}

	defer httpFileService.CloseMultipart(userImage)

	//得到新的文件名 ， 之后 并把它上传到对应的位置
	var userImageName = httpFileService.BuildFileName( id , userImageHeader )

	//在得到在服务器内的相对应路径
	var userImagePath = httpFileService.GetImagUserePath(userImageName)

	//之后， 我们将对应的服务器文件资源上传

	var path , uploadErr = httpFileService.Upload( userImagePath , userImage)

	if uploadErr != nil {
		return
	}

	var user = bean.User{}
	var ormService = db.GetOrmServiceInstance()

	var _ , transacErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		//现根据对应的id进行读取信息
		user.Id = id
		var readErr = o.Read(&user)
		if readErr != nil {
			return nil , readErr
		}

		user.Name = userName
		user.Image = path
		var _ , updateErr = o.Update(&user)
		if updateErr != nil {
			return nil , updateErr
		}
		return user , nil
	});

	if transacErr != nil {
		self.FailJson( transacErr )
		return
	}

	self.Json( user )


}
