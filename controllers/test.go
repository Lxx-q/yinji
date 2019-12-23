package controllers

import (
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
	"yinji/service/db"
	"yinji/service"
	"github.com/astaxie/beego/httplib"
	"yinji/service/url"
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
		self.FailJson(getFileErr)
		return
	}

	var httpFileService = service.GetHttpFileServiceInstance()
	var filePath , uploadImageErr = httpFileService.UploadImage( fileHeader.Filename , file )

	if uploadImageErr != nil {
		self.FailJson( uploadImageErr )
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

func (self *TestController) WebTest(){
	var urlString = url.BuildApiUrl("api/test")
	var req = httplib.Post( urlString )
	var string , err = req.String()

	if err != nil {
		self.FailJson(err)
	}

	self.String( string )
}

func (self *TestController) ApiTest(){
	self.String("hello , world")
}

func (self *TestController) ApiAudioComment(){

	var service = db.GetOrmServiceInstance()
	var comments []*bean.AudioComment
	service.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var qs = o.QueryTable(bean.GetAudioCommentTableName())
		qs.All(&comments)
		return nil,nil
	})
	for index:= 0 ; index < len(comments) ; index++ {
		var comment = comments[index]
		comment.Parse()
	}

	self.Json( comments )
}
