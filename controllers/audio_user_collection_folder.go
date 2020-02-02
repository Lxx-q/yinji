package controllers

import (
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
	"yinji/service"
	"yinji/models/bean"
)

type AudioCollectionFolderController struct {
	BeegoController
}

/**
	根据用户的id 来获取对应的信息体
 */
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

/**
	新建对应的收藏文件夹
*/
 func ( self *AudioCollectionFolderController ) AddCollectionFolder(){

	var userId , getUserIdErr = self.GetInt64("userId")

	if getUserIdErr != nil {
		self.FailJson( getUserIdErr )
		return
	}

 	//先收集对应的数据信息
 	var name = self.GetString("name")

 	var introduction = self.GetString("introduction")

	//下面将对应的信息体存入对应的 bean  结构体
	var folder = bean.AudioCollectionFolder{}

	folder.Name = name
	folder.Introduction = introduction
	folder.UserId = userId
	folder.New()

	//下面我们准备开始进行对应的数据库操作
	var ormService = db.GetOrmServiceInstance()

	var _ , tranErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		//下面我们进行插入操作
		return o.Insert( &folder )
	})

	if tranErr != nil {
		self.FailJson(tranErr)
		return
	}

	self.Json( folder )

 }

/**
   更新目标收藏夹的信息
*/
func ( self *AudioCollectionFolderController ) UpdateCollectionFolder(){

	var id , getIdErr = self.GetInt64("id")

	if getIdErr != nil {
		self.FailJson( getIdErr )
		return
	}

	//先收集对应的数据信息
	var name = self.GetString("name")

	var introduction = self.GetString("introduction")

	//下面将对应的信息体存入对应的 bean  结构体
	var folder = bean.AudioCollectionFolder{}

	folder.Id = id

	//下面我们准备开始进行对应的数据库操作
	var ormService = db.GetOrmServiceInstance()

	var _ , tranErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		//先根据对应的 id 值 来获取其他的信息 ， 然后更新对应的信息
		o.Read(&folder)
		folder.Name = name
		folder.Introduction = introduction
		folder.Refresh()
		//下面我们进行插入操作
		return o.Update(&folder)
	})

	folder.Parse()

	if tranErr != nil {
		self.FailJson(tranErr)
		return
	}

	self.Json( folder )
}

/**
	根据对应的收藏夹的id ， 删除对应的那一行的收藏夹信息
*/
func ( self *AudioCollectionFolderController ) DeleteCollectionFolder(){

	//先收集对应的信息
	var id , getIdErr = self.GetInt64("id")

	if getIdErr != nil {
		self.FailJson( getIdErr )
		return
	}

	var collectionFolder = bean.AudioCollectionFolder{}

	collectionFolder.Id = id

	var ormService = db.GetOrmServiceInstance()

	var _ , transacErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		return o.Delete(&collectionFolder)
	})

	if transacErr != nil {
		self.FailJson( transacErr )
		return
	}

	//输出最后的结果
	self.Json( collectionFolder )

}

