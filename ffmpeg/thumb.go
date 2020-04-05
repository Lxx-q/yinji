package ffmpeg

import (
	"os/exec"
	"yinji/command"
	"fmt"
)

/**
	压缩图片资源操作
 */
func Thumb( originFileName string , thumbFileName string ) error{
	var cmd = exec.Command( CMD_HEADER_FFMPEG , "-hide_banner","-i" , originFileName,"-pix_fmt","pal8" , thumbFileName )
	var result , runErr = command.ExecRun( cmd )
	fmt.Println( result )
	return runErr
}
