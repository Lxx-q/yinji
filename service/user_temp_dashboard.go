package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
	"yinji/models/base"
	"yinji/service/db"
	"yinji/utils"
)

type UserTempDashboardService struct {

}

/**
	真正意义上使用jdbc的方法
 */
func ( self *UserTempDashboardService ) AddDashboardCount( id int64 , dashboardBase *base.DashboardBase ){
	var ormService = db.GetOrmServiceInstance()
	ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		return self.UpdateDashboardCount(o , id , dashboardBase)
	});
}
/**
	数据库上对应的 修改数据的方法
 */
func ( self *UserTempDashboardService ) UpdateDashboardCount( o orm.Ormer , id int64 , dashboardBase *base.DashboardBase) (*bean.UserTempDashboard , error){

	/**
		现根据对应的信息 ， 来获取对应的信息
	 */

	var dashboard , findErr = self.FindOrNew(o , id)

	if findErr != nil{
		//倘若没有则开始添加
		return nil , findErr
	}

	dashboard.Add(dashboardBase)

	var _ , updateErr = o.Update(dashboard)
	return dashboard , updateErr

}
/**

 */
func ( self *UserTempDashboardService ) FindOrNew( o orm.Ormer , id int64) ( *bean.UserTempDashboard , error) {
	var dashboard , findErr = self.FindById(o , id );
	if findErr != nil {
		dashboard , findErr = self.New(o,dashboard , id )
	}
	return dashboard , findErr
}
/**
	根据id ， 来生成对应的信息
 */
func ( self *UserTempDashboardService ) FindById ( o orm.Ormer , id int64 ) ( *bean.UserTempDashboard , error ){
	var dashboard = &bean.UserTempDashboard{}
	dashboard.UserId = id
	var today = utils.Today()
	dashboard.WriteDate = *today
	var readErr = o.Read( dashboard , "userId" , "writeDate")
	return dashboard , readErr
}

func ( self *UserTempDashboardService ) New( o orm.Ormer , dashboard *bean.UserTempDashboard , id int64) ( *bean.UserTempDashboard , error ) {
	dashboard.UserId = id
	dashboard.New()
	var _ , insertErr = o.Insert( dashboard )
	return dashboard , insertErr
}

func ( self *UserTempDashboardService ) NewById( o orm.Ormer , id int64 ) ( *bean.UserTempDashboard , error ) {
	var dashboard = &bean.UserTempDashboard{}
	dashboard.Id = id
	return self.New( o ,dashboard , id)
}

var USER_TEMP_DASHBOARD_SERVICE_INSTANCE = &UserTempDashboardService{}

func GetUserTempDashboardServiceInstance() *UserTempDashboardService {
	return USER_TEMP_DASHBOARD_SERVICE_INSTANCE
}