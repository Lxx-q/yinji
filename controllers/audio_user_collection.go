package controllers

import (
	"errors"
	"yinji/service"
	"yinji/models/bean"
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type CollectionController struct {
	BeegoController
}

func ( self *CollectionController ) InsertCollection() {
	//获取对应的属性
	var userId, getUserIdErr= self.GetInt64("userId")

	if getUserIdErr != nil {
		self.FailJson(getUserIdErr)
		return
	}

	var audioId, getAudioIdErr= self.GetInt64("audioId")

	if getAudioIdErr != nil {
		self.FailJson( getAudioIdErr )
		return
	}

	//将收集而来的信息全部进行收入
	var collection = &bean.AudioUserCollection{}
	//生成对应的信息
	collection.New()

	collection.UserId = userId
	collection.AudioId = audioId

	var collectionService = service.GetCollectonServiceInstance()

	var ormService = db.GetOrmServiceInstance()

	var _ , insertErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		return collectionService.Insert( o , collection)
	})

	if insertErr != nil {
		self.FailJson( insertErr )
		return
	}
	//输出插入后的整合的信息
	self.Json( collection )
}

func ( self *CollectionController ) DeleteCollection(){
	//这个的话 ， 只需要获取对应的id ， 便可以进行对应的操作了
	var collectionId , getCollectionId = self.GetInt64("id")

	if getCollectionId != nil {
		self.FailJson( getCollectionId )
		return
	}

	var collection = &bean.AudioUserCollection{}
	collection.Id = collectionId

	var collectionResult *bean.AudioUserCollection= collection

	var collectionService = service.GetCollectonServiceInstance()

	var ormService = db.GetOrmServiceInstance()

	var _ , deleteErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		var readCollectionErr = o.Read(collectionResult)

		if readCollectionErr != nil {
			//输出对应的信息
			return nil, nil
		}

		var err = collectionService.Delete(o , collection)
		if err != nil {
			return nil, nil
		}
		return nil , err
	})

	if deleteErr != nil {
		self.FailJson( deleteErr )
		return
	}

	collectionResult.Parse()
	self.Json( collectionResult )
}


func ( self *CollectionController ) FindByUserAndAudio() {
	//获取对应的属性
	var userId , getUserIdErr = self.GetInt64("userId")
	var audioId , getAudioIdErr = self.GetInt64("audioId")

	if getUserIdErr != nil || getAudioIdErr != nil {
		self.FailJson( errors.New("xxx") )
		return
	}

	var collectionService = service.GetCollectonServiceInstance()

	var ormService = db.GetOrmServiceInstance()

	var collection *bean.AudioUserCollection = nil

	ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		collection = collectionService.FindByUserAndAudio( o , userId , audioId )
		return nil, nil
	})

	//直接在这里输出对应的结果
	self.Json( collection )
}

/**
	根据对应的 用户信息 ， 来获取 对应的 对应的 信息 ， 并且附带着对应的信息
 */
 func ( self *CollectionController ) SearchBindCollection() {
 	//先获取对应的信息
 	var userId , getUserIdErr = self.GetInt64("userId")

 	if getUserIdErr != nil {
		self.FailJson(getUserIdErr)
		return
	}

	var ormService = db.GetOrmServiceInstance()
	var collectionService = service.GetCollectonServiceInstance()
	var audioService = service.GetAudioServiceInstance()
	//输出对应的信息
	ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var collections = collectionService.FindByUser( o  , userId )
		service.ForCollection( collections , func(collection *bean.AudioUserCollection, index int) {
			//先转化对应的信息
			collection.Parse()

			//之后根据对应的 id  来进行获取信息
			var audio , readErr = audioService.FindById( o , collection.AudioId )

			if readErr != nil {
				return
			}

		})
		return nil,nil
	})



	fmt.Println(userId)

 }
