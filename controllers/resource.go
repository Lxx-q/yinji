package controllers

import "fmt"

type ResourceController struct {
	BeegoController
}

/**
	根据 audio id 可以直接找到对应的 audio 音频资源信息
 */
func ( self *ResourceController ) Audio(){
	var id , getIdErr = self.GetInt64("id")

	if getIdErr != nil {
		self.FailJson( getIdErr )
		return
	}

	fmt.Println( id )
	self.Json( id )
}
