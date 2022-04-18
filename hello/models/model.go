package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	Username string
	Password string
	Name     string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Local")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}
