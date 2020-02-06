package controllers

import "fmt"

type UserDateDashboardController struct {
	BeegoController
}


/**
	根据一段特定的时间来进行对应的规划
 */
func ( self *UserDateDashboardController ) SearchByAudioId(){
	var audio , getAudioErr = self.GetInt64("audioId")

	if getAudioErr != nil{
		self.Json( getAudioErr )
		return
	}

	fmt.Println( audio )
}