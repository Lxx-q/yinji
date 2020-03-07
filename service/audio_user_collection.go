package service

import (
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
	"yinji/config"
)

type CollectionService struct {

}

func ( self *CollectionService ) InsertCollection( o orm.Ormer , collection *bean.AudioUserCollection ) ( *bean.AudioUserCollection , *bean.AudioDashboard, error){
	//先进行对应的插入数据
	var _ , insertErr = self.Insert( o , collection )

	if insertErr != nil {
		return  nil , nil , insertErr
	}

	var audioDashboardService = GetAudioDashboardServiceInstance()
	//只加载一个的关系
	audioDashboardService.AddCollectionCount( o , collection.AudioId , 1 )

	return collection , nil , nil

}

/**
	下面是基本的原子操作方法
 */

/**
	根据对应的二者的关系 ， 进行插入数据
 */
func ( self *CollectionService ) Insert( o orm.Ormer ,  collection *bean.AudioUserCollection ) (int64 , error)  {
	return o.Insert(collection)
}

/**
	删除对应的收藏关系关系
	1.判断是否拥有这样的关系存在
	2.然后再进行对应的删除计划
 */
func ( self *CollectionService ) Delete( o orm.Ormer ,  collection *bean.AudioUserCollection ) error {
	var _ , deleteErr = o.Delete(collection)

	return  deleteErr
}

/**
	查看目标用户是否与目标音频存在关系
 */
func ( self *CollectionService ) FindByUserAndAudio( o orm.Ormer , userId int64 , auidoId int64) (*bean.AudioUserCollection ,error) {

	var collection = bean.AudioUserCollection{}
	collection.UserId = userId
	collection.AudioId = auidoId

	var err = o.Read(&collection , "user_id" , "audio_id")

	return &collection , err

}

/**
	获取对应的目标用户的收藏信息
	PS. 但是我们不知道是一次性输出所有的关系 ， 还是说根据（folder_id) 关系来进行输出
	还是二者兼得 ，
	并且是否还需要进行对应的分页操作。

	PPS.目前的程序逻辑 ： 输出目标用户的说有关注信息。
 */
func ( self *CollectionService ) FindByUser( o orm.Ormer , userId int64 ) []*bean.AudioUserCollection {
	var collections []*bean.AudioUserCollection

	var qt = o.QueryTable(bean.GetAudioUserCollectionTableName())
	qt = qt.Filter("userId").OrderBy("-create_time").Limit(config.LIMIT_COUNT)
	var _ , err = qt.All(&collections)

	if err != nil {
		return nil
	}

	ForCollection( collections , func(collection *bean.AudioUserCollection, index int) {
		collection.Parse()
	} )

	return collections
}

// 根据对应的 collections 的关系 来进行 获取 信息
func ( self *CollectionService ) SearchBindsUser( o orm.Ormer , collections []*bean.AudioUserCollection) {

	ForCollection( collections , func(collection *bean.AudioUserCollection, index int) {

	})


}

//进行对应的循环的操作
func ForCollection( array []*bean.AudioUserCollection ,function func( collection *bean.AudioUserCollection , index int)){
	var _array = array
	var _Len = len( _array )

	for index:= 0 ; index < _Len ; index ++ {
		var item = _array[ index ]
		function( item , index )
	}

}
var COLLECTION_SERVICE_INSTANCE = &CollectionService{}

func GetCollectonServiceInstance() *CollectionService{
	return COLLECTION_SERVICE_INSTANCE
}
