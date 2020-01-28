package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
)

type AudioCollectionFolderService struct {

}

// 根据对应的 用户 id 来进行获取所有的 folder 信息
func ( self *AudioCollectionFolderService ) AllByUserId( o orm.Ormer , userId int64 ) ( []*bean.AudioCollectionFolder , error ) {
	//根据对应的信息 来进行输出数据
	var folderList []*bean.AudioCollectionFolder

	var qt = o.QueryTable(bean.GetAudioCollectionFolderTableName())

	var _ , allErr = qt.Filter("userId",userId).All(&folderList)

	ForCollectionFolder( folderList , func(folder *bean.AudioCollectionFolder, index int) {
		folder.Parse()
	})

	return folderList , allErr

}

func ForCollectionFolder( array []*bean.AudioCollectionFolder ,function func( folder *bean.AudioCollectionFolder , index int)){
	var _Len = len( array )

	for index:= 0 ; index < _Len ; index ++ {
		var item = array[ index ]
		function( item , index )
	}
}

var AUDIO_COLLECTION_FOLDER_SERVICE = &AudioCollectionFolderService{}

//输出对应的 文件夹 的 service
func GetCollectionFolderServiceInstance() *AudioCollectionFolderService {
	return AUDIO_COLLECTION_FOLDER_SERVICE
}