package main

import (
	"yinji/ffmpeg"
	"fmt"
)

func main(){
	var amr_f = "E:/Go/code/src/yinji/static/music/mp3/2.amr"

	var mp3_f = "E:/Go/code/src/yinji/static/music/mp3/2.mp3"

	var err = ffmpeg.AmrToMp3( amr_f , mp3_f )

	if err != nil {
		fmt.Println(err.Error())
	}
}
