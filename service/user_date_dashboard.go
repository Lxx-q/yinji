package service

import (
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
	"time"
	"yinji/models/base"
	"yinji/service/db"
	"yinji/utils"
)

type UserDateDashboardService struct {

}

/**
	将数组转化为 map
*/
func ( self *UserDateDashboardService ) ToMap( array *[]*bean.UserDateDashboard ) ( *map[int64] *bean.UserDateDashboard ) {
	var _map = make( map[int64] *bean.UserDateDashboard )
	for _ , dateDashboard := range *array {
		_map[dateDashboard.WriteDate.Unix() * 1000 ] = dateDashboard
	}
	return &_map
}

/**
	根据对应的 audioId 来查询对应的信息
 */
func ( self *UserDateDashboardService ) SearchByAudioId( o orm.Ormer , audioId int64 , startDate *time.Time , endDate *time.Time) ( []*bean.UserDateDashboard , error ){
	var userDateDashboard []*bean.UserDateDashboard
	var qt = o.QueryTable(bean.GetUserDateDashboardTableName())
	qt = qt.Filter("userId" , audioId).Filter("writeDate__gte",*startDate).OrderBy("writeDate" )

	if endDate != nil {
		qt = qt.Filter("writeDate__lte" , *endDate)
	}

	var _ , allErr = qt.All(&userDateDashboard)
	return userDateDashboard , allErr
}

/**
	新建方法的默认操作
 */
func ( self *UserDateDashboardService ) NewByTemp( o orm.Ormer , tempDashboard *bean.UserDateDashboard) ( *bean.UserDateDashboard , error ) {
	tempDashboard.New()
	tempDashboard.DashboardBase = tempDashboard.DashboardBase
	var _ , insertErr = o.Insert( tempDashboard )
	return tempDashboard , insertErr
}

/**
	根据对应的userId 来直接输出
 */
func ( self *UserDateDashboardService ) NewByUserId( o orm.Ormer , userId int64 ) ( *bean.UserDateDashboard , error ){
	var dashboard = &bean.UserDateDashboard{}
	dashboard.UserId = userId
	return self.NewByTemp( o , dashboard)
}

func( self *UserDateDashboardService ) FindOrNew( o orm.Ormer , userId int64 ) ( *bean.UserDateDashboard , error ){
	var dashboard , findErr = self.FindById( o , userId )
	if findErr != nil {
		dashboard , findErr = self.NewByUserId( o , userId )
	}
	return dashboard , findErr
}

/**
	根据对应的userId 来搜索对应的信息
 */
func ( self *UserDateDashboardService ) FindById( o orm.Ormer , userId int64 )( *bean.UserDateDashboard , error ){
	var dashboard = &bean.UserDateDashboard{}
	dashboard.UserId = userId
	dashboard.WriteDate = *utils.Today()
	var readErr = o.Read( dashboard , "userId","writeDate")
	return dashboard , readErr
}


/**
	数据库上对应的 修改数据的方法
 */
func ( self *UserDateDashboardService ) UpdateDashboardCount( o orm.Ormer , id int64 , dashboardBase *base.DashboardBase) ( *bean.UserDateDashboard , error ){

	/**
		现根据对应的信息 ， 来获取对应的信息
	 */

	var dashboard , findErr = self.FindOrNew(o , id)

	if findErr != nil{
		//倘若没有则开始添加
		return nil, findErr
	}

	dashboard.Add(dashboardBase)

	var _ , updateErr = o.Update(dashboard)
	return dashboard , updateErr

}

/**
	真正意义上使用jdbc的方法
 */
func ( self *UserDateDashboardService ) AddDashboardCount( id int64 , dashboardBase *base.DashboardBase ){
	var ormService = db.GetOrmServiceInstance()
	ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		return self.UpdateDashboardCount(o , id , dashboardBase)
	});
}


var USER_DATE_DASHBORAD_SERVICE_INSTANCE = &UserDateDashboardService{}


func GetUserDateDashboardServiceInstance() *UserDateDashboardService{
	return USER_DATE_DASHBORAD_SERVICE_INSTANCE
}
