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
	var audioComment = &controllers.AudioCommentController{}

	var audioCommentReply = &controllers.AudioCommentReplyController{}

	var collection = &controllers.CollectionController{}

	var dashborad = &controllers.DashboardContrlller{}

	var collectionFolder = &controllers.AudioCollectionFolderController{}

	var love = &controllers.AudioUserLoveController{}

	var page = &controllers.PageController{}

	beego.Router("/", &controllers.MainController{})
	beego.Router("/test/json", test_controller, "*:JsonResult")
	beego.Router("/test/xml", test_controller, "*:XmlResult");
	beego.Router("/test/page/:page", test_controller, "*:FindHtml")

	beego.Router(app.GetUrl("/music/player"), audio, "*:Player")

	beego.Router(app.GetUrl("/user/new"), user, "*:NewToDb")
	beego.Router(app.GetUrl("/user/new/test"), user, "*:UploadTest")
	/**
		这下面便是输出页面的方法
	 */

	//主要的操作便是主页面操作
	beego.Router(app.GetUrl("/page/index/main"), audio, "*:IndexPage")
	//输出瀑布流视图
	beego.Router(app.GetUrl("/page/pbl/main"), audio, "*:AudioPblPage")

	//登录页面
	beego.Router(app.GetUrl("/page/mit/login"), login, "*:LoginPage")

	//注册页面
	beego.Router(app.GetUrl("/page/mit/register"), login, "*:RegisterPage")

	//对应的登录注册协议的页面
	beego.Router(app.GetUrl("/page/mit/policy"), login, "*:PolicePage")

	//对应的 index 页面
	beego.Router(app.GetUrl("/page/mit/index"),page , "*:MitIndexPage")

	//上传信息的页面
	beego.Router(app.GetUrl("/page/upload/audio"), audio, "*:AudioUploadPage")

	//修改音频信息的页面
	beego.Router(app.GetUrl("/page/upload/update"), audio, "*:AudioUpdatePage")

	//详细页面
	beego.Router(app.GetUrl("/page/details/main"),audioComment , "*:PageDetails")

	/**
		对应的 api的接口
	 */
	//利用账号密码进行注册
	/*
		我认为相对应的注册方法与登录方法应该一样
		用户应该可以支持account ， wx ， 支付宝等等方式 ， 来进行创建
		虽然，account 的方法 ， 可以说应该是应该是 其他注册方法的基石 ， 但是如何设计那个是后话了

	*/
	beego.Router(app.GetUrl("/api/register/by/account"), login, "*:RegisterByAccount")
	//进行登录的api
	beego.Router(app.GetUrl("/api/login/by/account"), login, "*:LoginByAccount")
	//获取当前session中的user
	beego.Router(app.GetUrl("/api/login/current/user"), login, "*:CurrentUser")
	//根据对应的id来获取user信息
	beego.Router(app.GetUrl("/api/user/find/id"),user , "*:FindUserById")
	//进行对应的上传页面的 api
	beego.Router(app.GetUrl("/api/audio/upload"), audio, "*:AudioUpload")

	//对应的修改页面的api
	beego.Router(app.GetUrl("/api/audio/update"), audio, "*:AudioUpdate")

	//进行目标用户下的 所有音频文件的信息
	beego.Router(app.GetUrl("/api/audio/user"), audio, "*:SearchAudioByUserId")

	//删除目标的audio 的 信息
	beego.Router(app.GetUrl("/api/audio/delete"), audio, "*:Delete")

	//下面 进行对应的搜索
	beego.Router(app.GetUrl("/api/audio/search"), audio, "*:SearchByString")
	beego.Router(app.GetUrl("/api/audio/favorite"), audio, "*:Favorites")
	beego.Router(app.GetUrl("/api/audio/find/id"), audio, "*:FindAudioById")

	//获取某个audio的 评论
	beego.Router(app.GetUrl("/api/comment/find/audio"),audioComment,"*:ByAudioId")

	beego.Router(app.GetUrl("/api/comment/insert"),audioComment,"*:InsertComment")

	//书写对应的评论的回复
	beego.Router(app.GetUrl("/api/reply/comment/insert"), audioCommentReply,"*:InsertCommentReply")
	beego.Router(app.GetUrl("/api/reply/comment/find/comment") , audioCommentReply , "*:ByCommentId")
	//收藏的接口
	beego.Router(app.GetUrl("/api/collection/insert") ,collection ,"*:InsertCollection")
	//取消收藏的接口
	beego.Router(app.GetUrl("/api/collection/delete"),collection,"*:DeleteCollection")
	//根据对应的信息， 来获取对应的信息
	beego.Router(app.GetUrl("/api/collection/find/relationship"), collection , "*:FindByUserAndAudio")

	//根据对应的 收藏夹的 id 来搜索出对应的 收藏信息与其绑定的 audio 信息
	beego.Router( app.GetUrl("/api/collection/all/and"),collection,"*:SearchCollectionAndAudio")

	//转移收藏信息
	beego.Router( app.GetUrl("/api/collection/update"),collection , "*:UpdateCollection")
	//仪表板输出格式相关的接口
	beego.Router(app.GetUrl("/api/dashborad/audio") , dashborad , "*:FindByAudioId")

	//对应的文件夹
	beego.Router( app.GetUrl("/api/collection/folder/all"), collectionFolder , "*:AllByUserId")

	//新建收藏夹
	beego.Router( app.GetUrl("/api/collection/folder/insert"),collectionFolder,"*:AddCollectionFolder")

	//更改收藏夹的对应的信息
	beego.Router( app.GetUrl( "/api/collection/folder/update") , collectionFolder ,"*:UpdateCollectionFolder")
	//删除收藏夹
	beego.Router( app.GetUrl("/api/collection/folder/delete") , collectionFolder , "*:DeleteCollectionFolder");
	//点赞
	beego.Router( app.GetUrl("/api/love/insert"),love,"*:InsertLove")
	//取消点赞
	beego.Router( app.GetUrl( "/api/love/delete"),love,"*:DeleteLove")
	//查询点赞关系
	beego.Router( app.GetUrl("/api/love/find"),love,"*:FindLove")

	//api测试输出
	beego.Router(app.GetUrl("/api/test"),test_controller,"*:ApiTest")

	/*
		对应的 web 的接口， 该接口的主要目的 ， 只是服务与 web 前端服务 交互的前端
	*/

	//在用户的模拟下 ， 开始获取对应的信息
	//beego.Router(app.GetUrl("/web/audio/audio/user") , audio,"")
	beego.Router(app.GetUrl("/web/test"),test_controller,"*:WebTest")

	//下面是测试api

	//测试对应的上传图片的功能是否有效
	beego.Router(app.GetUrl("/test/upload/image"), test_controller, "*:UploadImage")

	beego.Router(app.GetUrl("/test/session/set"), test_controller, "*:SetSession")
	beego.Router(app.GetUrl("/test/session/get"), test_controller, "*:GetSession")

	beego.Router(app.GetUrl("/test/audio/comments"),test_controller,"*:ApiAudioComment")

	beego.Router(app.GetUrl("/test/resource/redirec"),test_controller,"*:TestResourceRedirect")
}
