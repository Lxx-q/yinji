package ffmpeg

import (
	"os/exec"
	"yinji/command"
	"fmt"
)

const CMD_HEADER_FFMPEG = "ffmpeg"

/**
	将对应的 amr 文件 转化成对应的 mp3 的文件
 */
func AmrToMp3( originFileName string , encodeFileName string ) error {

	//获取对应的 cmd
	var cmd = exec.Command( CMD_HEADER_FFMPEG , "-i" , originFileName , encodeFileName )
	//之后进行运行
	var result , runErr = command.ExecRun( cmd )
	fmt.Println( result )
	return runErr
}
