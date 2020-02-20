package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
)

type AudioFastService struct {

}

/**
	根据目标的 audioId 来搜索信息

		audioId int64
	得出结果s

 */
func ( self *AudioFastService ) SearchByAudioId( o orm.Ormer , audioId int64 ){

}

/**
	根据对应的id 来搜索对应的信息
 */
func( self *AudioFastService ) FindById( o orm.Ormer , id int64 ) ( *bean.AudioFast , error){
	var audioFast = &bean.AudioFast{}
	audioFast.Id = id
	var readErr = o.Read( audioFast )

	if readErr ==nil {
		audioFast.Parse()
	}

	return audioFast , readErr
}

var AUDIO_FAST_SERVICE_INSTANCE = &AudioFastService{}

func GetAudioFastServiceInstance() * AudioFastService{
	return AUDIO_FAST_SERVICE_INSTANCE
}
