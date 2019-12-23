package service

import (
	"yinji/models/bean"
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
)

type UserService struct {

}

//下面我们 根据 对应的 信息 来进行 设置

func( self *UserService) FindUserById ( id int64) *bean.User{

	//对应的 信息
	var user = new(bean.User);
	user.Id = id;
	var ormService = db.GetOrmServiceInstance();

	ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		o.Read( user ,"Id" )
		return user , nil
	});

	return user;
}

func( self *UserService) FindUsersByIds( codes [] int64) []*bean.User{
	var users []*bean.User
	var ormService = db.GetOrmServiceInstance();
	ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {

		var qs = self.OrmFuncByIds( o , codes)
		return qs.All(&users)
	});

	return users;
}

func ( self *UserService ) OrmFuncByIds( o orm.Ormer , codes [] int64 ) orm.QuerySeter {
	var qs = o.QueryTable("user");
	qs = qs.Filter("user__Id__in" , codes )
	return qs
}

func ( self *UserService ) ParseIdsToMap( users []*bean.User ) map[int64] *bean.User {
	var userMap = make(map[int64] *bean.User)
	for index:= 0 ; index < len(users) ; index ++ {
		var user = users[index]
		userMap[user.Id] = user;
	}

	return userMap
}


var USER_SERVICE_INSTANCE = new(UserService)

func  GetUserServiceInstance() *UserService{
	return USER_SERVICE_INSTANCE
}