package controllers

import (
	"github.com/astaxie/beego"
	"yinji/protocol"
)

type BeegoController struct {

	beego.Controller
}

func (controller *BeegoController ) String( content string ){
	controller.Ctx.WriteString( content );
}

func ( controller *BeegoController) Json( target interface{}){
	var commStruct = controller.Success( target )
	controller.json( commStruct )
}

func ( controller *BeegoController) Xml( target interface{}){
	var commStruct = controller.Success( target )
	controller.xml( commStruct )
}

func ( controller *BeegoController) Jsonp (target interface{}){
	var commStruct = controller.Success( target )
	controller.jsonp( commStruct )
}

//返回 相对应的 页面
func (self *BeegoController) Resource( path string ){
	self.TplName = path ;
}

//返回失败的结果，json

func (self *BeegoController ) FailJson( err error ){
	//var response = models.FailResponse(err)
	var commStruct = protocol.CommStruct{}
	commStruct.Type = 2;
	commStruct.Content = err
	self.json(commStruct)
}

func ( self *BeegoController ) Success( content interface{}) *protocol.CommStruct{
	var comm = protocol.CommStruct{}
	comm.Type = protocol.COMM_TYPE_SUCCESS
	comm.Content = content
	return &comm
}

//下面作为返回对应结果的基石

/**
	下面的函数暂时不开放 ， 主要的原因 主要便是由于 ，
	在当前业务中 ， 我们 主要只是用 json 一种格式 ， 其他两种格式尚未解除到 ， 一次性解锁过多 ， 反而增加了耦合
	因此 ， 在这里 ， 我只是做了一下
*/
func ( controller *BeegoController ) json( target interface{} ){
	var data = controller.Data;
	data["json"] = target
	controller.ServeJSON();
}

func ( controller *BeegoController) xml( target interface{}){
	var data = controller.Data;
	data["xml"] = target;
	controller.ServeXML();
}

func ( controller *BeegoController) jsonp (target interface{}){
	var data = controller.Data;
	data["jsonp"] = target;
	controller.ServeJSONP();
}
