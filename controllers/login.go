package controllers

import (
	"yinji/models/bean"
	"time"
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	BeegoController
}

//对应的 注册接口
func ( controller *LoginController ) Register(){

	var time = time.Now()

	var user = bean.User{}
	user.NewEntity(time)

	//根据对方输入的密码 ， 之后进行生成对应的 信息
	var login = bean.Login{}

	login.NewEntity( time )

	var ormService = db.GetOrmServiceInstance()

	ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		o.Insert( login )
		return nil, nil
	})
}

/**
	由于一次性开发容错性高的接口 ，容易造成 高耦合 ，
	因此  ，我个人在这里建议先把登录接口分开，再合并
 */

 //利用账号密码登录
func ( self *LoginController ) LoginByAccount(){
	//开始对应的输出信息

	//获取对应的账号密码 , 并且进行组装到对应的 login 之中
	var account = self.GetString("account")

	var password = self.GetString("password")

	var login = bean.Login{}

	login.Acount = account

	login.Password = password

	var user = bean.User{}

	//这里 ， 我们设定对应的程序

	//程序进入下一个阶段
	var ormService = db.GetOrmServiceInstance()

	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {

		var readLoginErr = o.Read(&login , "account","password")

		if readLoginErr != nil {
			//查询错误 ，
			return nil , readLoginErr
		}

		//再之后 ， 我们便可以获取到对应的 user 的 信息

		user.Id = login.Id;

		var readUserErr = o.Read( &user )

		if readUserErr != nil {
			return nil , readUserErr
		}

		return &user , nil
	})

	if jdbcErr != nil {
		self.Fail( jdbcErr )
	}

	self.CruSession.Set("LOCAL_USER" , user)

	//之后返回结果
	self.Json( user )
}