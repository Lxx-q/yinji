package controllers

import (
	"yinji/utils"
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
	"yinji/models/base"
	"errors"
)

type DashboardController struct {
	BeegoController
}

/**
	根据特定的用户id ， 获取该目标今天（一个周期 temp）获取的信息总数
*/
func ( self  *DashboardContrlller ) TempFindUserId(){
	var userId , getUserIdErr = self.GetInt64("userId")

	if getUserIdErr != nil {
		self.FailJson( getUserIdErr )
		return
	}

	var today = utils.Today()

	var sql = "SELECT SUM(atd.browse_count) as browse_count , SUM(atd.forward_count) as forward_count , SUM(atd.love_count) as love_count , SUM(atd.collection_count) as collection_count , SUM(atd.comment_count) as comment_count FROM audio_temp_dashboard atd WHERE atd.write_date= ? AND atd.audio_id IN ( SELECT a.id FROM audio a WHERE a.user_id= ? );"

	var ormService = db.GetOrmServiceInstance()
	var row  []base.DashboardBase
	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var _ , listErr = o.Raw(sql,today,userId).QueryRows(&row)
		return nil , listErr
	})

	if jdbcErr != nil {
		self.FailJson(jdbcErr)
	}

	if( len(row) != 1 ){
		self.FailJson(errors.New("this is a err"))
		return
	}
	self.Json( row[0] )

}

/**
	获取一个用户的所有的信息
*/
func ( self  *DashboardContrlller ) AudioDashboardFindUserId(){
	var userId , getUserIdErr = self.GetInt64("userId")

	if getUserIdErr != nil {
		self.FailJson( getUserIdErr )
		return
	}

	var sql = "SELECT  SUM(ad.browse_count) as browse_count , SUM(ad.forward_count) as forward_count , SUM(ad.love_count) as love_count , SUM(ad.collection_count) as collection_count , SUM(ad.comment_count) as comment_count  FROM audio_dashboard ad WHERE ad.id IN ( SELECT a.id FROM audio a WHERE a.user_id = ? );"

	var ormService = db.GetOrmServiceInstance()
	var row  []base.DashboardBase
	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var _ , listErr = o.Raw(sql,userId).QueryRows(&row)
		return nil , listErr
	})

	if jdbcErr != nil {
		self.FailJson(jdbcErr)
		return
	}

	if( len(row) != 1 ){
		self.FailJson(errors.New("this is a err"))
		return
	}
	self.Json( row[0] )

}