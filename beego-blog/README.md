## 基于go + beego + gorm开发的博客
## beego小项目文件结构
├─.idea
├─conf
├─controllers
├─models
├─routers
├─static
│  ├─css
│  ├─img
│  └─js
├─tests
└─views


## 运行
打开终端：bee run 
或者运行main.go  


## 新增

default.go

```go
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
```

## 查询

```go
	/**
	  查询
	*/
	// 1，orm对象
	o := orm.NewOrm()
	// 2，查询的对象
	user := models.User{}
	// 3，指定查询对象的值
	user.Id = 10
	// 4， 查询
	err := o.Read(&user)
	if err != nil {
		beego.Info("查询失败", err)
		return
	}
	beego.Info("查询成功", user)
```

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220418225725.png)

## 修改

```go
	/* 修改
	1，先查询存不存在
	2，存在则修改
	*/
	o := orm.NewOrm()
	user := models.User{}
	user.Id = 10
	err := o.Read(&user)
	if err == nil {
		user.Username = "邵桐杰"
		user.Name = "公众号：千羽的编程时光"
		_, err := o.Update(&user)
		if err != nil {
			beego.Info("修改失败", user)
			return
		}
	}

```

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220418225457.png)

## 删除

```go
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
```

