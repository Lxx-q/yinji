package bean

import "yinji/models/base"

//转发
type AudioUserForward struct {
	base.IdStruct
	base.CreateTimeStruct
	AudioId int64 `orm:"column(audio_id)" json:"audioId"`
	UserId int64 `orm:"column(user_id)" json:"userId"`

	//可能将来还会出现添加类型这一个说法，但是，我们也不需要太着急添加便是
}


func ( self *AudioUserForward) TableName() string {
	return GetAudioUserForwardTableName()
}

func GetAudioUserForwardTableName() string{
	return "audio_user_forward"
}

