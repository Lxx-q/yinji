package bean

import (
	"yinji/models/base"
	"time"
)

type UserFollowInterest struct {
	base.IdStruct
	UserId int64
	TargetUserId int64
	base.CreateTimeStruct
}

/**
	初始化函数
*/

func (self *UserFollowInterest) NewEntity( t time.Time ){
	self.IdStruct.NewEntity( t )
	self.CreateTimeStruct.NewEntity( t )
}

func (self *UserFollowInterest) New(){
	var t = time.Now()
	self.NewEntity( t )
}

func (self *UserFollowInterest) TableName() string{
	return GetUserFollowInterestTableName()
}

func GetUserFollowInterestTableName() string{
	return "user_follow_interest"
}
