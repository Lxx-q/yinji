package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
)

type AudioBrowseHistroyService struct {

}

//对应的 音频 次数信息
func ( self *AudioBrowseHistroyService ) FindById( o orm.Ormer , id int64 ) ( *bean.AudioBrowseHistroy , error ){
	var history = &bean.AudioBrowseHistroy{}
	history.Id = id
	var readErr = o.Read( history )
	return history , readErr
}

func ( self *AudioBrowseHistroyService ) New( o orm.Ormer , browseHistory *bean.AudioBrowseHistroy ) error {
	var  _ , insertError = o.Insert( browseHistory )
	return insertError
}

/**
	根据对应的Audio id 信息来生成字符信息
 */
func ( self *AudioBrowseHistroyService ) NewByAudioId( o orm.Ormer , audioId int64 )  ( *bean.AudioBrowseHistroy  , error ){
	var browseHistory = &bean.AudioBrowseHistroy{}
	browseHistory.Id = audioId
	browseHistory.BrowseAllCount = 0
	browseHistory.New()
	var newErr = self.New( o , browseHistory )
	return browseHistory , newErr
}

var AUDIO_BROWSE_HISTORY_SERVICE_INSTANCE = &AudioBrowseHistroyService{}

func GetAUdioBrowseHistoryServiceInstance() *AudioBrowseHistroyService{
	return AUDIO_BROWSE_HISTORY_SERVICE_INSTANCE
}