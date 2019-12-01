package main

import (
	_ "yinji/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
	_ "github.com/go-sql-driver/mysql"
	"yinji/service"
)


func Init(){
	RegistOrm();
}

func RegistOrm(){

	// 进行注册  相对应的 信息
	orm.RegisterDriver("mysql",orm.DRMySQL);

	/**
		var maxId =
		var maxConnection = 30;
	 */
	orm.RegisterDataBase("default","mysql","root:980621@/none?charset=utf8&loc=Local")
	//设置 相对应的 信息 内容
	orm.RegisterModel(
		new(bean.Login),
		new(bean.User),
		new(bean.Audio),
		);

	orm.Debug = true;

}

func TemplateInit(){
}

func main() {
	Init();
	//设置相对应的 静态文件的 代码
	var app = service.AppService{"yinji"}
	beego.SetStaticPath(app.GetUrl("/music"),"static/music")
	beego.SetStaticPath(app.GetUrl("/index") , "static/index")
	beego.SetStaticPath(app.GetUrl("/resources") , "static/resources")
	beego.SetStaticPath(app.GetUrl("/framework"),"static/framework")
	beego.SetStaticPath(app.GetUrl("/upload"), "static/upload")
	beego.SetStaticPath(app.GetUrl("/audio"),"static/audio")
	beego.SetStaticPath(app.GetUrl("/image"),"static/image")
	beego.SetStaticPath(app.GetUrl("/js"),"static/js")

	//瀑布流界面 ， 音乐显示的布局
	beego.SetStaticPath(app.GetUrl("/page/pbl"),"static/page/pbl")

	beego.Run();
}

