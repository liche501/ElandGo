package routers

import (
	"BeegoDemo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/demo1", &controllers.MainController{}, "*:Demo1")
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.BookmarkController{})

	//	beego.Include(&controllers.UserController{})

}
