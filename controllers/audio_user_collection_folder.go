package controllers

import (
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
	"yinji/service"
)

type AudioCollectionFolderController struct {
	BeegoController
}

//根据对应的 用户信息 来进行搜索
func ( self *AudioCollectionFolderController ) AllByUserId(){
	//获取对应的 user 的 信息

	var userId , getUserIdErr = self.GetInt64("userId")

	if getUserIdErr != nil {
		self.FailJson(getUserIdErr)
		return
	}
	var ormService = db.GetOrmServiceInstance()
	var folderService = service.GetCollectionFolderServiceInstance()
	var folderList , tranErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		//对应的信息
		return folderService.AllByUserId(o , userId )
	})

	if tranErr != nil {
		self.FailJson( tranErr )
	}

	self.Json( folderList )

}


