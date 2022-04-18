package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"hello/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	o := orm.NewOrm()
	user := models.User{}
	user.Name = "111"
	user.Username = "千羽"
	user.Password = "1234"
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入失败", err)
		return
	}
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.Data["data"] = "千羽的编程时光"
	c.TplName = "test.html"

}

func (c *MainController) Post() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.Data["data"] = "Go开发的编程时光"
	c.TplName = "test.html"

}
