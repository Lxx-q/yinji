package utils

import (
	"github.com/astaxie/beego/orm"
	"math"
	"math/rand"
	"yinji/config"
	"time"
)

func Limit( seter orm.QuerySeter ) orm.QuerySeter {
	return seter.Limit( config.LIMIT_COUNT )
}

/**
	生成随机数
*/
func Rann( number  int) int64 {
	//先获取最大数
	rd :=rand.New(rand.NewSource(time.Now().UnixNano()))
	var maxCount = math.Pow10( number )
	//输出任意数
	return rd.Int63n( int64( maxCount ))
}

