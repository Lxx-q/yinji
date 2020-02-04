package service


import (
	"yinji/service/db"
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
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

/**
 */
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

var AUDIO_COMMENT_SERVICE_INSTANCE = &AudioCommentService{ db.GetOrmServiceInstance()}

func GetAudioCommentServiceInstance() *AudioCommentService{
	return AUDIO_COMMENT_SERVICE_INSTANCE;
}
