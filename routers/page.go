package routers

import (
	"github.com/astaxie/beego"
	"bytes"
)

type PageController struct {

	beego.Controller

}

func ( controller *PageController) UsingPageByName(){

	var suffix = ".html";

	var head = "";

	var buffer = bytes.Buffer{};

	var input = controller.Ctx.Input;

	var page = input.Param("page");

	buffer.WriteString(head)

	buffer.WriteString( page )

	buffer.WriteString(suffix)

	controller.TplName = buffer.String();
}
