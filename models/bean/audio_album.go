package bean

import "yinji/models/base"

type AudioAlbum struct {
	Name string
	Introduction string
	UserId int64
	base.IdStruct
	base.CreateTimeAndModifyTimeStruct
}