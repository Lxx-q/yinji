package controllers

import (
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
	"yinji/service"
	"yinji/models/base"
)

type AudioUserHistoryController struct {
	BeegoController
}

/**
	根据用户的id ， 来搜索对应的信息
*/
func ( self *AudioUserHistoryController ) SearchByUserId(){
	var userId , getUserIdErr = self.GetInt64("userId")

	if getUserIdErr != nil {
		self.FailJson( getUserIdErr )
		return
	}

	var page , count = self.GetPageAndCount(15)

	var ormService = db.GetOrmServiceInstance()

	var historyArr []*bean.AudioUserHistroy

	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var qt = o.QueryTable(bean.GetAudioUserHistroyTableName())
		var _ , allErr = qt.Filter("user_id",userId).OrderBy("-modify_time").Limit(count,page*count).All(&historyArr)
		return historyArr , allErr
	})

	if jdbcErr != nil {
		self.FailJson( jdbcErr )
		return
	}

	for _ , history := range historyArr {
		history.Parse()
	}

	self.Json( historyArr )
}

/**
	根据目标的audio和 user 的操作来进行搜索
*/
func ( self *AudioUserHistoryController ) AddCount(){
	//收集对应的信息
	var userId , getUserIdErr = self.GetInt64("userId");

	if getUserIdErr != nil {
		self.FailJson( getUserIdErr )
		return
	}

	var audioId , getAudioIdErr = self.GetInt64("audioId")
	if getAudioIdErr != nil{
		self.FailJson( getAudioIdErr )
		return
	}

	var count , getCountErr = self.GetInt64("count")

	if getCountErr != nil {
		count = 1
	}

	var lastTime , getLastTimeErr = self.GetInt("lastTime")

	if getLastTimeErr != nil {
		lastTime = 0
	}

	var ormService = db.GetOrmServiceInstance()

	var audioUserHistory = service.GetAudioUserHistoryServiceInstance()
	var history *bean.AudioUserHistroy

	var _ , transacErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		var findErr error
		history ,findErr = audioUserHistory.FindOrNew(o , userId , audioId )
		if findErr != nil {
			return history , findErr
		}
		history.BrowseAllCount = history.BrowseAllCount + count;
		history.LastTime = lastTime
		history.Refresh()
		o.Update(history)

		return history , findErr
	})

	if transacErr != nil {
		self.FailJson( transacErr )
		return
	}

	//添加记录
	var dashboardService = service.GetDashboardServiceInstance()
	var dashboardBase = base.NewDashboardBase()
	dashboardBase.BrowseCount = count
	dashboardService.AddCount(audioId,userId,dashboardBase)

	self.Json( history )


}
