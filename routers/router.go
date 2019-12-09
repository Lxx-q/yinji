package routers

import (
	"yinji/controllers"
	"github.com/astaxie/beego"
	"yinji/service"
)

func init() {


	var app = service.AppService{"yinji"};

	var test_controller = &controllers.TestController{}

	var audio = &controllers.AudioController{};

	var user = &controllers.UserController{}

	var login = &controllers.LoginController{}


    beego.Router("/", &controllers.MainController{})
    beego.Router("/test/json" , test_controller , "*:JsonResult")
	beego.Router("/test/xml",test_controller,"*:XmlResult");
	beego.Router("/test/page/:page",test_controller,"*:FindHtml")

	beego.Router(app.GetUrl("/music/player"),audio , "*:Player")

	beego.Router(app.GetUrl("/user/new") , user , "*:NewToDb")
	beego.Router(app.GetUrl("/user/new/test"),user,"*:UploadTest")

	beego.Router(app.GetUrl("/upload/audio"),audio , "*:AudioUploadPage")

	//进行对应音频页面更改信息的操作
	beego.Router(app.GetUrl("/upload/update"),audio,"*:AudioUpdatePage")

	/**
		这下面便是输出页面的方法
	 */

	//主要的操作便是主页面操作
	beego.Router(app.GetUrl("/page/index/main"),audio,"*:IndexPage")
	//输出瀑布流视图
	beego.Router(app.GetUrl("/page/pbl/main"),audio,"*:AudioPblPage")

	/**
		对应的 api的接口
	 */

	 //进行登录的api

	 beego.Router(app.GetUrl("/api/login/by/account"),login , "*:LoginByAccount")
	//进行对应的上传页面的 api
	beego.Router(app.GetUrl("/api/audio/upload") ,audio , "*:AudioUpload")

	//对应的修改页面的api
	beego.Router( app.GetUrl("/api/audio/update") , audio , "*:AudioUpdate")

	//进行目标用户下的 所有音频文件的信息
	beego.Router(app.GetUrl("/api/audio/user"),audio,"*:SearchAudioByUserId")

	//删除目标的audio 的 信息
	beego.Router(app.GetUrl("/api/audio/delete") , audio , "*:Delete")

	//下面 进行对应的搜索
	beego.Router(app.GetUrl("/api/audio/search"),audio,"*:SearchByString")
	beego.Router(app.GetUrl("/api/audio/favorite"),audio,"*:Favorites")
	beego.Router( app.GetUrl("/api/audio/find/id") , audio ,"*:FindAudioById")
	//下面是测试api

	//测试对应的上传图片的功能是否有效
	beego.Router( app.GetUrl("/test/upload/image") , test_controller , "*:UploadImage")

	 beego.Router(app.GetUrl("/test/session/set") , test_controller , "*:SetSession")
	beego.Router(app.GetUrl("/test/session/get") , test_controller , "*:GetSession")



}
