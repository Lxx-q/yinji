package controllers

import (
	"github.com/astaxie/beego/orm"
	"yinji/service/db"
	"yinji/service"
	"yinji/models/bean"
)

type DashboardContrlller struct {
	BeegoController
}

func ( self *DashboardContrlller ) FindByAudioId( ){
	var audioId , getAudioIdErr = self.GetInt64("audioId"  )

	if getAudioIdErr != nil {
		self.FailJson(getAudioIdErr)
		return
	}

	var dashboradService = service.GetDashboardServiceInstance()

	var dashboradList []*bean.AudioUserDashboard
	var dashboradMap map[int]*bean.AudioUserDashboard

	var ormService = db.GetOrmServiceInstance()
	ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		_dashboradList , _dashboradErr := dashboradService.FindDashboardByAudio( o , audioId )
		dashboradList = _dashboradList
		return _dashboradList ,_dashboradErr
	})

	dashboradMap = dashboradService.ToMap(dashboradList)

	self.Json(dashboradMap)
}
