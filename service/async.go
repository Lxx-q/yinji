package service

/**
	异步操作结构化参数
 */
func Async( function func()  ){
	go func() {
		function()
	}()
}
