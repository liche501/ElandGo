package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	. "github.com/astaxie/beego/utils"
)

type Bookmark struct {
	BookmarkId   int
	BookmarkName string
	Desc         string
	FixChk       bool
	NumChk       bool
	Placeholder  string
	UseChk       bool
	InUserId     string
	InDatetime   time.Time
	ModiUserId   string
	ModiDatetime time.Time
}

func init() {
	fmt.Println("======================models.Bookmark.init 已执行")
}

func (this *Bookmark) CreateBookmark() error {
	o := orm.NewOrm()

	_, err := o.Raw(`INSERT INTO bookmark(bookmark_name, desc, fix_chk, num_chk, placeholder, use_chk, in_user_id, in_datetime) values(?,?,?,?,?,?,?,?)`, this.BookmarkName, this.Desc, this.FixChk, this.NumChk, this.Placeholder, this.UseChk, this.InUserId, time.Now()).Exec()

	if err == nil {
		fmt.Println("bookmark create success")
	}
	return err
}

func (this *Bookmark) GetBookmarkByName(bookmarkName string) ([]Bookmark, error) {
	var bookmarkList []Bookmark

	o := orm.NewOrm()

	_, err := o.Raw("select * from bookmark where bookmark_name = ? ", bookmarkName).QueryRows(&bookmarkList)
	if err == nil {
		Display("GetBookmarkByName_Model==", bookmarkList)
	}

	return bookmarkList, err
}

func (this *Bookmark) DeleteBookmarkById(bookmarkId int) error {
	o := orm.NewOrm()

	res, err := o.Raw(`DELETE FROM bookmark WHERE bookmark_id = ?`, bookmarkId).Exec()

	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Printf("%d rows of bookmark delete success", num)
	}

	return err
}

func (this *Bookmark) UpdateBookmark() error {
	o := orm.NewOrm()

	_, err := o.Raw(`Update bookmark SET bookmark_name = ?, desc = ?, fix_chk = ?, num_chk = ?, placeholder = ?, use_chk = ?, modi_user_id = ?, modi_datetime = ? WHERE bookmark_id = ?`, this.BookmarkName, this.Desc, this.FixChk, this.NumChk, this.Placeholder, this.UseChk, this.InUserId, time.Now(), this.BookmarkId).Exec()

	if err == nil {
		fmt.Println("bookmark update success")
	}
	return err
}
