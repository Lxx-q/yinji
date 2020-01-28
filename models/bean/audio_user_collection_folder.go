package bean

import "yinji/models/base"

type AudioCollectionFolder struct {
	base.IdStruct
	Name string `orm:"column(name)" json:"name" `
	Introduction string `orm:"column(introduction)" json:"introduction"`
	base.CreateTimeAndModifyTimeStruct
	UserId int64 `orm:"column(user_id)" json:"userId"`
}

func ( self *AudioCollectionFolder) TableName() string {
	return GetAudioCollectionFolderTableName()
}

func GetAudioCollectionFolderTableName() string{
	return "audio_user_collection_folder"
}


