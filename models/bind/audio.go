package bind

import "yinji/models/bean"

/*
	主要便是相对应的 Audio 附带上 相对应的 User 的 部分信息
*/
type AudioAndUser struct {

	//根据对应的 bean.Audio 来作为主题
	bean.Audio

	//然后 我们 便可以附带上对应的 User 的 一些 我们 所需要的 信息

	UserName string
	//对应的 user的 图片的 地址
	UserImage string
}

//利用对应的 User 来填充对应的 所需要的 信息
func (self *AudioAndUser ) BindUser( user bean.User){
	//将相对应的 信息 进行 设置 进去
	self.UserName =  user.Name;
	self.UserImage = user.Image;
}

type CollectionAndAudio struct {
	*bean.AudioUserCollection
	Audio *bean.Audio `json:"audio"`
}