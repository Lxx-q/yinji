package service

import "yinji/models/base"

type DashboardService struct {
	AudioDashboard *map[int64]*base.DashboardBase //对应audioDashboard
	AudioTempDashboard *map[int64]*base.DashboardBase
	UserDashboard *map[int64]*base.DashboardBase
	UserTempDashboard *map[int64]*base.DashboardBase
}

func ( self *DashboardService ) Start() *DashboardService{
	self.AudioDashboard = new(map[int64]*base.DashboardBase)
	self.AudioTempDashboard = new(map[int64]*base.DashboardBase)
	self.UserDashboard = new(map[int64]*base.DashboardBase)
	self.UserTempDashboard = new(map[int64]*base.DashboardBase)
	return self
}
/**
	内部方法 ， 从map中获取对应的dashboard
 */
func ( self *DashboardService ) getDashboardFMap( dashboardMap map[int64]*base.DashboardBase ,  id int64 ) *base.DashboardBase {
	var dashboard = dashboardMap[id]
	if dashboard == nil {
		//如果对应的 dashboard 为空的话 ， 我们则自动生成一个新的并且添加进去
		dashboard = new(base.DashboardBase)
		dashboard.New()
		dashboardMap[id] = dashboard
	}
	return dashboard
}

/**
	开发的另一个添加数量的方法
*/
func ( self *DashboardService ) AddCountT(audioId int64 , userId int64 , base *base.DashboardBase) bool{

	var audioDashboardBase = self.getDashboardFMap(*self.AudioDashboard , audioId)
	var audioTempDashboardBase = self.getDashboardFMap(*self.AudioTempDashboard , audioId )
	var userDashboardBase = self.getDashboardFMap(*self.UserDashboard , userId)
	var userTempDashboardBase = self.getDashboardFMap(*self.AudioTempDashboard , userId)
	audioDashboardBase.Add(base)
	audioTempDashboardBase.Add(base)
	userDashboardBase.Add(base)
	userTempDashboardBase.Add(base)

	return true
}


/**
	根据对应的信息 ，来根据对应的信息 来进行操作
 */
func( self *DashboardService ) AddCount(audioId int64 , userId int64 , base *base.DashboardBase) bool{
	//var audioDashboardService = GetAudioDashboardServiceInstance()
	var audioTempDashboardService = GetAudioTempDashboardServiceInstance()
	//var userDashboardService = GetUserDashboardService()
	var userTempDashboardService = GetUserTempDashboardServiceInstance()
	//var userDateDashboardService = GetUserDateDashboardServiceInstance()

	//添加对应的信息
	//audioDashboardService.AddDashboardCount(audioId , base)
	audioTempDashboardService.AddDashboardCount(audioId , base)
	//userDashboardService.AddDashboardCount(userId , base)
	userTempDashboardService.AddDashboardCount(userId,base)
	//userDateDashboardService.AddDashboardCount(userId,base)

	return true
}

var sDASHBOARD_SERVICE_INSTANCE = &DashboardService{}

var sDASHBOARD_SERVICE_INSTANCE_INIT = false

func GetDashboardServiceInstance()  *DashboardService {
	if sDASHBOARD_SERVICE_INSTANCE_INIT {
		sDASHBOARD_SERVICE_INSTANCE.Start()
		sDASHBOARD_SERVICE_INSTANCE_INIT = true
	}
	return sDASHBOARD_SERVICE_INSTANCE
}
