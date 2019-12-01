package utils

import "time"

func Now() time.Time{
	var location , _ = time.LoadLocation("Asia/Chongqing")
	time.Local = location

	return time.Now().Local()
}
