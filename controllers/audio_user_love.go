package controllers

import (
	"yinji/service"
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
)

type AudioUserLoveController struct {
	BeegoController
}

/**
	根据目标id 查找是否有点赞关系
*/
func ( self *AudioUserLoveController ) FindLove(){

	var userId , getUserIdErr = self.GetInt64("userId")

	if getUserIdErr != nil {
		self.FailJson(getUserIdErr)
		return
	}

	var audioId , getAudioIdErr = self.GetInt64("audioId")

	if getAudioIdErr != nil {
		self.FailJson( getAudioIdErr )
		return
	}

	var audioUserLoveService = service.GetAudioUserLoveServiceInstance()
	var ormService = db.GetOrmServiceInstance()

	var love , findLoveErr =  ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		return audioUserLoveService.FindLove( o , userId , audioId )
	})

	if findLoveErr != nil {
		self.FailJson( findLoveErr )
		return
	}

	self.Json( love )

}

/**
	点赞的接口
*/
func ( self *AudioUserLoveController ) InsertLove(){
	//先收取对应的信息 ， user [ 用户的id ] ， audio [ 音频的 id ]
	var userId , getUserIdErr = self.GetInt64("userId")

	if getUserIdErr != nil {
		self.FailJson(getUserIdErr)
		return
	}

	var audioId , getAudioIdErr = self.GetInt64("audioId")

	if getAudioIdErr != nil {
		self.FailJson( getAudioIdErr )
		return
	}

	var audioUserLoveService = service.GetAudioUserLoveServiceInstance()
	var ormService = db.GetOrmServiceInstance()

	var love , insertLoveErr =  ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		return audioUserLoveService.InsertLove( o , userId , audioId )
	})

	if insertLoveErr != nil {
		self.FailJson( insertLoveErr )
		return
	}

	self.Json( love )

}

/**
	趋向点赞关系
 */
func ( self *AudioUserLoveController ) DeleteLove(){

	//先收取对应的信息 ， user [ 用户的id ] ， audio [ 音频的 id ]
	var userId , getUserIdErr = self.GetInt64("userId")

	if getUserIdErr != nil {
		self.FailJson(getUserIdErr)
		return
	}

	var audioId , getAudioIdErr = self.GetInt64("audioId")

	if getAudioIdErr != nil {
		self.FailJson( getAudioIdErr )
		return
	}

	var audioUserLoveService = service.GetAudioUserLoveServiceInstance()
	var ormService = db.GetOrmServiceInstance()


	var love , deleteLoveErr =  ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		return audioUserLoveService.DeleteLove( o , userId , audioId )
	})

	if deleteLoveErr != nil {
		self.FailJson(deleteLoveErr)
		return
	}

	self.Json( love )

}