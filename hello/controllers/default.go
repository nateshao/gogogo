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
	/**新增*/
	//o := orm.NewOrm()
	//user := models.User{}
	//user.Name = "111"
	//user.Username = "千羽"
	//user.Password = "1234"
	//_, err := o.Insert(&user)
	//if err != nil {
	//	beego.Info("插入失败", err)
	//	return
	//}

	/**查询*/
	//// 1，orm对象
	//o := orm.NewOrm()
	//// 2，查询的对象
	//user := models.User{}
	//// 3，指定查询对象的值
	//user.Id = 10
	//// 4， 查询
	//err := o.Read(&user)
	//if err != nil {
	//	beego.Info("查询失败", err)
	//	return
	//}
	//beego.Info("查询成功", user)

	/* 修改
	1，先查询存不存在
	2，存在则修改
	*/
	//o := orm.NewOrm()
	//user := models.User{}
	//user.Id = 10
	//err := o.Read(&user)
	//if err == nil {
	//	user.Username = "邵桐杰"
	//	user.Name = "公众号：千羽的编程时光"
	//	_, err := o.Update(&user)
	//	if err != nil {
	//		beego.Info("修改失败", user)
	//		return
	//	}
	//}

	/*删除*/
	o := orm.NewOrm()
	user := models.User{}
	user.Id = 1
	_, err := o.Delete(&user)
	if err != nil {
		beego.Info("删除失败", err)
		return
	}
	beego.Info("删除成功", user)

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
