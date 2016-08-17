package routers

import (
	"fvCloud/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	initUserRouter()

}

func initUserRouter() {
	beego.Router("/login.html", &controllers.UserControllers{}, "get:LoginPage")
	beego.Router("/login.html", &controllers.UserControllers{}, "post:Login")
}
