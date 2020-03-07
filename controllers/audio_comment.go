package controllers

import (
	"yinji/service"
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
	"yinji/service/db"
	"yinji/models/base"
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

	/*
	//查询对应的page , 与 count

	var page , getPageErr = self.GetInt("page")

	if getPageErr != nil {
		page = 0
	}

	var count , getCountErr = self.GetInt("count")

	if getCountErr != nil {
		count = 10
	}

	var offset = page * count
	*/
	var instance = service.GetAudioCommentServiceInstance()
	var ormService = db.GetOrmServiceInstance()
	var comments interface{}
	ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		comments = instance.FindAudioCommentsAndUser( o , func(o orm.Ormer) orm.QuerySeter {
			var qs = o.QueryTable(bean.GetAudioCommentTableName()).Filter("AudioId",audioId)
			qs = qs.OrderBy("-create_time")
			//之后进行对应的搜索
			return qs
		})

		return nil, nil
	})
	//之后开始进行对应的参数

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

	var dashboardService = service.GetDashboardServiceInstance()
	var dashboardBase = base.NewDashboardBase()
	dashboardBase.CommentCount = 1
	dashboardService.AddCount(audioId,userId,dashboardBase)

	self.Json( audioComment )
}

/**
	根据对应的  audio Comment 的 id 来搜索信息
 */
func ( self *AudioCommentController ) DeleteCommentById(){
	//收集对应的信息
	var id , getIdErr = self.GetInt64("id")

	if getIdErr != nil {
		self.FailJson( getIdErr )
		return
	}

	//之后进行请求
	var ormService = db.GetOrmServiceInstance()

	//填充对应的 audioComment 的 信息
	var audioComment = bean.AudioComment{}
	audioComment.Id = id

	var audioCommentService = service.GetAudioCommentServiceInstance()

	var _ , transacErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		var readErr = o.Read(&audioComment)
		if readErr != nil  {
			return nil , readErr
		}
		var deleteCommentErr = audioCommentService.DeleteComment(o , &audioComment)
		return nil , deleteCommentErr
	})

	if transacErr != nil {
		self.FailJson(transacErr)
		return
	}

	self.Json(audioComment)

}

/**
	返回 hbulderx 下面的数据
	说白了一次性
 */
func ( self *AudioCommentController) ByAudioInHB(){

}

func (self *AudioCommentController ) PageDetails(){
	self.Resource("details/main.html")
}
