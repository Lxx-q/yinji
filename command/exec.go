package command

import "os/exec"

/**
	头部的信息 ， 便是信息头 比如 说 dir  ， ffmpeg
	后面便是对应的 下面的信息 ，
	同步使用
 */
func ExecRun( cmd *exec.Cmd ) ( string , error ){
	var runErr = cmd.Run()
	if runErr != nil {
		return "" , runErr
	}

	return "" , nil

}

/**
	异步使用
 */
func ExecStart( cmd *exec.Cmd)  ( string , error ){
	var runErr = cmd.Start()
	if runErr != nil {
		return "" , runErr
	}

	return "" , nil
}
