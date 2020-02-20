package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
)

type AudioUserHistoryService struct {

}

/**
	查询或者新建
*/
func ( self *AudioUserHistoryService ) FindOrNew( o orm.Ormer , userId int64 , audioId int64 ) ( *bean.AudioUserHistroy , error ){
	var history , findErr = self.FindById( o , userId , audioId )
	if findErr != nil {
		var newErr error
		history , newErr = self.New( o , history)
		return history , newErr
	}

	return history , nil

}

func ( self *AudioUserHistoryService ) New( ormer orm.Ormer , history *bean.AudioUserHistroy ) ( *bean.AudioUserHistroy , error ){
	history.New()
	var _ , errors = ormer.Insert( history )
	return history , errors
}

//根据对应的id 来进行搜索
func ( self *AudioUserHistoryService ) FindById( o orm.Ormer , userId int64 , audioId int64 )(*bean.AudioUserHistroy , error ){
	var history = &bean.AudioUserHistroy{}
	history.UserId = userId
	history.AudioId = audioId
	var readErr = o.Read( history ,"userId","audioId")
	if readErr == nil{
		history.Parse()
	}
	return history , readErr
}


var AUDIO_USER_HISTORY_SERVICE_INSTANCE = &AudioUserHistoryService{}

func GetAudioUserHistoryServiceInstance() *AudioUserHistoryService{
	return AUDIO_USER_HISTORY_SERVICE_INSTANCE
}