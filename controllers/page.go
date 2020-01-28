package controllers

type PageController struct {
	BeegoController
}

func ( self *PageController ) MitIndexPage() {
	self.Resource("mit/index.html")
}