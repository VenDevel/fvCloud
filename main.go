package main

import (
	_ "fvCloud/routers"
	_ "fvCloud/sqlite"
	"github.com/astaxie/beego"
)

func main() {
	//sqlite.AddUser("admin", "admin", "admin", 1)
	beego.Run()
}
