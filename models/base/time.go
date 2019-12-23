package base

import (
	"time"
)


type TimeStruct struct {
	//毫秒
	MilliTimes int64
	//下面显示基本的信息
	Year int
	Month int
	Day int

	Hours int
	Minutes int
	Second int
	WeekDay int
}

// 根据对应的时间 来得到 相对应的时间
func ( self *TimeStruct) ParseTime( t time.Time ){
	//利用对应的方法进行输出
	self.MilliTimes = t.Unix()
	self.Year , _ , self.Day = t.Date()
	self.Month = int(t.Month())
	self.Hours = t.Hour()
	self.Minutes = t.Minute()
	self.Second = t.Second()
	self.WeekDay = int( t.Weekday() )
}

