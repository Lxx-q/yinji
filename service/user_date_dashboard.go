package service

import (
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
	"time"
)

type UserDateDashboardService struct {

}

/**
	将数组转化为 map
*/
func ( self *UserDateDashboardService ) ToMap( array *[]*bean.UserDateDashboard ) ( *map[int64] *bean.UserDateDashboard ) {
	var _map = make( map[int64] *bean.UserDateDashboard )
	for _ , dateDashboard := range *array {
		_map[dateDashboard.WriteDate.Unix() * 1000] = dateDashboard
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

func ( self *UserDateDashboardService ) NewByTemp( tempDashboard *bean.UserTempDashboard) {
	var userDateDashboard = &bean.UserDateDashboard{}
	userDateDashboard.New()
	userDateDashboard.DashboardBase = tempDashboard.DashboardBase
}

var USER_DATE_DASHBORAD_SERVICE_INSTANCE = &UserDateDashboardService{}


func GetUserDateDashboardServiceInstance() *UserDateDashboardService{
	return USER_DATE_DASHBORAD_SERVICE_INSTANCE
}
