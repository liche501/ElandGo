package controllers

import (
	//	"encoding/json"
	//	"strconv"
	//	"time"
	//	"reflect"
	"BeegoDemo/models"
	"BeegoDemo/tools/libs"

	//	"github.com/astaxie/beego"
	//	. "github.com/astaxie/beego/utils"
)

type UserController struct {
	//	beego.Controller
	BaseController
}

//============
func (this *UserController) CreateUser() {
	//	beego.Alert(this.Ctx.Request.Header.Get("Content-Type"))

	var user models.User
	user.UserId = this.GetString("UserId")
	user.UserName = this.GetString("UserName")
	password := this.GetString("Password")
	password = libs.Md5([]byte(password))
	//	beego.Alert(password)
	user.Password = password
	if this.GetString("UseChk") == "true" {
		user.UseChk = true
	} else {
		user.UseChk = false
	}

	if this.GetString("AuthChk") == "true" {
		user.AuthChk = true
	} else {
		user.AuthChk = false
	}
	user.InUserId = this.userId
	//	beego.Alert(time.Now().Format("2006-01-02 15:04:05"))
	err := user.CreateUser()
	if err == nil {
		this.Ctx.WriteString("success")
	} else {
		this.Ctx.WriteString("error")
	}
}

func (this *UserController) UserGetById() {
	userId := this.GetString("UserId")

	if userResults, err := models.UserGetById(userId); err == nil {

		this.jsonResult(userResults)

	} else {
		this.ajaxMsg("UserGetById请求错误", 500)
	}
}

func (this *UserController) UpdateUser() {
	user := models.User{}
	user.UserId = this.GetString("UserId")
	user.UserName = this.GetString("UserName")
	//	user.Password = "1234_modify"
	if this.GetString("UseChk") == "true" {
		user.UseChk = true
	} else {
		user.UseChk = false
	}
	if this.GetString("AuthChk") == "true" {
		user.AuthChk = true
	} else {
		user.AuthChk = false
	}
	//	user.ModiUserId = this.GetString("UserId")
	user.ModiUserId = this.userId

	if err := user.UpdateUser(); err == nil {
		this.ajaxMsg("UserupDate成功", 200)
	} else {
		this.ajaxMsg("UserupDate失败", 500)
	}
}
