package controllers

import (
	"yinji/service/db"
	"yinji/service"
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
	"yinji/config"
	"strconv"
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

/**
	根据对应的hostResource表来进行获取数据
*/

func ( self *ResourceController ) HostResource(){
	var id , getIdErr = self.GetInt64("id")
	var image string = config.RESOURCE_IMAGE_404

	if getIdErr != nil {
		var urlPath = "/" + config.SERVER_NAME + "/" + image
		self.Redirect(urlPath,302)
		return
	}

	var hostResource = bean.HostResource{}
	hostResource.Id = id
	var ormService = db.GetOrmServiceInstance()
	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var readErr  = o.Read(&hostResource)
		return hostResource , readErr
	})

	if jdbcErr != nil {
		var urlPath = "/" + config.SERVER_NAME + "/" + image
		self.Redirect(urlPath,302)
		return
	}
	image = hostResource.Path
	var urlPath = "/" + config.SERVER_NAME + "/" + image
	self.Redirect(urlPath,302)
}

/**

 */
func ( self *ResourceController) ResourceById(){

	var id , getIdErr = self.GetInt64("id")
	var image string = config.RESOURCE_IMAGE_404

	if getIdErr != nil {
		var urlPath = "/" + config.SERVER_NAME + "/" + image
		self.Redirect(urlPath,302)
		return
	}

	var resource = bean.Resource{}
	resource.Id = id
	var hostResource = bean.HostResource{}
	var ormService = db.GetOrmServiceInstance()
	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var readResourceErr  = o.Read(&resource)
		if readResourceErr != nil {
			return nil , readResourceErr
		}
		hostResource.Id = resource.HostResourceId
		var readHostResourceErr = o.Read(&hostResource)
		return resource, readHostResourceErr
	})

	if jdbcErr != nil {
		var urlPath = "/" + config.SERVER_NAME + "/" + image
		self.Redirect(urlPath,302)
		return
	}
	image = hostResource.Path
	var urlPath = "/" + config.SERVER_NAME + "/" + image
	self.Redirect(urlPath,302)


}

/**
	根据对应的 resourceimage的id 以及类型来获取对应的信息
 */
func ( self *ResourceController ) ResourceImage(){

	var id , getIdErr = self.GetInt64("id")
	var image =  config.RESOURCE_IMAGE_404

	if getIdErr != nil {
		var urlPath = "/" + config.SERVER_NAME + "/" + image
		self.Redirect(urlPath,302)
		return
	}

	var imageType , getImageTypeErr = self.GetInt("type")

	if getImageTypeErr != nil {
		imageType = bean.RESOURCE_IMAGE_ORIGIN_TYPE
	}

	var ormService = db.GetOrmServiceInstance()
	var resourceImage = bean.ResourceImage{}
	resourceImage.Id = id
	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var readErr = o.Read(&resourceImage)
		return nil , readErr
	})

	if jdbcErr != nil {
		var urlPath = "/" + config.SERVER_NAME + "/" + image
		self.Redirect(urlPath,302)
		return
	}

	var resource_id = resourceImage.OriginResourceId

	if imageType == bean.RESOURCE_IMAGE_THUMB_TYPE {
		if resourceImage.ThumbResourceId != 0 {
			resource_id = resourceImage.ThumbResourceId
		}
	}

	var resourceId = strconv.FormatInt(resource_id , 10)
	var url = "/yinji/api/resource/id?id=" + resourceId
	self.Redirect( url , 302 )

}

/**
	根据对应的 resourceAudio 信息来进行查询信息
*/
func ( self *ResourceController ) ResourceAudio(){

	var id , getIdErr = self.GetInt64("id")

	if getIdErr != nil {
		//返回错误的类型
		self.FailJson(getIdErr)
		return
	}

	//由于暂时没有对应的信息，因此并没有设置对应的信息
	/*
	var audioType , getAudioTypeErr = self.GetInt("type")

	if getAudioTypeErr != nil {
		self.FailJson( getAudioTypeErr )
		return
	}
	*/

	var ormService = db.GetOrmServiceInstance()
	var resourceAudio = bean.ResourceAudio{}
	resourceAudio.Id = id
	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var readErr = o.Read(&resourceAudio)
		//得到对应的信
		return nil , readErr
	})

	if jdbcErr != nil {
		self.FailJson( jdbcErr )
		return
	}

	var resourceId = resourceAudio.OriginResourceId //暂时默认为 origin resource 的对应信息

	var url = "/yinji/api/resource/id?id=" + strconv.FormatInt(resourceId,10)
	self.Redirect( url , 302 )

}