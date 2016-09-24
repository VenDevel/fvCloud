package controllers

import (
	"fmt"
	"fvCloud/models"
	"fvCloud/sqlite"
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

func (this *UserControllers) Register() {
	user := &models.AddUserInfo{}
	_, err := this.GetRequestJsonToObject(user)
	if err != nil {
		this.WriteData(-1, err.Error(), "")
		return
	}
	if user.Account == "" || user.Password == "" || user.NickName == "" {
		this.WriteData(-2, "参数错误", "")
		return
	}
	err = sqlite.AddUser(user.Account, user.Password, user.NickName, 1)
	if err != nil {
		this.WriteData(-3, err.Error(), "")
		return
	}
	this.WriteData(1000, "SUCCESS", "")
}

func (this *UserControllers) Login() {
	user := &models.AddUserInfo{}
	_, err := this.GetRequestJsonToObject(user)
	if err != nil {
		this.WriteData(-1, err.Error(), "")
		return
	}
	if user.Account == "" || user.Password == "" {
		this.WriteData(-2, "参数错误", "")
		return
	}
	us, err := sqlite.GetUserByAccount(user.Account)
	if err != nil {
		fmt.Println(err.Error())
		this.WriteData(-3, "", "") //一般是账号不存在
		return
	}

	if us.Password != user.Password {
		this.WriteData(-4, "密码不正确", "")
		return
	}

	this.WriteData(1000, "SUCCESS", "")
}

func (this *UserControllers) GTest() {

}

func (this *UserControllers) PTest() {
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
