package controllers

import (
	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller
}

func (this *baseController) WriteList(code int, msg string, list interface{}) {
	this.Data["json"] = map[string]interface{}{
		"Code": code,
		"Msg":  msg,
		"List": list,
	}
	this.ServeJSON()
}

func (this *baseController) WriteData(code int, msg string, data interface{}) {
	this.Data["json"] = map[string]interface{}{
		"Code": code,
		"Msg":  msg,
		"data": data,
	}
	this.ServeJSON()
}
