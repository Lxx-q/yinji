package controllers

import (
	"yinji/models/bean"
	"yinji/service"
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