package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
)

type AudioUserLoveService struct {

}

func ( self *AudioUserLoveService ) FindLove( o orm.Ormer , userId int64 , audioId int64 ) ( *bean.AudioUserLove , error ){

	var love = &bean.AudioUserLove{}

	//插入对应的信息
	love.UserId = userId
	love.AudioId = audioId

	//将结果输入
	var  readErr = o.Read(love , "userId" , "audioId")

	return love , readErr
}

/**
	进行点赞的操作
 */

func ( self *AudioUserLoveService ) InsertLove( o orm.Ormer ,  userId int64 , audioId int64 ) ( *bean.AudioUserLove , error ) {
	var love , readErr = self.FindLove( o , userId , audioId )

	//只有当 love  无数据的时候 ， 我们才能插入数据
	if readErr == nil {
		return nil , readErr
	}

	love.New()

	var _ , insertErr = o.Insert( love )

	var audioDashboardService = GetAudioDashboardServiceInstance()

	audioDashboardService.AddLoveCount( o , love.AudioId , 1 )

	return love , insertErr
}

func ( self *AudioUserLoveService ) DeleteLove( o orm.Ormer , userId int64 , audioId int64 ) ( *bean.AudioUserLove , error) {
	var love ,  findErr = self.FindLove( o , userId , audioId )

	if findErr != nil {
		return nil , findErr
	}

	var _ , deleteErr = o.Delete( love )

	return love , deleteErr
}

var AUDIO_USER_LOVE_SERVICE_INSTANCE = &AudioUserLoveService{}

func GetAudioUserLoveServiceInstance() *AudioUserLoveService{
	return AUDIO_USER_LOVE_SERVICE_INSTANCE
}