package controllers

import (
	//	"encoding/json"
	"strconv"
	//	"time"
	"BeegoDemo/models"
	//	"reflect"

	//	"strings"

	"github.com/astaxie/beego"
	//	. "github.com/astaxie/beego/utils"
)

type BookmarkController struct {
	//	beego.Controller
	BaseController
}

//============
func (this *BookmarkController) CreateBookmark() error {
	beego.Alert(this.Ctx.Request.Header.Get("Content-Type"))

	var bookmark models.Bookmark
	bookmark.BookmarkName = this.GetString("BookmarkName")
	bookmark.Desc = this.GetString("Desc")
	if this.GetString("FixChk") == "true" {
		bookmark.FixChk = true
	} else {
		bookmark.FixChk = false
	}

	if this.GetString("NumChk") == "true" {
		bookmark.NumChk = true
	} else {
		bookmark.NumChk = false
	}
	bookmark.Placeholder = this.GetString("Placeholder")

	if this.GetString("UseChk") == "true" {
		bookmark.UseChk = true
	} else {
		bookmark.UseChk = false
	}
	bookmark.InUserId = this.GetString("InUserId")
	err := bookmark.CreateBookmark()
	if err == nil {
		this.Ctx.WriteString("success")
	} else {
		this.Ctx.WriteString("error")
	}
	return err
}

func (this *BookmarkController) GetBookmarkByName() {
	name := this.GetString("Name")
	beego.Alert(name)
	bookmark := models.Bookmark{}

	if bookmarkResults, err := bookmark.GetBookmarkByName(name); err == nil {
		this.jsonResult(bookmarkResults)
	} else {
		this.ajaxMsg("GetBookmarkByName请求错误", 500)
	}
}

func (this *BookmarkController) DeleteBookmarkById() error {
	beego.Alert(this.Ctx.Request.Header.Get("Content-Type"))

	var bookmark models.Bookmark

	bookmarkId, err := strconv.Atoi(this.GetString("BookmarkId"))
	beego.Alert(bookmarkId)

	if err != nil {
		return err
	}

	err = bookmark.DeleteBookmarkById(bookmarkId)

	if err == nil {
		this.Ctx.WriteString("success")
	} else {
		this.Ctx.WriteString("error")
	}
	return err
}

func (this *BookmarkController) UpdateBookmarkById() error {
	beego.Alert(this.Ctx.Request.Header.Get("Content-Type"))

	var bookmark models.Bookmark

	bookmarkId, err := strconv.Atoi(this.GetString("BookmarkId"))

	if err != nil {
		return err
	}
	bookmark.BookmarkId = bookmarkId

	bookmark.BookmarkName = this.GetString("BookmarkName")
	bookmark.Desc = this.GetString("Desc")
	if this.GetString("FixChk") == "true" {
		bookmark.FixChk = true
	} else {
		bookmark.FixChk = false
	}

	if this.GetString("NumChk") == "true" {
		bookmark.NumChk = true
	} else {
		bookmark.NumChk = false
	}
	bookmark.Placeholder = this.GetString("Placeholder")

	if this.GetString("UseChk") == "true" {
		bookmark.UseChk = true
	} else {
		bookmark.UseChk = false
	}
	bookmark.InUserId = this.GetString("InUserId")
	err = bookmark.UpdateBookmark()
	if err == nil {
		this.Ctx.WriteString("success")
	} else {
		this.Ctx.WriteString("error")
	}
	return err
}
