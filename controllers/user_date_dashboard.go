package controllers

import (
	"yinji/utils"
	"yinji/service/db"
	"yinji/service"
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
	"time"
)

type UserDateDashboardController struct {
	BeegoController
}

/**
	根据分化的时间来进行获取特定的时间
 */
func ( self *UserDateDashboardController ) SearchByAudioId(){
	var userId , getUserIdErr = self.GetInt64("userId")

	if getUserIdErr != nil{
		self.Json( getUserIdErr )
		return
	}
	var start , getStartErr = self.GetInt64("start")

	if getStartErr != nil {
		self.FailJson( getStartErr )
		return
	}

	var end , getEndErr = self.GetInt64("end")

	var startDate = utils.Unix( start )
	//获取对应的时间
	var endDate *time.Time = nil

	if getEndErr == nil {
		endDate = utils.Unix( end )
	}

	var ormService = db.GetOrmServiceInstance()
	var userDateDashboardService = service.GetUserDateDashboardServiceInstance()

	var userDateDashboard []*bean.UserDateDashboard
	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var searchErr  error
		userDateDashboard , searchErr = userDateDashboardService.SearchByAudioId(o , userId , startDate , endDate )
		if searchErr != nil {
			return userDateDashboard , searchErr
		}
		return userDateDashboard , nil
	})

	if jdbcErr != nil {
		self.FailJson( jdbcErr )
		return
	}

	self.Json( userDateDashboardService.ToMap(&userDateDashboard ) )

}


/**
	根据分化的时间来进行获取特定的时间
 */
func ( self *UserDateDashboardController ) SearchByAudioIdByTime(){
	var audioId , getAudioErr = self.GetInt64("audioId")

	if getAudioErr != nil{
		self.Json( getAudioErr )
		return
	}
	//获取起始时间 ，
	var startYear , getStartYearErr = self.GetInt("startYear")

	if getStartYearErr != nil {
		self.FailJson( getStartYearErr )
		return
	}

	var startMonth , getStartMonthErr = self.GetInt("startMonth")

	if getStartMonthErr != nil {
		self.FailJson( getStartMonthErr )
		return
	}

	var startDay , getStartDayErr = self.GetInt("startDay")

	if getStartDayErr != nil {
		self.FailJson( getStartDayErr )
		return
	}


	var endYear , getEndYearErr = self.GetInt("endYear")

	var endMonth , getEndMonthErr = self.GetInt("endMonth")

	var endDay , getEndDayErr = self.GetInt("endDay")

	var startDate =  utils.Date( startYear , startMonth , startDay )
	//获取对应的时间

	var endDate *time.Time
	if getEndYearErr != nil  || getEndMonthErr != nil || getEndDayErr != nil {
		endDate = nil
	}else{
		endDate = utils.Date( endYear , endMonth , endDay )
	}

	var ormService = db.GetOrmServiceInstance()
	var userDateDashboardService = service.GetUserDateDashboardServiceInstance()

	var userDateDashboard []*bean.UserDateDashboard
	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var searchErr  error
		userDateDashboard , searchErr = userDateDashboardService.SearchByAudioId(o , audioId , startDate , endDate )
		if searchErr != nil {
			return userDateDashboard , searchErr
		}
		return userDateDashboard , nil
	})

	if jdbcErr != nil {
		self.FailJson( jdbcErr )
		return
	}

	self.Json( userDateDashboardService.ToMap(&userDateDashboard ) )

}