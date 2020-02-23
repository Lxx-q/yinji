package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
	"yinji/models/base"
	"yinji/service/db"
)

type AudioTempDashboardService struct {

}

/**
	根据对应的id 来进行搜索
*/
func ( self *AudioTempDashboardService ) FindById( o orm.Ormer , audioId int64) ( *bean.AudioTempDashboard , error){
	var dashboard = bean.AudioTempDashboard{}
	dashboard.Id = audioId
	var readErr = o.Read( &dashboard)
	return &dashboard , readErr
}

/**
	根据对应的方法 ， 来进行 新建
 */
func ( self *AudioTempDashboardService ) New( o orm.Ormer , audioId int64) ( *bean.AudioTempDashboard , error) {
	var dashboard = &bean.AudioTempDashboard{}
	dashboard.New()
	dashboard.Id = audioId
	dashboard.AudioId = audioId
	var _ , insertErr = o.Insert( dashboard )
	return dashboard , insertErr
}

/**
	新建或者 查询
 */
func ( self *AudioTempDashboardService ) FindOrNew( o orm.Ormer , audioId int64 ) ( *bean.AudioTempDashboard , error)  {
	var dashboard , findErr = self.FindById(o , audioId );
	if findErr != nil {
		dashboard , findErr = self.New(o,audioId)
	}

	return dashboard , findErr
}

/**

*/

/**
	数据库上对应的 修改数据的方法
 */
func ( self *AudioTempDashboardService ) UpdateDashboardCount( o orm.Ormer , audioId int64 , dashboardBase *base.DashboardBase) (*bean.AudioTempDashboard , error){

	/**
		现根据对应的信息 ， 来获取对应的信息
	 */

	var dashboard , findOrNewErr = self.FindOrNew(o , audioId)

	if findOrNewErr != nil{
		//倘若没有则开始添加
		return nil , findOrNewErr
	}

	dashboard.Add(dashboardBase)

	var _ , updateErr = o.Update(dashboard)
	return dashboard , updateErr

}

/**
	真正意义上使用jdbc的方法
 */
func ( self *AudioTempDashboardService ) AddDashboardCount( id int64 , dashboardBase *base.DashboardBase ){
	var ormService = db.GetOrmServiceInstance()
	ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		return self.UpdateDashboardCount(o , id , dashboardBase)
	});

}


var AUDIO_TEMP_DASHBOARD_SERVICE_INSTANCE = &AudioTempDashboardService{}

func GetAudioTempDashboardServiceInstance() *AudioTempDashboardService{
	return AUDIO_TEMP_DASHBOARD_SERVICE_INSTANCE
}