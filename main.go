package main

import (
	_ "BeegoDemo/routers"
	//	"fmt"
	"BeegoDemo/models"

	"github.com/astaxie/beego"
)

func init() {
	//	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	//	beego.SetStaticPath("/views", "views")
	//路由不区分大小写
	beego.BConfig.RouterCaseSensitive = false
	models.Init()
}

func main() {

	beego.Run()
}
