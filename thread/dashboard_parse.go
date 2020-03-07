package thread

import (
	"yinji/service"
	"yinji/models/bean"
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
	"yinji/utils"
	"time"
)

//主要的工作便是将对应的 userTempDashboard 转化到 userDateDashboard 之中
func ParseUserTempToUserDate(){

	service.Async(func() {

		for true {

			var ormService = db.GetOrmServiceInstance()
			var userDashboardService = service.GetUserDashboardService()
			var userTempDashboardArr []*bean.UserTempDashboard
			ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
				//先获取不属于今天的日子
				var today = utils.Today()
				var qt = o.QueryTable( bean.GetUserTempDashboardTableName())
				var _ , allErr = qt.Exclude("writeDate",today).All( &userTempDashboardArr)

				var dashboardLen = len(userTempDashboardArr)
				for index := 0; index <  dashboardLen ; index++  {
					//之后将目标
					var tempDashboard = userTempDashboardArr[index]
					// 之后我们将对应的 temp 转化成对应的 DateDashboard
					var dateDashboard = bean.UserDateDashboard{}
					dateDashboard.DashboardBase = tempDashboard.DashboardBase
					dateDashboard.Id = tempDashboard.Id
					dateDashboard.WriteDate = tempDashboard.WriteDate
					dateDashboard.UserId = tempDashboard.UserId

					//之后插入对应的数据
					var _ , insertErr = o.Insert(&dateDashboard)
					if insertErr != nil {
						continue;
					}

					var userId = tempDashboard.UserId

					var _ , updateErr = userDashboardService.UpdateDashboardCount( o , userId , &tempDashboard.DashboardBase)

					if updateErr != nil {
						continue
					}

					//下面将对应的数据插入到对应的userDashboard

					o.Delete(tempDashboard)

				}

				return nil ,  allErr
			})

			//下面開始每隔相對應的時間再運行 , 目前的時間暫定為 3 個小時
			time.Sleep(time.Duration(3) * time.Hour)
		}
	})


	// 之后将所有的数据输出
}


//主要的工作便是将对应的 audioTemp 转化到 audioDateDashboard之中
func ParseAudioTempToUserDate(){

	service.Async(func() {

		for true {

			var ormService = db.GetOrmServiceInstance()
			var audioTempDashboardArr []*bean.AudioTempDashboard
			var audioDashboardService = service.GetAudioDashboardServiceInstance()
			ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
				//先获取不属于今天的日子
				var today = utils.Today()
				var qt = o.QueryTable( bean.GetAudioTempTempDashboard())
				var _ , allErr = qt.Exclude("writeDate",today).All( &audioTempDashboardArr)

				var dashboardLen = len(audioTempDashboardArr)
				for index := 0; index <  dashboardLen ; index++  {
					//之后将目标
					var tempDashboard = audioTempDashboardArr[index]
					// 之后我们将对应的 temp 转化成对应的 DateDashboar
					var audioId = tempDashboard.AudioId
					var _ , updateErr = audioDashboardService.UpdateDashboardCount(o, audioId , &tempDashboard.DashboardBase);

					if( updateErr != nil ){
						//若有错，则直接添加
						continue
					}

					o.Delete(tempDashboard)

				}

				return nil ,  allErr
			})

			//下面開始每隔相對應的時間再運行 , 目前的時間暫定為 3 個小時
			time.Sleep(time.Duration(3) * time.Hour)
		}
	})


	// 之后将所有的数据输出
}