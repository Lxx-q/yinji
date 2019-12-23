package service

import (
	"yinji/service/db"
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
)

//进行输入对应的信息
type AudioCommentReplyService struct{
	ormService *db.OrmService
}

/**
	插入对应的信息
 */
func (self *AudioCommentReplyService) InsertCommentReply( audioCommentReply *bean.AudioCommentReply ) error{

	var _ , insertErr = self.ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		return o.Insert(audioCommentReply)
	})
	return insertErr
}

//获取对应的信息
var AUDIO_COMMENT_REPLY_SERVICE = &AudioCommentReplyService{}

func GetAudioCommentReplyServiceInstance() *AudioCommentReplyService {
	return AUDIO_COMMENT_REPLY_SERVICE
}