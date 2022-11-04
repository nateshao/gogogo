---
title: day5-设计模式之Database-SQL-与GORM实践
date: 2022-05-15 12:40:59
tags: 
- Go学习路线
- 字节跳动青训营
---



这是我参与「第三届青训营 -后端场」笔记创作活动的的第5篇笔记

## 「设计模式之 Database/SQL 与 GORM 实践」 第三届字节跳动青训营

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220709231205.jpg)

同时这也是课表的第5天课程

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220703104730.png)

PC端阅读效果更佳，点击文末：**阅读原文**即可。

# 01 理解database/sql

基本用法，设计原理，基础概念

## 1.1基本用法-Quick Start

```go
package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)

func main() {
	db, err := sql.Open ("mysqL", "user:password@tcp(127.0.0.1:3306])/hello")
	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err == nil {
		//xxx
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user. ID, &user,Name)
		if err!=nil{
		//..
		}
	}
		users = append(users, user)
		if rows.Err() != nil {
		}
}
```

了解 DSN 是什么：

- https://github.com/go-sql-driver/mysql#dsn-data-source-name
- https://en.wikipedia.org/wiki/Data_source_name

## 1.2 设计原理

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515171109.png)



查看go的源码

```go
type DB struct {
	// Atomic access only. At top of struct to prevent mis-alignment
	// on 32-bit platforms. Of type time.Duration.
	waitDuration int64 // Total time waited for new connections.

	connector driver.Connector
	// numClosed is an atomic counter which represents a total number of
	// closed connections. Stmt.openStmt checks it before cleaning closed
	// connections in Stmt.css.
	numClosed uint64

	mu           sync.Mutex // protects following fields
	freeConn     []*driverConn
	connRequests map[uint64]chan connRequest
	nextRequest  uint64 // Next key to use in connRequests.
	numOpen      int    // number of opened and pending open connections
	// Used to signal the need for new connections
	// a goroutine running connectionOpener() reads on this chan and
	// maybeOpenNewConnections sends on the chan (one send per needed connection)
	// It is closed during db.Close(). The close tells the connectionOpener
	// goroutine to exit.
	openerCh          chan struct{}
	resetterCh        chan *driverConn
	closed            bool
	dep               map[finalCloser]depSet
	lastPut           map[*driverConn]string // stacktrace of last conn's put; debug only
	maxIdle           int                    // zero means defaultMaxIdleConns; negative means 0
	maxOpen           int                    // <= 0 means unlimited
	maxLifetime       time.Duration          // maximum amount of time a connection may be reused
	cleanerCh         chan struct{}
	waitCount         int64 // Total number of connections waited for.
	maxIdleClosed     int64 // Total number of connections closed due to idle.
	maxLifetimeClosed int64 // Total number of connections closed due to max free limit.

	stop func() // stop cancels the connection opener and the session resetter.
}
```

接下来看连接池

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515174512.png)

连接池配置

```go
func (db *DB) SetConnMaxIdleTime(d time.Duration)
func (db *DB) SetConnMaxLifetime(d time.Duration)
func (db *DB) SetMaxIdleConns(n int)
func (db *DB) SetMaxOpenConns(n int)
```

连接池状态

```go
func (db *DB) Stats() DBStats
```

Driver 连接接口

```go
	for i := 0; i < maxBadConnRetries; i++ {
		// 从连接池获取连接或通过driver新建连接
		dc, err := db.conn(ctx, strategy)
		//有空闲连接-> reuse -> max life time
		//新建连接-> max open.. .
		//将连接放回连接池
		defer dc.db.putConn(dc, err, true)
		/// validateConnection有无错误
		// max life time, max idle conns检查
		//连接实现driver.Queryer; driver.Execer等interface
		if err == nil{
			err = dc.ci.Query(sql, args...)
		}
		isBadConn = errors.Is(err, driver. ErrBadConn)
		if !isBadConn {
			break
		}
```

```go
// Open new Connection.
// See https://github.com/go-sql-driver/mysql#dsn-data-source-name for how
// the DSN string is formatted
func (d MySQLDriver) Open(dsn string) (driver.Conn, error) {
	cfg, err := ParseDSN(dsn)
	if err != nil {
		return nil, err
	}
	c := &connector{
		cfg: cfg,
	}
	return c.Connect(context.Background())
}
// github.com/go-sql-driver/mysql/driver.go
// 注册驱动
func init() {
	sql.Register("mysql", &MySQLDriver{})
}
```

https://dev.mysql.com/doc/internals/en/client-server-protocol.html

# 02 GORM基础使用

基本用法，Model定义，惯例约定，关联操作

## 2.1背景知识

**设计简洁、功能强大、自由扩展的全功能ORM**

- 设计原则： API精简、测试优先、最小惊讶、灵活扩展、无依赖可信赖

- 功能完善：
  - 关联：一对一、一对多、单表自关联、多态；Preload、 Joins 预加载、级联删除；关联模式；自定义关联表
  - 事务：事务代码块、嵌套事务、Save Point
  - 多数据库、读写分离、命名参数、Map、子查询、分组条件、代码共享、SQL表达式(查询、创建、更新)、自动选字段、查询优化器
  - 字段权限、软删除、批量数据处理、Prepared Stmt、自定义类型、命名策略、虚拟字段、自动track时间、SQL Builder、Logger
  - 代码生成、复合主键、Constraint、 Prometheus、 Auto Migration、真 跨数据库兼容...
  - 多模式灵活自由扩展
  - Developer Friendly

## 2.2基本用法

先install

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

