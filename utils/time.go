package utils

import "time"

func Now() time.Time{
	var location , _ = time.LoadLocation("Asia/Chongqing")
	time.Local = location

	return time.Now().Local()
}

func Date( year int , month int , day int ) *time.Time {
	var date = time.Date( year , time.Month(month) , day , 0 , 0 , 0 , 0 , time.Local )
	return &date
}

func Unix( times int64 ) *time.Time {
	var thousand int64 = 1000
	var t = time.Unix( times / thousand , times % thousand)
	return &t
}

func Today() *time.Time{
	var today = time.Now()
	var year = today.Year()
	var month = int(today.Month())
	var day = today.Day()
	return Date( year , month , day )
}
