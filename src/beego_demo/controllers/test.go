package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (o *TestController) Get() {
	id := o.Ctx.Input.Param(":id")
	o.Data["json"] = id
	o.ServeJSON()
}

func (o *TestController) List() {
	o.Data["json"] = o.Ctx.Request.Method
	o.ServeJSON()
}

func (o *TestController) Test() {
	o.Data["json"] = "test"
	o.ServeJSON()
}