```go
package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func main() {
  db, err := gorm.Open(sqlite.Open("deme.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&Product{})

  // Create
  db.Create(&Product{Code: "D42", Price: 100})

  // Read
  var product Product
  db.First(&product, 1) // find product with integer primary key
  db.First(&product, "code = ?", "D42") // find product with code D42

  // Update - update product's price to 200
  db.Model(&product).Update("Price", 200)
  // Update - update multiple fields
  db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  db.Delete(&product, 1)
}
```

## 2.2基本用法-CRUD

github：https://github.com/nateshao/gin-demo/blob/main/gin-demo-17-gorm-mysql/main.go

```go
package main

// gorm demo1

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo --> 数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// 连接MySQL数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 创建表 自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	//u1 := UserInfo{1, "七米", "男", "蛙泳"}
	//db.Create(&u1)
	// 查询
	var u UserInfo
	db.First(&u) // 查询表中第一天数据保存到u中
	fmt.Printf("u:%#v\n", u)
	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)
}
```

**新建**

```go
package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gorm demo03

// 1. 定义模型
type User struct {
	ID   int64
	Name sql.NullString `gorm:"default:'小王子'"`
	Age  int64
}

func main() {
	// 连接MySQL数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 2. 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 3. 创建
	u := User{Name: sql.NullString{String: "", Valid: true}, Age: 98} // 在代码层面创建一个User对象
	fmt.Println(db.NewRecord(&u))                                     // 判断主键是否为空 true
	db.Debug().Create(&u)                                             // 在数据库中创建了一条q1mi 18的记录
	fmt.Println(db.NewRecord(&u))                                     // 判断主键是否为空 false
}
```

## 2.3模型定义-惯例约定

**约定优于配置**

- 表名为struct name的snake_ cases复数格式
- 字段名为field name的snake_ case单数格式
- ID/ ld字段为主键，如果为数字，则为自增主键
- CreatedAt字段，创建时，保存当前时间
- UpdatedAt字段，创建、更新时，保存当前时间
- gorm.DeletedAt字段，默认开启soft delete模式

一切皆可配置：https://gorm.io/docs/conventions.html



![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515194116.png)

```go
package main

// gorm demo2
import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// 定义模型
type User struct {
	gorm.Model   // 内嵌gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"` // 零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(120);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

// 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

// 唯一指定表名
func (Animal) TableName() string {
	return "qimi"
}

func main() {
	// 修改默认的表明规则
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "SMS_" + defaultTableName
	}
	// 连接MySQL数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true) // 禁用复数

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

	// 使用User结构体创建名叫 xiaowangzi 的表
	//db.Table("xiaowangzi").CreateTable(&User{})
}
```

## 2.4关联介绍

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515194306.png)

## 2.4关联操作-CRUD

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515194402.png)



## 2.4关联操作-Preload / Joins预加载

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515194440.png)

## 2.4关联操作-级联删除

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515194716.png)



# 03 GORM设计原理

SQL生成，插件扩展，ConnPool，Dialector

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515195020.png)

1. SQL是怎么生成的
2. 插件是怎么工作的
3. ConnPool是什么
4. Dialector

## 3.1SQL是怎么生成的

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515201358.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515201600.png)



![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515201655.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515201824.png)

为什么这样处理呢？

1. 自定义Clause Builder
2. 方便扩展Clause
3. 自由选择Clauses

## 3.1 SQL是怎么生成的-自定义Builder

![image-20220515202111720](https://cdn.jsdelivr.net/gh/nateshao/images/20220515202111.png)

## 3.1SQL是怎么生成的-扩展子句

![image-20220515202429309](https://cdn.jsdelivr.net/gh/nateshao/images/20220515202429.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515202541.png)

## 3.2插件是怎么工作的

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515202711.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515202912.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515203038.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515203101.png)

## 3.2插件是怎么工作的-多租户

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515203657.png)



## 3.2插件是怎么工作的-多数据库、读写分离

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515203853.png)

## 3.3 ConnPool是什么

![image-20220515203958422](https://cdn.jsdelivr.net/gh/nateshao/images/20220515203958.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515204215.png)

## 3.3 ConnPool是什么

![image-20220515204441871](https://cdn.jsdelivr.net/gh/nateshao/images/20220515204442.png)

![image-20220515204617325](https://cdn.jsdelivr.net/gh/nateshao/images/20220515204617.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515204654.png)



## 3最开始的问题

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515204847.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515204944.png)

## 3.4 Dialector是什么

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515205043.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515205226.png)



# 04 GORM最佳实践
1. 数据序列化与SQL表达式

2. 批量数据操作

3. 代码复用、分库分表、Sharding
4. 混沌工程/压测
5. Logger/ Trace
6. Migrator
7. Gen代码生成/ Raw SQL
8. 安全

## 4.1数据序列化与SQL表达式- SQL表达式更新创建

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515225425.png)

## 4.1数据序列化与SQL表达式- SQL表达式查询

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515225545.png)

## 4.1数据序列化与SQL表达式-数据序列化

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515225645.png)

## 4.2批量数据操作-批量创建/查询

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515225906.png)

## 4.2批量数据操作-批量更新

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515230007.png)

## 4.2批量数据操作-批量数据加速操作

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515230131.png)

## 4.3代码复用、分库分表、Sharding -代码复用

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515230310.png)



![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515230321.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515230400.png)

## 4.4混沌工程/压测

![image-20220515230435595](https://cdn.jsdelivr.net/gh/nateshao/images/20220515230435.png)

![image-20220515230655898](https://cdn.jsdelivr.net/gh/nateshao/images/20220515230656.png)

## 4.5 Logger / Trace

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515230728.png)



## 4.6 Migrator -数据库迁移管理

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515230759.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515230851.png)

## 4.7 Gen代码生成, Raw SQL - Raw SQL

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515230922.png)



## 4.8安全问题

https://gorm.io/docs/security.html

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515231234.png)

### 总结

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515231336.png)





































































