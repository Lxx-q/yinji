package service


import (
	"yinji/service/db"
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
	"yinji/models/base"
)

type AudioCommentService struct {
	ormServiceInstance  *db.OrmService
}

func (self *AudioCommentService) Parse( comments []*bean.AudioComment){
	for index:= 0 ; index < len(comments) ; index ++ {
		var comment = comments[index]
		comment.Parse()
	}
}

func (self *AudioCommentService) FindAudoComments( function func(o orm.Ormer) orm.QuerySeter ) []*bean.AudioComment{
	var service = db.GetOrmServiceInstance()
	var comments []*bean.AudioComment
	service.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var qs = function( o ).OrderBy()
		qs.All(&comments)
		return nil,nil
	})
	return comments;
}

func (self *AudioCommentService ) FindAudioCommentsAndUser( o orm.Ormer ,  function func(o orm.Ormer) orm.QuerySeter ) [] *bean.AudioCommentAndUser{

	var userService = GetUserServiceInstance()

	var comments []*bean.AudioComment = self.FindAudoComments( function )

	var bindUser = make([] *bean.AudioCommentAndUser , 0 , 0 )

	//之后输出对应的信息
	for index:= 0 ; index < len(comments) ; index ++ {

		var comment = comments[ index ];
		var user , _  = userService.FindUserById( o , comment.UserId)
		var commentAndUser = &bean.AudioCommentAndUser{}

		commentAndUser.AudioComment = comment
		commentAndUser.Bind( user )
		comment.Parse()

		bindUser = append( bindUser, commentAndUser  )

	}

	return bindUser;
}

func (self *AudioCommentService) InsertComment( audioComment *bean.AudioComment ) error{

	var _ , insertErr = self.ormServiceInstance.Transaction(func(o orm.Ormer) (interface{}, error) {
		return o.Insert(audioComment)
	})

	return insertErr
}

/**
	删除对应的 comment 信息
*/
func ( self *AudioCommentService) DeleteComment( o orm.Ormer ,  audioComment *bean.AudioComment ) error {
	var _  , deleteErr = o.Delete(audioComment)
	// 之后我们将结果输出
	if deleteErr != nil {
		return deleteErr
	}

	var audioId = audioComment.AudioId
	var userId = audioComment.UserId

	//添加记录
	var dashboardService = GetDashboardServiceInstance()
	var dashboardBase = base.NewDashboardBase()
	dashboardBase.CollectionCount = -1
	dashboardService.AddCount(audioId,userId,dashboardBase)

	return nil
}

var AUDIO_COMMENT_SERVICE_INSTANCE = &AudioCommentService{ db.GetOrmServiceInstance()}

func GetAudioCommentServiceInstance() *AudioCommentService{
	return AUDIO_COMMENT_SERVICE_INSTANCE;
}
