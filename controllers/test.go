package controllers

import (
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
	"yinji/service/db"
	"yinji/service"
	"yinji/models"
)

type TestController struct {

	BeegoController

	Login bean.Login

}

func ( controller *TestController) JsonResult(){

	var data = controller.Data;
	//data["json"] = controller.Login;
	var ormService = db.GetOrmServiceInstance();

	var music [] *bean.Audio;
	var result , _ = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var qs = o.QueryTable("audios")
		qs.All(&music);
		return music , nil
	})

	data["json"] = result;
	controller.ServeJSON();

}

func ( controller *TestController) XmlResult(){
	var data = controller.Data;
	data["xml"] = controller.Login;
	controller.ServeXML();
}



func( controller *TestController) FindHtml(){

	var input = controller.Ctx.Input;

	var page = input.Param(":page")

	controller.TplName = page + ".html";

}

func ( self *TestController ) UploadImage(){
	var file , fileHeader , getFileErr = self.GetFile("image")

	if getFileErr !=  nil {
		models.FailResponse( getFileErr )
		return
	}

	var httpFileService = service.GetHttpFileServiceInstance()
	var filePath , uploadImageErr = httpFileService.UploadImage( fileHeader.Filename , file )

	if uploadImageErr != nil {
		self.Fail( uploadImageErr )
		return
	}

	self.Json( filePath )

}

const SESSION_TEST_KEY = "TEST_KEY"

const SESSION_TEST_VALUE = "TEST_SESSION_VALUE"

func (self *TestController ) SetSession(){
	//设置对应的 信息
	self.StartSession().Set( SESSION_TEST_KEY , SESSION_TEST_VALUE )
	self.String("helllo , world")
}

func (self *TestController) GetSession(){
	var value = self.StartSession().Get( SESSION_TEST_KEY)
	self.Json( value )
}
