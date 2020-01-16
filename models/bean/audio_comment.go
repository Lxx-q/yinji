package bean

import (
	"yinji/models/base"
	"time"
)

type AudioComment struct {

	UserId int64 `orm:"column(user_id)" json:"userId"`
	AudioId int64 `orm:"column(audio_id)" json:"audioId"`
	Content string `orm:"column(content)" json:"content"`

	base.IdAndCodeStruct
	//设置对应的 创造时间
	base.CreateTimeStruct

}

func ( self *AudioComment) TableName() string {
	return GetAudioCommentTableName()
}

func GetAudioCommentTableName() string{
	return "audio_comment"
}

func (self *AudioComment) New(){
	var t = time.Now()
	self.CreateTimeStruct.NewEntity(t)
	self.IdAndCodeStruct.NewEntity(t)
}


/**
	AudioCommentAndUser ， 对应的 附带上 对应的 user 的 信息
 */
type AudioCommentAndUser struct {
	*AudioComment
	UserName string `json:"userName"`
	UserImage string `json:"userImage"`
}

func ( self *AudioCommentAndUser ) Bind( user *User ){
	self.UserName = user.Name
	self.UserImage = user.Image
}