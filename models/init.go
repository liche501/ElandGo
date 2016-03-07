package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func Init() {
	orm.RegisterDataBase("default", "sqlite3", "eland.sqlite")

	//	orm.RegisterModel(new(User), new(Task), new(TaskGroup), new(TaskLog))
	//	orm.RunSyncdb("default", true, true)
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
