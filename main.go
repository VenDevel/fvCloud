package main

import (
	_ "fvCloud/routers"
	"fvCloud/sqlite"
	"github.com/astaxie/beego"
)

func main() {
	sqlite.AddCategory("aafads")
	beego.Run()
}
