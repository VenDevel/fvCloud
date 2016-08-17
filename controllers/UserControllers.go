package controllers

import ()

type UserControllers struct {
	baseController
}

func (this *UserControllers) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}

func (this *UserControllers) LoginPage() {
	this.TplName = "login.html"
}

func (this *UserControllers) Login() {
	//this.WriteData(1000, "test", test)
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}
