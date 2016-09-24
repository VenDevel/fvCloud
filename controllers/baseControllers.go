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

func (this *baseController) GetRequestJsonToObject(object interface{}) (interface{}, error) {
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &object)
	return object, err
}

func (this *baseController) GetRequestJsonToMap() (map[string]interface{}, error) {
	m := make(map[string]interface{}, 0)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &m)
	return m, err
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
