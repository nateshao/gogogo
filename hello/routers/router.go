package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/demo", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{})

	beego.Router("/login", &controllers.MainController{}, "get:ShowLogin;post:HandleLogin")
}
