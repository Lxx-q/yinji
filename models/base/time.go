package base

import (
	"time"
)


type TimeStruct struct {
	//毫秒
	MilliTime int64 `json:"millTime"`
	//下面显示基本的信息
	Year int `json:"year"`
	Month int `json:"month"`
	Day int `json:"day"`

	Hour int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
	WeekDay int `json:"weekDay"`
}

// 根据对应的时间 来得到 相对应的时间
func ( self *TimeStruct) ParseTime( t time.Time ){
	//利用对应的方法进行输出
	self.MilliTime = t.Unix()
	self.Year , _ , self.Day = t.Date()
	self.Month = int(t.Month())
	self.Hour = t.Hour()
	self.Minute = t.Minute()
	self.Second = t.Second()
	self.WeekDay = int( t.Weekday() )
}

