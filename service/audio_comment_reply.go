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

/**
 */
func (self *AudioCommentReplyService) FindAudoCommentReplys(function func(qs orm.QuerySeter ) orm.QuerySeter) []*bean.AudioCommentReply{
	var service = db.GetOrmServiceInstance()
	var replies []*bean.AudioCommentReply
	service.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var qs = o.QueryTable(bean.GetAudioCommentReplyTableName())
		qs = function( qs ).OrderBy()
		qs.All(&replies)
		return nil,nil
	})
	return replies;
}

func (self *AudioCommentReplyService ) FindAudoCommentReplyAndUser( o orm.Ormer , function func(qs orm.QuerySeter) orm.QuerySeter )[]*bean.AudioCommentReplyAndUser{
	var userService = GetUserServiceInstance()

	var replies []*bean.AudioCommentReply = self.FindAudoCommentReplys( function )

	var replyAndUser_list = make([] *bean.AudioCommentReplyAndUser , 0 , 0 )

	//之后输出对应的信息
	for index:= 0 ; index < len(replies) ; index ++ {
		var replyAndUser = &bean.AudioCommentReplyAndUser{}
		var reply = replies[index]

		var user ,_ = userService.FindUserById( o , reply.UserId)
		var target ,_= userService.FindUserById(o , reply.TargetId)

		reply.Parse()

		//绑定数据
		replyAndUser.AudioCommentReply = reply
		replyAndUser.BindTarget( target )
		replyAndUser.BindUser( user )


		//最后输出
		replyAndUser_list = append(replyAndUser_list , replyAndUser)
	}
	
	return replyAndUser_list
}

//获取对应的信息
var AUDIO_COMMENT_REPLY_SERVICE = &AudioCommentReplyService{}

func GetAudioCommentReplyServiceInstance() *AudioCommentReplyService {
	return AUDIO_COMMENT_REPLY_SERVICE
}