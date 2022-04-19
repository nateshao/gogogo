package main

import (
	_ "beego-blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.AddFuncMap("ShowNextPage", getNextPage)
	beego.AddFuncMap("ShowLastPage", getLastPage)
	beego.Run()
}

func getNextPage(pageindex int) int {
	pageindex++
	return pageindex
}
func getLastPage(pageindex int) int {
	if pageindex--; pageindex < 0 {
		pageindex = 0
	}
	return pageindex
}
