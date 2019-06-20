package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (o *MainController) Get() {
	o.Data["json"] = "hello world"
	o.ServeJSON()
}
