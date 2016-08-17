package controllers

import (
	"encoding/json"
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

func (this *baseController) GetRequestBody() []byte {
	return this.Ctx.Input.RequestBody
}

func (this *baseController) GetRequestBodyToJson(object interface{}) error {
	return json.Unmarshal(this.Ctx.Input.RequestBody, &object)
}

func (this *baseController) GetRequestParam(key string) string {
	return this.Ctx.Input.Param(key)
}

func (this *baseController) GetRequestParams() map[string]string {
	return this.Ctx.Input.Params()
}

func (this *baseController) Query(key string) string {
	return this.Ctx.Input.Query(key)
}
