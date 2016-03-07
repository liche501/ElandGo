package controllers

import (
	"BeegoDemo/models"
	"BeegoDemo/tools/libs"
	"strings"

	"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["UserName"] = c.userName
	c.TplName = "index.tpl"
}

//func (this *MainController) Demo1() {
//	this.Data["username"] = "liche"
//	this.TplName = "demo1.tpl"
//}
func (this *MainController) Demo1() {
	//	beego.Alert(this.Ctx.Request.URL.Query())
	if this.isPost() {
		//		this.Redirect("/user/index", 302)
		this.redirect("/user/index")
	}
	this.Data["username"] = "liche"
	this.TplName = "demo1.tpl"

}

func (this *MainController) Login() {
	beego.Alert(this.userId)
	if len(this.userId) > 0 {
		beego.Alert(33333333)
		this.redirect("/")
	}
	beego.ReadFromRequest(&this.Controller)

	if this.isPost() {
		flash := beego.NewFlash()

		userid := strings.TrimSpace(this.GetString("userid"))
		password := strings.TrimSpace(this.GetString("password"))
		remember := this.GetString("remember")
		if userid != "" && password != "" {

			user, err := models.UserGetById(userid)
			errorMsg := ""
			beego.Alert(userid, password)
			//			aaa := libs.Md5([]byte("admin888"))
			//			beego.Alert(aaa)
			if err != nil || user.Password != libs.Md5([]byte(password)) {
				errorMsg = "帐号或密码错误"
			} else if user.UseChk == false {
				errorMsg = "该帐号已禁用"
			} else {

				authkey := libs.Md5([]byte(this.getClientIp() + "|" + user.Password))
				if remember == "yes" {
					this.Ctx.SetCookie("auth", user.UserId+"|"+authkey, 7*86400)
				} else {

					beego.Alert(authkey)
					this.Ctx.SetCookie("auth", user.UserId+"|"+authkey)
				}
				msg := user.UserName + ":登录成功"
				this.ajaxMsg(msg, 200)
			}
			beego.Alert(errorMsg)
			flash.Error(errorMsg)
			flash.Store(&this.Controller)
			this.redirect("/login")
		}
	}

	this.TplName = "login.html"
}

func (this *MainController) Logout() {
    
	beego.Warning("logout......")

	this.Ctx.SetCookie("auth", "")

	this.redirect("/login")
}
