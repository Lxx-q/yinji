package controllers

import (
	"yinji/service/db"
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
)

type AudioFastController struct {
	BeegoController
}

func ( self *AudioFastController ) SearchArr() {
	var page , count = self.GetPageAndCount(15)

	var ormService = db.GetOrmServiceInstance()
	var audioFastArr []*bean.AudioFast
	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var qt = o.QueryTable(bean.GetAudioFastTableName())
		var _, allErr = qt.OrderBy("-create_time").Limit(count , page*count).All(&audioFastArr)
		return audioFastArr , allErr
	})

	if jdbcErr != nil {
		self.FailJson(jdbcErr)
		return
	}

	for _ , audioFast := range audioFastArr {
		audioFast.Parse()
	}

	self.Json( audioFastArr )
}

/**
 	新建对应的信息 audio_fast
	对应的参数
	audioId:音频的id
	userId:用户的id
	start:  开始时间(按照秒)
	end: 结束书剑( 按照秒)

	//输出对应的audioFast的信息
*/
func ( self *AudioFastController ) NewAudioFast(){
	//收集对应的信息
	var audioId , getAudioIdErr  = self.GetInt64("audioId")

	if getAudioIdErr != nil {
		self.FailJson( getAudioIdErr )
		return
	}

	var userId , getUserIdErr   = self.GetInt64("userId")

	if getUserIdErr != nil {
		self.FailJson( getUserIdErr )
		return
	}

	var start , getStartErr   = self.GetInt("start")

	if getStartErr != nil {
		self.FailJson( getStartErr )
		return
	}

	var end , getEndErr   = self.GetInt("end")

	if getEndErr != nil {
		self.FailJson( getEndErr )
		return
	}

	var introduction  = self.GetString("introduction")

	//下面设置对应的信息

	var audioFast = bean.AudioFast{}

	audioFast.New()
	audioFast.UserId = userId
	audioFast.AudioId = audioId
	audioFast.StartTime = start
	audioFast.EndTime = end
	audioFast.Introduction = introduction

	var ormService = db.GetOrmServiceInstance()

	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var _  , insertErr = o.Insert(&audioFast)
		return audioFast , insertErr
	})

	if jdbcErr != nil {
		self.FailJson(jdbcErr)
		return
	}

	self.Json( audioFast )
}