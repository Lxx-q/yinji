package controllers

import (
	"yinji/models/bean"
	"yinji/service"
	"github.com/astaxie/beego/orm"
	"yinji/service/db"
)

type AudioCommentReplyController struct {
	BeegoController
}

func ( self *AudioCommentReplyController ) InsertCommentReply(){
	//获取对应的信息
	var commentId , getAudioIdErr = self.GetInt64("commentId")

	if getAudioIdErr != nil {
		self.FailJson(nil)
		return
	}

	var userId , getUserIdErr = self.GetInt64("userId")

	if getUserIdErr != nil {
		self.FailJson( getUserIdErr )
		return
	}

	var targetId , getTargetIdErr = self.GetInt64("targetId")

	if getTargetIdErr != nil {
		self.FailJson(getTargetIdErr)
		return
	}

	var content  = self.GetString("content");

	var commentReply = bean.AudioCommentReply{}

	//输入对应的信息
	commentReply.AudioCommentId = commentId
	commentReply.UserId = userId
	commentReply.Content = content
	commentReply.TargetId = targetId

	commentReply.New()

	var instance = service.GetAudioCommentReplyServiceInstance()

	var insertCommentReplyErr = instance.InsertCommentReply( &commentReply )

	if insertCommentReplyErr != nil {
		self.FailJson(insertCommentReplyErr)
		return
	}

	self.Json( commentReply )
}

//根据对应的 commentId 来加载对应的信息
func ( self *AudioCommentReplyController ) ByCommentId(){
	//得到对应的 信息
	var commentId , commentIdErr = self.GetInt64("commentId")

	if commentIdErr != nil {
		self.FailJson( commentIdErr )
		return
	}

	var instance = service.GetAudioCommentReplyServiceInstance()
	var ormService = db.GetOrmServiceInstance()

	var result interface{}

	ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {

		result = instance.FindAudoCommentReplyAndUser( o , func( qs  orm.QuerySeter ) orm.QuerySeter {
			return qs.Filter("audioCommentId", commentId)
		})
		return nil, nil

	})

	self.Json( result )
}