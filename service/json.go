package service

import "github.com/astaxie/beego"

type ReturnResultService struct {
}

func ( *ReturnResultService) Json(controller *beego.Controller , target interface{}){
	var data = controller.Data;
	data["json"] = target;
	controller.ServeJSON();
}


func ( *ReturnResultService) Xml(controller *beego.Controller , target interface{}){
	var data = controller.Data;
	data["xml"] = target;
	controller.ServeXML();
}

func ( *ReturnResultService) Jsonp (controller *beego.Controller , target interface{}){
	var data = controller.Data;
	data["jsonp"] = target;
	controller.ServeJSONP();
}


var RETURN_RESULT_SERVICE_INSTANCE = new(ReturnResultService);

func GetReturnResultServiceInstance() *ReturnResultService{
	return RETURN_RESULT_SERVICE_INSTANCE;
}