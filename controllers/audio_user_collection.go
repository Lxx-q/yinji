package controllers

import (
	"errors"
	"yinji/service"
	"yinji/models/bean"
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
	"yinji/models/bind"
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

	//folderId ， 为非必要属性
	var folderId , getFolderIdErr = self.GetInt64("folderId")

	//将收集而来的信息全部进行收入
	var collection = &bean.AudioUserCollection{}
	//生成对应的信息
	collection.New()

	collection.UserId = userId
	collection.AudioId = audioId

	if getFolderIdErr == nil {
		collection.FolderId = &folderId
	}


	var collectionService = service.GetCollectonServiceInstance()

	var ormService = db.GetOrmServiceInstance()

	var _ , insertErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		//先根据对应的 audio 以及 folderId 进行查询 ， 查看是否能查询知道数据 ， 倘若能 ， 则退出查询
		var readError = o.Read( collection ,"userId","audioId","folderId")

		if readError == nil {
			return nil , errors.New("该实例已经存在")
		}
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
			self.FailJson( readCollectionErr )
			//输出对应的信息
			return nil, nil
		}

		var err = collectionService.Delete(o , collection)
		if err != nil {
			self.FailJson( err )
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
	获取某个文件夹下面的收藏信息
 */
 func ( self *CollectionController ) SearchCollectionAndAudio() {

 	/*
 	//先获取对应的信息
 	var userId , getUserIdErr = self.GetInt64("userId")

 	if getUserIdErr != nil {
		self.FailJson(getUserIdErr)
		return
	}

	*/
	var folderId , getFolderIdErr = self.GetInt64("folderId")

	if getFolderIdErr != nil {
		self.FailJson( getFolderIdErr )
		return
	}

	var ormService = db.GetOrmServiceInstance()
	var audioService = service.GetAudioServiceInstance()

	var shuchulist = make([]bind.CollectionAndAudio , 0)
	//输出对应的信息
	ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {

		var collections []*bean.AudioUserCollection

		//计算并且得出最后的时间
		var qt = o.QueryTable( bean.GetAudioUserCollectionTableName())

		qt.Filter("folderId" , folderId ).OrderBy("-create_time").All(&collections)

		service.ForCollection( collections , func(collection *bean.AudioUserCollection, index int) {
			//先转化对应的信息
			collection.Parse()

			//之后根据对应的 id  来进行获取信息
			var audio , readErr = audioService.FindById( o , collection.AudioId )

			if readErr != nil {
				return
			}

			//利用新型的输出结构体
			var collectionAndAudio = bind.CollectionAndAudio{}

			collectionAndAudio.AudioUserCollection = collection
			collectionAndAudio.Audio = audio

			shuchulist = append(shuchulist, collectionAndAudio)
		})

		return shuchulist , nil
	})

	self.Json(shuchulist)

 }

/**
	修改对应的 collection 信息
*/
func ( self *CollectionController ) UpdateCollection(){

	var id , getIdErr = self.GetInt64("id")

	if getIdErr != nil {
		self.FailJson( getIdErr )
		return
	}

	var collectionFolderId , getCollectionFolderIdErr = self.GetInt64("folderId")

	if getCollectionFolderIdErr != nil {
		self.FailJson( getCollectionFolderIdErr )
		return
	}

	var ormService = db.GetOrmServiceInstance()

	var collection = bean.AudioUserCollection{}

	var _ , transacErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		collection.Id = id
		var readErr = o.Read(&collection)
		if readErr != nil {
			return nil , readErr
		}
		collection.FolderId = &collectionFolderId
		return o.Update( &collection )
	})

	if transacErr != nil {
		self.FailJson( transacErr )
		return
	}

	self.Json( collection )
}

