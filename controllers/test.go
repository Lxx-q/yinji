package controllers

import (
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
	"yinji/service/db"
	"yinji/service"
	"github.com/astaxie/beego/httplib"
	"yinji/service/url"
	"yinji/models/base"
	"strconv"
	"os"
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

	var hostResourceService = service.GetResourceImageServiceInstance()
	var o = orm.NewOrm()
	var filename = fileHeader.Filename
	var resourceImage, uploadErr = hostResourceService.UploadImage( o ,filename  ,file)

	if uploadErr != nil {
		self.String( uploadErr.Error() )
	}

	self.String(  strconv.FormatInt(resourceImage.Id , 10) )

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

func (self *TestController) TestResourceRedirect(){
	self.Redirect("/yinji/resources/image/none.jpg",302)
}

func ( self *TestController ) TestUploadPage(){
	self.Resource("uploadFile.html")
}

func ( self *TestController ) TestDashboard(){
	audioId , getAudioIdErr := self.GetInt64("audioId")

	if getAudioIdErr != nil {
		audioId = 2
	}

	userId , getUserIdErr := self.GetInt64("userId")

	if getUserIdErr != nil {
		userId = 2
	}

	var dashboardService = service.GetDashboardServiceInstance()

	dashboardBase := &base.DashboardBase{}
	dashboardBase.BrowseCount = 1
	dashboardBase.CommentCount = 1
	dashboardBase.CommentCount = 1
	dashboardBase.LoveCount = 1
	dashboardService.AddCount(audioId , userId , dashboardBase)


	self.Json( dashboardBase )
}

/**
	将对应的 audio 的格式进行转化
*/

func ( self *TestController ) AudioFormatChange(){
	var ormService = db.GetOrmServiceInstance()
	var audioList []*bean.Audio
	var resourceImageService = service.GetResourceImageServiceInstance()
	var resourceAudioService = service.GetResourceAudioServiceInstance()
	var httpFileService = service.GetHttpFileServiceInstance()
	ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		var qs = o.QueryTable("audio")
		var _ , allErr = qs.All(&audioList)

		if allErr != nil {
			return nil , allErr
		}
		var audioLen = len(audioList)
		for index:=0; index < audioLen ; index ++   {

			var audio = audioList[index]
			if audio.Image != "" {
				var absImageFile = httpFileService.BuildServerPath(audio.Image)
				var imageFile, imageErr = os.Open(absImageFile)
				if imageErr != nil {

				}
				var resourceImage, _= resourceImageService.UploadImage(o, audio.Image, imageFile)
				audio.ResourceImageId = resourceImage.Id
			}

			var absAudioFile = httpFileService.BuildServerPath(audio.Url)
			var audioFile , audioErr = os.Open(absAudioFile)
			if audioErr != nil {
				continue
			}
			var resourceAudio , _ = resourceAudioService.UploadAudio( o , audio.Url , audioFile )
			audio.ResourceAudioId = resourceAudio.Id
			o.Update(audio)
		}

		return nil , allErr
	})

	self.Json(audioList)

}

/**

*/

func ( self *TestController) ChangeResourceAudioFormatChange(){

	var ormService = db.GetOrmServiceInstance()

	var resource_image_list []*bean.ResourceImage
	ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		var qt = o.QueryTable("resource_image")
		qt.All(&resource_image_list)
		var audioLen = len(resource_image_list)
		for index:=0; index < audioLen ; index++ {
			var resourceImage= resource_image_list[index]
			if resourceImage.ThumbResourceId == 0 {
				resourceImage.ThumbResourceId = resourceImage.OriginResourceId
				o.Update(resourceImage)
			}

		}
		return nil, nil
	})

	self.String("helloworld")
}