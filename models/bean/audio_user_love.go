package bean

import "yinji/models/base"

//点赞类
type AudioUserLove struct {
	base.IdStruct
	base.CreateTimeStruct
	AudioId int64 `orm:"column(audio_id)" json:"audioId"`
	UserId int64 `orm:"column(user_id)" json:"userId"`
}

func ( self *AudioUserLove) TableName() string {
	return GetAudioUserLoveTableName()
}

func GetAudioUserLoveTableName() string{
	return "audio_user_love"
}