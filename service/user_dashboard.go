package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
	"yinji/models/base"
	"yinji/service/db"
)

type UserDashboardService struct {

}

/**
	真正意义上使用jdbc的方法
 */
func ( self *UserDashboardService ) AddDashboardCount( id int64 , dashboardBase *base.DashboardBase ){
	var ormService = db.GetOrmServiceInstance()
	ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		return self.UpdateDashboardCount(o , id , dashboardBase)
	});

}

/**
	数据库上对应的 修改数据的方法
 */
func ( self *UserDashboardService ) UpdateDashboardCount( o orm.Ormer , id int64 , dashboardBase *base.DashboardBase) (*bean.UserDashboard , error){

	/**
		现根据对应的信息 ， 来获取对应的信息
	 */

	var dashboard , findErr = self.FindOrNew(o , id)

	if findErr != nil{
		//倘若没有则开始添加

	}

	dashboard.Add(dashboardBase)

	var _ , updateErr = o.Update(dashboard)
	return dashboard , updateErr

}

func ( self *UserDashboardService ) FindOrNew( o orm.Ormer , id int64) ( *bean.UserDashboard , error) {
	var dashboard , findErr = self.FindById(o , id );
	if findErr != nil {
		dashboard , findErr = self.New(o,dashboard)
	}
	return dashboard , findErr
}

/**
	根据id ， 来生成对应的信息
 */
func ( self *UserDashboardService ) FindById ( o orm.Ormer , id int64 ) ( *bean.UserDashboard , error ){
	var dashboard = &bean.UserDashboard{}
	dashboard.Id = id
	var readErr = o.Read( dashboard )
	return dashboard , readErr
}

func ( self *UserDashboardService ) New( o orm.Ormer , dashboard *bean.UserDashboard ) ( *bean.UserDashboard , error ) {
	dashboard.DashboardBase.New()
	var _ , insertErr = o.Insert( dashboard )
	return dashboard , insertErr
}

func ( self *UserDashboardService ) NewById( o orm.Ormer , id int64 ) ( *bean.UserDashboard , error ) {
	var dashboard = &bean.UserDashboard{}
	dashboard.Id = id
	return self.New( o ,dashboard )
}

var USER_DASHBOARD_SERVICE_INSTANCE = &UserDashboardService{}

func GetUserDashboardService() *UserDashboardService {
	return USER_DASHBOARD_SERVICE_INSTANCE
}
