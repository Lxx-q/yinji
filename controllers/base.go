package controllers

import (
	"github.com/astaxie/beego"
	"yinji/models"
)

type BeegoController struct {

	beego.Controller
}

func (controller *BeegoController ) String( content string ){
	controller.Ctx.WriteString( content );
}

func ( controller *BeegoController) Json( target interface{}){
	var data = controller.Data;
	data["json"] = target;
	controller.ServeJSON();
}


func ( controller *BeegoController) Xml( target interface{}){
	var data = controller.Data;
	data["xml"] = target;
	controller.ServeXML();
}

func ( controller *BeegoController) Jsonp (target interface{}){
	var data = controller.Data;
	data["jsonp"] = target;
	controller.ServeJSONP();
}

//返回 相对应的 页面
func (self *BeegoController) Resource( path string ){
	self.TplName = path ;
}

func (self *BeegoController ) Fail( err error ){
	var response = models.FailResponse(err)
	self.Json(response)
}

