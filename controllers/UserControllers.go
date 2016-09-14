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

func (this *UserControllers) RegisterPage() {
	this.TplName = "register.html"
}

func (this *UserControllers) Login() {
	//this.WriteData(1000, "test", test)
	// body := this.GetRequestBody()
	// fmt.Println(this.Query("password"))
	// fmt.Println("aaaaaa")
	// fmt.Println(string(body))
	// this.Data["Website"] = "beego.me"
	// this.Data["Email"] = "astaxie@gmail.com"
	// this.TplName = "index.tpl"

	account := this.GetRequestParam("account")
	password := this.GetRequestParam("password")
	test := this.GetRequestParam("11111")

	fmt.Println("a----------")

	fmt.Println(account)
	fmt.Println(password)
	fmt.Println(test)

	fmt.Println("a----------")
	mp := this.GetRequestParams()
	by := this.GetRequestBody()
	qusrt := this.Query("account")
	fmt.Println("b----------")
	fmt.Println(mp)
	fmt.Println(string(by))
	fmt.Println(qusrt)
	fmt.Println("b----------")

	this.WriteData(1000, "test", 1)
}
