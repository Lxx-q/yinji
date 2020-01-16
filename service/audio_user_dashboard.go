package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
)

type DashboardService struct {

}

func ( self *DashboardService) ToMap(dashboards []*bean.AudioUserDashboard)  map[int]*bean.AudioUserDashboard {
	var dashboardLen = len( dashboards )

	var dashboardMap =  make(map[int]*bean.AudioUserDashboard)
	for index:= 0 ; index < dashboardLen ; index++ {
		var dashboradItem = dashboards[index]
		dashboradItem.Parse()
		dashboardMap[dashboradItem.DashboardType] = dashboradItem
	}

	return dashboardMap
}

//获取对应的信息 ， 最后输出对应的信息
func ( self *DashboardService ) FindDashboardByAudio( o orm.Ormer , audioId int64 ) ([]*bean.AudioUserDashboard , error ){
	var dashboards []*bean.AudioUserDashboard
	var qt = o.QueryTable( bean.GetAudioUserDashboardTableName() );
	qt = Limit(qt)
	var _ , allErr = qt.Filter("AudioId",audioId).All( &dashboards )
	//输出信息
	return dashboards , allErr
}

//添加对应的信息量
func ( self *DashboardService ) AddCount( o orm.Ormer , audioId int64 , dashboradType int ) ( *bean.AudioUserDashboard , error ) {
	var dashborad = &bean.AudioUserDashboard{}
	var readErr =  o.Read( dashborad , "AudioId","DashboradType")

	if readErr != nil {
		return nil , readErr
	}

	//原本的 点击数量
	var count = dashborad.Count

	dashborad.Count = count + 1

	var updateSql = "UPDATE audio_user_dashborad aud SET aud.count = ? WHERE aud.count = ? AND aud.id = ? "
	var _ , err = o.Raw( updateSql ,dashborad.Count , count , dashborad.Id ).Exec()

	if err != nil {
		return nil , err
	}

	return dashborad , nil
}

//对应的数据
var DASHBOARD_SERVICE_INSTANCE *DashboardService = &DashboardService{}

func GetDashboardServiceInstance() *DashboardService{
	return DASHBOARD_SERVICE_INSTANCE
}