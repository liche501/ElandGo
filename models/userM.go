package models

import (
	"fmt"
	//	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "github.com/astaxie/beego/utils"
)

type User struct {
	UserId       string `orm:"pk"`
	UserName     string
	Password     string
	UseChk       bool
	AuthChk      bool
	InUserId     string
	InDatetime   time.Time `orm:"auto_now_add"`
	ModiUserId   string
	ModiDatetime time.Time `orm:"auto_now"`
}

func init() {
	fmt.Println("======================models.user.init 已执行")
	Display("init", "==")
	// 需要在init中注册定义的model

	orm.RegisterModel(new(User))

}

func (this *User) CreateUser() error {
	o := orm.NewOrm()
	//	_, err := o.Raw(`insert into user(user_id,user_name,password,use_chk,auth_chk,in_user_id,in_datetime)
	//					values(?,?,?,?,?,?,?)`, this.UserId, this.UserName, this.Password, this.UseChk, this.AuthChk, this.UserId, time.Now().UTC()).Exec()
	//	if err == nil {

	if _, err := o.Insert(this); err == nil {
		return nil
	} else {
		return err
	}

}

func (this *User) UpdateUser() error {
	o := orm.NewOrm()
	//	fmt.Println(time.Now())
	//	orm.DefaultTimeLoc
	//	fmt.Println(this.ModiDatetime.Location())

	if num, err := o.Update(this, "user_name", "use_chk", "auth_chk", "modi_user_id"); err == nil {
		beego.Debug("UpdateUser更新的行数为:", num)
		return nil
	} else {
		beego.Error("UpdateUser更新失败:", err)
		return err
	}

}

//func (this *User) ChangeUseChk() error {
//	o := orm.NewOrm()
//	if num, err := o.Update(this, "use_chk"); err == nil {
//		beego.Debug("ChangeUseChk更新的行数为:", num)
//		return nil
//	} else {
//		beego.Error("ChangeUseChk更新失败:", err)
//		return err
//	}

//}

func UserGetById(userId string) (*User, error) {
	//	var userList []User
	//	beego.Alert(this.UserId)
	var userList User
	//	var userList []orm.Params
	o := orm.NewOrm()
	//	_, err := o.Raw("select *from user ").QueryRows(&userList)
	err := o.Raw("select *from user where user_id = ? ", userId).QueryRow(&userList)
	if err == nil {

		//		Display("UserGetById_Model==", userList)
		//				beego.Alert(userList[0]["user_name"])
		//		beego.Alert(userList.Password)
		return &userList, nil
	} else {

		return &userList, err
	}
}
