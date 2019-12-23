package bean

import (
	"yinji/models/base"
	"time"
)

type AudioCommentReply struct{

	UserId int64 `orm:"column(user_id)" json:"userId"`
	AudioCommentId int64 `orm:"column(audio_comment_id)" json:"audioCommentId"`
	Content string `orm:"column(content)" json:"content"`
	TargetId int64 `orm:"column(target_id)" json:"targetId"`

	base.IdStruct
	//设置对应的 创造时间
	base.CreateTimeStruct
}

//设置连接苗斌
func ( self *AudioCommentReply ) TableName() string {
	return GetAudioCommentReplyTableName()
}

func GetAudioCommentReplyTableName() string {
	return "audio_comment_reply"
}

func ( self *AudioCommentReply ) New(){
	var t = time.Now()
	self.CreateTimeStruct.NewEntity(t)
	self.IdStruct.NewEntity(t)
}