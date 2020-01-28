package service

import (
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
)

type DashboardService struct {

}

const (
	AUDIO_DASHBOARD_C0LLECTION = 1
	AUDIO_DASHBOARD_FAVORITE = 2
	AUDIO_DASHBOARD_FORWARD = 3
)

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

//封装上的关注事件
func ( self *DashboardService ) AddDashboradCountEvent( o orm.Ormer , dashboradType int , audioId int64   ) ( *bean.AudioUserDashboard , error){
	//先获取对应的信息 ， 如无的话 ， 则添加一条新信息
	var dashborad , findErr = self.FindDashboradByAudioAndType( o , dashboradType , audioId )

	if findErr != nil {
		//倘若对应的是未查询

	}

	//倘若出现错误 ， 这说明该类并没有对应的信息行。我们便可将其删除
	dashborad , newErr := self.NewDashboard( o ,dashboradType , audioId)

	if newErr != nil {
		//出现错误
		return nil , newErr
	}

	dashborad , updateErr :=self.AddDashboradCount( o , dashborad )

	return dashborad , updateErr

}

/**
	下面便是基础的数据操作方法
 */

//获取对应的信息 ， 最后输出对应的信息
func ( self *DashboardService ) FindDashboardByAudio( o orm.Ormer , audioId int64 ) ([]*bean.AudioUserDashboard , error ){
	var dashboards []*bean.AudioUserDashboard
	var qt = o.QueryTable( bean.GetAudioUserDashboardTableName() );
	qt = Limit(qt)
	var _ , allErr = qt.Filter("AudioId",audioId).All( &dashboards )
	//输出信息
	return dashboards , allErr
}

func ( self *DashboardService ) FindDashboradByAudioAndType ( o orm.Ormer , dashboardType int  , audioId int64) (*bean.AudioUserDashboard , error) {
	var dashboread = &bean.AudioUserDashboard{}
	dashboread.DashboardType = dashboardType
	dashboread.AudioId = audioId
	//根据对应的信息来获取信息
	var readErr = o.Read(dashboread , "DashboardType","AudioId")
	return dashboread , readErr
}

//性插入数据
func ( self *DashboardService ) NewDashboard( o orm.Ormer , dashboardType int  , audioId int64 ) (*bean.AudioUserDashboard , error) {
	var dashboard = &bean.AudioUserDashboard{}
	dashboard.New()
	dashboard.AudioId = audioId
	dashboard.DashboardType = dashboardType
	var _ , insertErr = o.Insert(dashboard)

	return dashboard , insertErr
}

//添加对应的信息量
func ( self *DashboardService ) AddDashboradCount( o orm.Ormer , dashborad *bean.AudioUserDashboard ) ( *bean.AudioUserDashboard , error ) {

	//原本的 点击数量
	var count = dashborad.Count

	dashborad.Count = count + 1

	var updateSql = "UPDATE audio_user_dashborad aud SET aud.count = ? WHERE aud.count = ? AND aud.id = ? ; "
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