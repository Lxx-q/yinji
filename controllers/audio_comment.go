package controllers

import (
	"yinji/service"
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
)

type AudioCommentController struct {
	BeegoController
}

func (self *AudioCommentController) ByAudioId(){
	//得到对应的 信息
	var audioId , audioIdErr = self.GetInt64("audioId")

	//之后关闭对应的 信息 窗口
	if audioIdErr != nil {
		self.FailJson( audioIdErr )
		return;
	}

	var instance = service.GetAudioCommentServiceInstance()

	//之后开始进行对应的参数
	var comments = instance.FindAudioCommentsAndUser(func(o orm.Ormer) orm.QuerySeter {
		var qs = o.QueryTable(bean.GetAudioCommentTableName()).Filter("audioId",audioId)
		//之后进行对应的搜索
		return qs
	})

	self.Json( comments )
}

func (self *AudioCommentController) InsertComment(){
	var audioId , getAudioIdErr = self.GetInt64("audioId")

	if getAudioIdErr != nil {
		self.FailJson(nil)
		return
	}

	var userId , getUserIdErr = self.GetInt64("userId")

	if getUserIdErr != nil {
		self.FailJson( getUserIdErr )
		return
	}

	var content  = self.GetString("content");
	var instance = service.GetAudioCommentServiceInstance()

	var audioComment = bean.AudioComment{}

	//设置对应的信息
	audioComment.New()
	audioComment.AudioId = audioId
	audioComment.UserId = userId
	audioComment.Content = content

	var insertCommentErr = instance.InsertComment( &audioComment )

	if insertCommentErr != nil {
		self.FailJson( insertCommentErr )
		return
	}

	self.Json( audioComment )
}

func (self *AudioCommentController ) PageDetails(){
	self.Resource("details/main.html")
}
