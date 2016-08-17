package controllers

import (
	"fmt"
)

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
	body := this.GetRequestBody()
	fmt.Println(this.Query("password"))
	fmt.Println("aaaaaa")
	fmt.Println(string(body))
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}
