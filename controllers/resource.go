package controllers

import (
	"yinji/service/db"
	"yinji/service"
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
	"yinji/config"
)

type ResourceController struct {
	BeegoController
}

/**
	根据对应的path , 来直接获取对应的信息
*/
func ( self *ResourceController ) ResourcePath(){
	var path = self.GetString("path")

	if path == "" {
		path = config.RESOURCE_IMAGE_404
	}

	var urlPath = "/" + config.SERVER_NAME + "/" + path
	self.Redirect( urlPath , 302 )
}

/**
	根据 audio id 可以直接找到对应的 audio 音频资源信息
*/
func ( self *ResourceController ) Audio() {
	var id , getIdErr = self.GetInt64("id")

	if getIdErr != nil {
		self.FailJson( getIdErr )
		return
	}

	var ormService = db.GetOrmServiceInstance()
	var audioService = service.GetAudioServiceInstance()
	var audio *bean.Audio
	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var findErr error
		audio , findErr = audioService.FindById( o , id )
		return audio , findErr
	})

	if jdbcErr != nil {
		self.FailJson( jdbcErr )
		return
	}

	var audioPath = audio.Url

	var urlPath = "/" + config.SERVER_NAME + "/" + audioPath

	self.Redirect( urlPath , 302 )

}

/**
	根据audio 来直接获取对应的图片资源
 */
func ( self *ResourceController ) ImageAudio(){
	var id , getIdErr = self.GetInt64("id")

	if getIdErr != nil {
		self.FailJson( getIdErr )
		return
	}

	var ormService = db.GetOrmServiceInstance()
	var audioService = service.GetAudioServiceInstance()
	var audio *bean.Audio
	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var findErr error
		audio , findErr = audioService.FindById( o , id )
		return audio , findErr
	})

	if jdbcErr != nil {
		self.FailJson( jdbcErr )
		return
	}

	var image = audio.Image

	if image == "" {
		image = config.RESOURCE_IMAGE_404
	}

	var urlPath = "/" + config.SERVER_NAME + "/" + image

	self.Redirect( urlPath , 302 )
}

/**
	根据目标user 的id 来直接转化数据

 */
func ( self *ResourceController ) ImageUser() {

	//暂时默认的操作
	//全部返回 404 图片
	/*
	var image = config.RESOURCE_IMAGE_404

	var urlPath = "/" + config.SERVER_NAME + "/" + image

	self.Redirect( urlPath , 302 )
	*/


	var id , getIdErr = self.GetInt64("id")

	if getIdErr != nil {
		self.FailJson( getIdErr )
		return
	}

	var ormService = db.GetOrmServiceInstance()
	var userService = service.GetUserServiceInstance()
	var user *bean.User
	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var findErr error
		user , findErr = userService.FindUserById( o , id )
		return user , findErr
	})

	if jdbcErr != nil {
		self.FailJson( jdbcErr )
		return
	}

	var image = user.Image

	if image == "" {
		image = config.RESOURCE_IMAGE_404
	}

	var urlPath = "/" + config.SERVER_NAME + "/" + image

	self.Redirect( urlPath , 302 )



}

