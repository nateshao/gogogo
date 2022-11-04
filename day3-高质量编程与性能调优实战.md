---

title: day3-高质量编程与性能调优实战
date: 2022-05-11 09:48:48
tags: 
- Go学习路线
- 字节跳动青训营
---



[TOC]

>关于性能分析这一块，我在之前也写过两篇文章：
>[java性能分析与常用工具](https://mp.weixin.qq.com/s?__biz=MzIyNjE0MDI1NQ==&mid=2247484291&idx=1&sn=a8b5762f1abe5ffed7c28b665a15137f&chksm=e8744ab3df03c3a558ce16149f7a382b7fb0d8eadf0c43ab479d11d201483d407d9e9de8b0a3&token=1303895085&lang=zh_CN#rd)和[Java性能分析之火焰图](https://mp.weixin.qq.com/s?__biz=MzIyNjE0MDI1NQ==&mid=2247484467&idx=1&sn=4a7b01532fc3f38fb0699ce51bcd572f&chksm=e8744d03df03c415f8db852d671d919ac5b27bfa7b1c84c01a16d1a2a264d8b2db0b61673352&token=1303895085&lang=zh_CN#rd)

## 「高质量编程与性能调优实战」第三届字节跳动青训营 - 后端专场

这是我参与「第三届青训营 -后端场」笔记创作活动的的第3篇笔记

同时这也是课表的第三天课程

<img src="https://cdn.jsdelivr.net/gh/nateshao/images/20220511095046.png" style="zoom:50%;" />

PC端阅读效果更佳，点击文末：**阅读原文**即可。源码已经上传到这里：

> 源码：https://github.com/nateshao/gogogo/tree/master/day03-05-11/go-pprof-practice

**这篇文章可以收获什么**

1. 如何编写更简洁清晰的代码
2. 常用Go语言程序优化手段
3. 熟悉Go程序性能分析工具
4. 了解工程中性能优化的原则和流程

今天的内容主要分成两大部分

1. **高质量编程**
   - 编程能够完成功能是基本要求，那么什么是高质量代码，有哪些实践规范，以及常见的性能优化建议有哪些？

2. **性能调优实战**
   - 平时解决算法问题也在追求效率越来越高的算法，在工作中对程序也需要进行不断的优化，这种场景不像算法题那样有明确的流程，如何分析性能瓶颈，使用什么工具，实际服务的优化流程是什么样的会具体说明



# 01.高质量编程

## 1.1简介

**什么是高质量：编写的代码能够达到正确可靠、简洁清晰的目标可称之为高质量代码**

1. 各种边界条件是否考虑完备
2. 异常情况处理，稳定性保证
3. 易读易维护

**编程原则**

实际应用场景千变万化，各种语言的特性和语法各不相同。但是高质量编程遵循的原则是相通的

- **简单性**
  - 消除“多余的复杂性”，以简单清晰的逻辑编写代码
  - 不理解的代码无法修复改进
- **可读性**
  - 代码是写给人看的，而不是机器
  - 编写可维护代码的第一 步是确保代码可读网
- **生产力**
  - > 团队整体工作效率非常重要          -- Go语言开发者Dave Cheney

## 1.2编码规范

**如何编写高质量的Go代码**

1. 代码格式
2. 注释
3. 命名规范
4. 控制流程
5. 错误和异常处理

## 1.2.1编码规范-代码格式

**推荐使用gofmt自动格式化代码**

`gofmt`自动格式化代码，保证所有的`Go`代码与官方推荐格式保持一致。而且可以很方便的进行配置，像`Goland`内置 了相关功能，直接开启即可在保存文件的时候自动格式化。

**gofmt**
`Go`语言官方提供的工具，能自动格式化`Go`语言代码为官方统一风格。常见IDE都支持方便的配置

<img src="https://cdn.jsdelivr.net/gh/nateshao/images/20220511115213.png" style="zoom:50%;" />

<img src="https://cdn.jsdelivr.net/gh/nateshao/images/20220511115311.png" style="zoom:50%;" />

**goimports ** 也是`Go`语言官方提供的工具

`goimports`会对依赖包进行管理，自动增删依赖的包引用，按字母序排序分类，具体可以根据团队实际情况配置使用

之所以将格式化放在第一条， 因为这是后续规范的基础，团队合作`review`其他人的代码时就能体会到这条规范的作用了

## 1.2.2编码规范-注释

#### 简介

**注释应该做的**

1. 注释应该解释代码作用
2. 注释应该解释代码如何做的
3. 注释应该解释代码实现的原因
4. 注释应该解释代码什么情况会出错

```go
Good code has lots of comments, bad code requires lots of comments
好的代码有很多注释，坏代码需要很多注释 
									---Dave Thomas and Andrew Hunt
```

> 多年以后，你再回去看看以前的代码，你是否还看得懂？

**注释应该解释代码作用**

- 适合注释公共符号

  ```go
  // Open opens the named file for reading. If successful, methods on
  // the returned file can be used. for reading; the, 28 6 eC iated file
  // descriptor has mode 0_ RDONLY.
  // If there is an error, it will be of type *PathEr-or.5206
  func Open(name string) (*File, error) {
  	return OpenFile(name, 0_RDONLY,0)
  }
  ```

  原文：https://github.com/golang/go/blob/master/src/os/file.go#L313

  ```go
  // Returns true if the table cannot hold any more entries
  func IsTableFull( ) bool
  ```

**注释应该解释代码如何做的**

- 适合注释实现过程

  ```go
  // Add the Referer header from the most recent
  // request URL to the new one，if it's not https->http:
  if ref := refererForURL( reqs[len(reqs)-1].URL,req.URL); ref !="" {
  	req.Header.Set( "Referer", ref)
  }
  ```

  原文链接：https://github.com/golang/go/blob/master/src/net/http/client.go#L678

  ```go
  // Process every element in the list
  for e := range e Lements {
  	process(e)
  }
  ```

**注释应该解释代码实现的原因**

- 适合解释代码的外部因素

- 提供额外上下文

  ```go
  switch resp.StatusCode {
  // ...
  case 307, 308:
  redirectMethod = reqMethod
  	shouldRedirect = true
  	includeBody = true
  	if ireq.GetBody == nil && ireq.outgoingLength() != 0 {
  		// We had a request body, and 307/308 require
  		// re-sending it, but GetBody is not defined. So just
  		// return this response to the user instead of an
  		// error, like we did in Go 1.7 and earlier.
  		shouldRedirect = false
  	}
  }
  ```

  原文链接：https://github.com/golang/go/blob/master/src/net/http/client.go#L521

**注释应该解释代码什么情况会出错**

- 适合解释代码的限制条件

  ```go
  // parseTimeZone parses a time zone string and returns its length. Time zones
  // are human-generated and unpredictable can't do precise error checking.
  // 0n the other hand, fora correct parsl there must be a time zone at the
  // beginning of the string, so it's almost always true that there's one
  // there. We look at the beginning of the string for a run of upper-case letters.
  // If there are more than 5，it's an error.
  // If there are 4 or 5 and the last is aT，it's a time zone.
  // If there are3，it's a time zone.06
  // Otherwise, other than special cases, it's not a time zone.
  // GMT is special because it can have an hour offset.
  func parseTimeZone(value string) (length int， ok bool)
  ```

  原文链接：https://github.com/golang/go/blob/master/src/time/format.go#L1344

**公共符号始终要注释**

- 包中声明的每个公共的符号：变量、常量、函数以及结构都需要添加注释

- 任何既不明显也不简短的公共功能必须予以注释

- 无论长度或复杂程度如何，对库中的任何函数都必须进行注释

  ```go
  // ReadAll reads from r until an error or EOF and returns the data it read.
  // A successful call returns err == nil, not err == EOF. Because ReadAll is
  // defined to read from src until EOF, it does not treat an EOF from Read
  // as an error to be reported. 
  func ReadAll(r Reader) ([]byte,error)
  ```

  原文链接：https://github.com/golang/go/blob/master/src/io/io.go#L638

**公共符号始终要注释**

- 有一个例外，不需要注释实现接口的方法。具体不要像下面这样做

  ```go
  // Read implements the io.Rladet, interface
  func (r *FileReader) Read(bur []byte) (int error)
  ```

  

**公共符号始终要注释**

- 对于公共符号都有注释说明尽管L imitedReader.Read本身没有注释，但它紧跟LimitedReader结构的声明，明确它的作用

  ```go
  // LimitReader. returns a Reader that reads from r
  // but stops with EOF after n bytes.
  // The underlying imp lementation is a *L im i tedReader.
  func LimitReader(r Reader, n int64) Reader { return &LimitedReader{r, n} }
  // A LimitedReader reads from R but limits the amount of
  // data returned to just N bytes. Each call to Read
  // updates N to reflect the new amount remaining.
  // Read returns EOF when N <= 0 or when the underlying R returns EOF.
  type LimitedReader struct {
  	R Reader // underlying reader
  	N int64 // max bytes remaining
  }
  func (l *LimitedReader) Read(p []byte) (n int, err error) {
  	if l.N<=0{
  		return 0，EOF
  	}
  	if int64(len(p)) > l.N {
  		p = p[0:l.N]
  	n.err = l.R.Read(p)
  	l.N -= int64(n)
  	return
  }
  ```

### 小结

- 代码是最好的注释
- 注释应该提供代码未表达出的上下文信息 

## 1.2.3编码规范-命名规范

**variable**

- 简洁胜于冗长
- 缩略词全大写，但当其位于变量开头且不需要导出时，使用全小写
  - 例如使用`ServeHTTP`而不是`ServeHttp`
  - 使用`XMLHTTPRequest`或者`xmIHTTPRequest`

- 变量距离其被使用的地方越远，则需要携带越多的下 文信息
  - 全局变量在其名字中需要更多的上下文信息，使得在不同地方可以轻易辨认出其含义

**举例：**

```go
// Good
func (C *Client) send(req *Request, deadline time.Time)
// Bad
func (C *Client) send(req *Request, t time.Time)
```

- 将`deadline`替换成t降低了变量名的信息量
- `t`常代指任意时间
- `deadline`指截止时间，有特定的含义

> 所以建议不要把`deadline`换成 `t` ,这样降低了变量名的信息量

## 1.2.3编码规范-命名规范

**function** 

- 函数名不携带包名的上下文信息，因为包名和函数名总是成对出现的
- 函数名尽量简短
- 当名为`foo`的包某个函数返回类型`Foo`时，可以省略类型信息而不导致歧义
- 当名为`foo`的包某个函数返回类型T时(T并不是`Foo`),可以在函数名中加入类型信息

> http包中创建服务的函数如何命名更好?  

```go
func Serve(l net._istener, handler Handler) error
func ServeHTTP(l net.Listener, handler Handler) error
```

推荐使用第一种，原因在于你在调用http包的时候是.Serve

**package**

- 只由小写字母组成。不包含大写字母和下划线等字符

- 简短并包含一定的上下文信息。例如`schema`、`task `等

- 不要与标准库同名。例如不要使用`sync`或者`strings`

以下规则尽量满足，以标准库包名为例

- 不使用常用变量名作为包名。例如使用`bufio`而不是`buf`
- 使用单数而不是复数。例如使用`encoding`而不是`encodings`
- 谨慎地使用缩写。例如使用fmt在不破坏上下文的情况下比`format`更加简短

#### 小结：

- 核心目标是降低阅读理解代码的成本
- 重点考虑上下文信息，设计简洁清晰的名称

>Good naming is like a good joke. If you have to explain it, it s not funny
>好的命名就像一个好笑话。 如果你必须解释它，那就不好笑了
>																						--Dave Cheney



## 1.2.4编码规范-控制流程

**避免嵌套，保持正常流程清晰**

如果两个分支中都包含return语句，则可以去除冗余的else。 

```go
// Bad
if foo {
	return X
} else {
	return nil
}
// Good
if foo {
	return X
}
return nil
```

**尽量保持正常代码路径为最小缩进**
优先处理错误情况特殊情况，尽早返回或继续循环来减少嵌套

```go
// Bad
func OneFunc( ) error {
	err := doSomething( )
	if err == nil{
		err := doAnotherThing( )
		if err == nil{
		return nil // normal case
	}
		return err
	}
	return err
}
```

- 最常见的正常流程的路径被嵌套在两个if条件内
- 成功的退出条件是return nil,必须仔细匹配大括号来发现
- 函数最后一行返回一个错误，需要追溯到匹配的左括号，才能了解何时会触发错误
- 如果后续正常流程需要增加一步操作，调用新的函数，则又会增加一层嵌套

<center>调整后</center>

```go
// Good
func OneFunc( ) error {
	if err := doSomething(); err != nil {
		return err
	}
	if err. := doAnotherThing(); err != nil {
		return err
	}
	returrpnil // normal case
}
```

---

```go
func (b *Reader) UnreadByte( ) error {
	if b.lastByte < 0 II b.r == 0 && b.W> 0{
		return ErrInvalidUnreadByte
	}千
	// b.r> 0//b.w== 0
	ifb.r>0{
		b.r--
		206
	} else {
		// b.r==20 && b.w == 0
		b.w=1
	}
	b.buf[b.r] = byte(b. lastByte)
	b.lastByte = -1
	b. lastRuneSize = -1 ( 6
	return nil
}
```



#### 小结

- 线性原理，处理逻辑尽量走直线，避免复杂的嵌套分支
- 正常流程代码沿着屏幕向下移动
- 提升代码可维护性和可读性
- 故障问题大多出现在复杂的条件语句和循环语句中

## 1.2.5编码规范-错误和异常处理

**简单错误**

1. 简单的错误指的是仅出现一次的错误， 且在其他地方不需要捕获该错误
2. 优先使用`errors.New`来创建匿名变量来直接表示简单错误
3. 如果有格式化的需求，使用`fmt.Errorf`

```go
func defaultCheckRedirect(req *Request, via []*Request) error {
	if len(via) >= 10 {
		return errors . New( "stopped after 10 redirects")
	}
	return nil
}
```

原文链接：https://github.com/golang/go/blob/master/src/net/http/client.go



**错误的Wrap和Unwrap**

- 错误的Wrap实际上是提供了一个`error`嵌套另一个`error`的能力，从而生成一个`error`的跟踪链

- 在`fmt.Errorf`中使用: %w关键字来将一个错误关联至错误链中

  ```go
  list,_,err := c.GetBytes( cache. Subkey(a.actionID, "srcfiles"))
  if err != nil{
  	return fmt. Errorf("reading srcfiles list: %W",err )
  }
  ```

  https://github.com/golang/go/blob/master/src/cmd/go/internal/work/exec.go#L.983

>Go1.13在errors中新增了三个新API和一个新的format关键字，分别是errors.Is errors.As,errors.Unwrap以及fmt.Errorf的%W。
>
>如果项目运行在小于Go1.13的版本中，导入golang.org/x/ xerrors来使用



**错误判定**

- 判定一个错误是否为特定错误，使用`errors.Is`

- 不同于使用==,使用该方法可以判定错误链.上的所有错误是否含有特定的错误

  ```go
  data, err = lockedfile. Read( targl2209
  	if errors.Is(err, fs.ErrNotExist) {
  		// Treat non-existent as empty，to bootstrap the "latest" file
  		// the first time we connect to a given database.
  		return []byte{}, nil
  	}
  return data, err
  ```

  https://github.com/golang/go/blob/master/src/cmd/go/internal/modfetch/sumdb.go#L208

**错误判定**

- 在错误链上获取特定种类的错误，使用`errors.As`

  ```go
  if _,err := os.0pen( "non-existing"); err != nil {
  	var pathError *fs. PathEr r
  		if errors.As(err, &pathEroor) {
  		fmt. Println("Failed at path:", pathError .Path)
  	} else {
  		fmt.Println(err)
  	}
  }
  ```

  

**panic** 

1. 不建议在业务代码中使用`panic`

2. 调用函数不包含`recover`会造成程序崩溃

3. 若问题可以被屏蔽或解决，建议使用error代替`panic`

4. 当程序启动阶段发生不可逆转的错误时，可以在init或main函数中使用`panic`

   ```go
   func main() {
   	ctx, cancel := context.WithCancel(context.Background( ))
   	client,err := sarama.NewConsumerGroup(strings.Split(brokers,"), group, config)
   	if err!=nil{
   		log.Panicf("Error creating consumer group client: %V",err)
   	}
   }
   // Panicf is equivalent to Printf() followed by a call to panic().
   func Panicf(format string, V ... interface{}) {
   	S := fmt.Sprintf(format, V... )
   	std.output(2,s)
   	panic(s)
   }
   ```

   

**recover**

1. `recover`只能在被defer的函数中使用

2. 嵌套无法生效

3. 只在当前`goroutine`生效

4. `defer`的语句是后进先出

   ```go
   func (s *ss) Token(skipSpace bool, f func( rune) bool)
   	(tok []byte, err error) {
   		defer func () {
   			if e := recover(); e != nil {
   				if se, ok := e.(scanError); ok {
   					err = se.err
   				} else {
   				panic(e)
   			}
   		}
   	}()
   	// ...
   }
   ```

5. 如果需要更多的上下文信息，可以recover 后在log中记录当前的调用栈

   ```go
   func (t *treeFS) Open(name string) (f fs.File, err error) {
   	defer func() {
   		if e := recover(); e != nitt {
   			f = nil
   			err = fmt.Errorf("gitfs panic: %V\n%S", e, debug.Stack())
   		}
   		// ...
   	}
   }
   ```


### 小结

1. `error`尽可能提供简明的.上下文信息链，方便定位问题
2. `panic`用于真正异常的情况
3. `recover`生效范围，在当前`goroutine`的被`defer`的函数中生效







## 1.2编码规范

**哪种命名方式更好?**

```go
package time
// A functione returns the current local t ime.
// which one is better?
func Now() Time
// or
func NowTime() Time
```

看看调用的时候，`time.Now()`更简洁

```go
t := time.Now()
t := time.NowTime()
```

---

```go
package time
// A function parses a duration string.
// such as "300ms", "-1.5h” or "2h45m'
func Parse(s string) (Duration, error)
// or
func ParseDuration(s string) (Duration, error)
```

实际调用的时候

```go
duration := time.Parse(s)
duration := time.ParseDuration(s)
```

---

**程序的输出是什么?**

```go
func main() {
	
	if true {
		defer fmt.Printf("1")
	} else {
		defer fmt.Printf("2")
	}
	defer fmt.Printf("3")
}
// 输出31
```

`defer`语句会在函数返回前调用
多个`defer`语句是后进先出

## 1.3性能优化建议

- 性能优化的前提是满足正确可靠、简洁清晰等质量因素
- 性能优化是综合评估，有时候时间效率和空间效率可能对立
- 针对`Go`语言特性，介绍`Go`相关的性能优化建议

高质量的代码能够完成功能，但是在大规模程序部署的场景，仅仅支持正常功能还不够，我们还要尽可能的提升性能，节省资源成本。

接下来就主要介绍**性能相关的建议**

- 高性能代码为了效率会用到许多技巧，没有相关背景的人难以理解，不过有些基础性能问题是和语言本身相关的，接下来主要介绍这类内容，对应的调整对可读性和可维护性影响不大
- 在满足正确性、可靠性、健壮性、可读性等质量因素的前提下，设法提高程序的性能
- 有时候时间效率和空间效率可能对立，此时应当分析那个更重要，作出适当的折衷。例如多花费一些内存来提高性能。针对Go语言编程，介绍`Go`相关的性能优化建议

## 1.3.1 性能优化建议-Benchmark

**如何使用？**

- 性能表现需要实际数据衡量
- `Go`语言提供了支持基准性能测试的`benchmark `工具

```go
// from fib.go
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// from fib_test.go
func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
```

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511230441.png)

我们如何评估性能， 性能表现要用数据说话，实际情况和想象中的并不一定一致， 要用数据来验证我们写的代码是否真的有性能提升

`Go`自带了性能评估工具

以计算斐波拉契数列的函数为例，分两个文件，`fib.go`编写 函数代码,`fib._test.go`编写`benchmark`的逻辑，通过命令运行`benchmark`可以得到测试结果`benchmem`表示也统计内存信息

## 1.3.2 性能优化建议-Slice

**slice预分配内存**

- 尽可能在使用`make()`初始化切片时提供容量信息

```go
func NoPreAlloc(size int) {
	data := make([]int, 0)
	for k := 0; k < size; k++ {
		data = append(data, k)
	}
}
```

```go
func PreAlloc(size int) {
	data := make([]int, 0, size)
	for k := 0; k < size; k++ {
		data = append(data, k)
	}
}
```

`slice`是`go`中最常用的结构，也很方便，那么在使用过程中有哪些点需要注意几点：

1. 第一条建议就是**预分配**，尽可能在使用make()初始化切片时提供容量信息，特别是在追加切片时对比看下两种情况的性能表现，左边是没有提供初始化容量信息，右边是设置了容量大小。结果中可以看出执行时间相差很多，预分配只有一次内存分配

图来源：https://ueokande.github.io/go-slice-tricks

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511183140.png)

我们看看`slice`的结构：**切片本质是一个数组片段的描述，包括数组指针，片段的长度，片段的容量(不改变内存分配情况下的最大长度)**

**，切片操作并不复制切片指向的元素**。创建一个新的切片 会复用原来切片的底层数组，以切片的`append`为例，`append`时有两种场景 :

1. 当`append`之后的长度小于等于`cap`，将会直接利用原底层数组剩余的空间。
2. 当`append`后的长度大于`cap`时，则会分配-块更大的区域来容纳新的底层数组。

因此，为了避免内存发生拷贝，如果能够知道最终的切片的大小，预先设置`cap`的值能够避免额外的内存分配，获得更好的性能

---

- 切片本质是一个数组片段的描述
  - 包括数组指针
  - 片段的长度
  - 片段的容量(不改变内存分配情况下的最大长度)

- 切片操作并不复制切片指向的元素
- 创建一个新的切片会复用原来切片的底层数组

```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```

------

**另一个陷阱：大内存未释放**
在已有切片基础上创建切片，不会创建新的底层数组

场景

- 原切片较大，代码在原切片基础上新建小切片
- 原底层数组在内存中有引用，得不到释放

**可使用copy替代re-slice**

```go
func GetLastBySlice(origin []int) []int {
	return origin[len(origin)-2:]
}
func GetLastByCopy(origin []int) []int {
	result := make([]int, 2)
	copy(result, origin[len(origin)-2:])
	return result
}
```

```go
func testGetLast(t *testing.T, f func([]int) []int) {
	result := make([][]int, 0)
	for k := 0; k < 100; k++ {
		origin := generateWithCap(128 * 1024) // 1M
		result = append(result, f(origin))
	}
	printMem(t)
	_ = result
}
```

Go 语言高性能编程：https://geektutu.com/post/hpg-slice.html

## 1.3.3性能优化建议-Map

**map预分配内存**

除了`slice`，`map`也是编程中常用的结构，是不是也有预分配的性能优化点呢？可以对比验证下

```go
func NoPreAlloc(size int) {
	data := make(map[int]int)
	for i := 0; i < size; i++ {
		data[i] = 1
	}
}
```

---

```go
func PreAlloc(size int) {
	data := make(map[int]int, size)
	for i := 0; i < size; i++ {
		data[i] = 1
	}
}
```

**分析**

- 不断向`map`中添加元素的操作会触发`map`的扩容
- 提前分配好空间可以减少内存拷贝和`Rehash`的消耗
- 建议根据实际需求提前预估好需要的空间

## 1.3.4性能优化建议-字符串处理

**使用strings.Builder**

编程过程中除了`slice`和`map`，平时很多得功能都和子字符用处理相关的，字符串处理也是高频操作。

- 常见的字符串拼接方式

  ```go
  func Plus(n int, str string) string {
  	s := ""
  	for i := 0; i < n; i++ {
  		s += str
  	}
  	return s
  }
  ```

  ---

  ```go
  func StrBuilder(n int, str string) string {
  	var builder strings.Builder
  	for i := 0; i < n; i++ {
  		builder.WriteString(str)
  	}
  	return builder.String()
  }
  ```

  ---

  ```go
  func ByteBuffer(n int, str string) string {
  	buf := new(bytes.Buffer)
  	for i := 0; i < n; i++ {
  		buf.WriteString(str)
  	}
  	return buf.String()
  }
  ```

**使用+拼接性能最差，strings.Builder, bytes.Buffer 相近，strings.Buffer 更快**
分析

- 字符串在`Go`语言中是不可变类型，占用内存大小是固定的
- 使用+每次都会重新分配内存
- `strings.Builder`, ` bytes.Buffer`底层都是[]byte 数组
- 内存扩容策略，不需要每次拼接重新分配内存

当使用+拼接2个字符串时，生成一个新的字符串，那么就需要开辟一段新的空间， 新空间的大小是原来两个字符串的大小之和。拼接第三个字符串时，再开辟一段新空间， 新空间大小是三个字符串大小之和，以此类推。

---

**为什么stringbuilder会比bytebuffer更快一些**，可以看看实际的代码

注意注释里也提到如果想用更高效的字符串构造方法，可以使用string builder

联系上面的内容，有没有办法再次提升字符串拼接的效率？关键字是什么？预分配

```go
// To build strings more efficiently, see the strings . Builder type.
func (b *Buffer) String() string {
	if b == nil {
		// Special case, useful in debugging.
		return "<nil>"
	}
	return string(b.buf[b.off:])
}
```

`bytes.Buffer`转化为字符串时重新申请了一块空间
`strings.Builder`直接将底层的[]byte转换成了字符串类型返回

```go
// String returns the accumulated string.
func (b *Builder) String() string {
	return *(*string)(unsafe.Pointer(&b.buf))
}
```

---

字符串拼接和`slice`一样，同样支持预分配，在预知字符串长度的情况下，我们可以进一步提升拼接性能
注意这里能确认`stringbuiler`只有一次内存分配，`bytebuffer`有两次

```go
func PreStrBuilder(n int, str string) string {
	var builder strings.Builder
	builder.Grow(n * len(str))
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}
```

```go
func PreByteBuffer(n int, str string) string {
	buf := new(bytes.Buffer)
	buf.Grow(n * len(str))
	for i := 0; i < n; i++ {
		buf.WriteString(str)
		return buf.String()
	}
}
```

## 1.3.5 性能优化建议-空结构体

性能优化有时是时间和空间的平衡，之前提到的都是提高时间效率的点，而空结构体是节省内存空间的一个手段。

**使用空结构体节省内存**

- 空结构体`structf`实例不占据任何的内存空间
- 可作为各种场景下的占位符使用
  - 节省资源
  - 空结构体本身具备很强的语义，即这里不需要任何值，仅作为占位符

```go
func EmptyStructMap(n int) {
	m := make(map[int]struct{})
	for i := 0; i < n; i++ {
		m[i] = struct{}{}
	}
}
func BoolMap(n int) {
	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		m[i] = false
	}
}
```

空结构体占用内存更少些，在元素更多的情况下会更明显。实际应用场景有哪些，容易想到的是set实现

- 实现Set,可以考虑用map来代替
- 对于这个场景，只需要用到map的键，而不需要值
- 即使是将map的值设置为bool类型，也会多占据1个字节空间

一个开源实现： https://github.com/deckarep/golang-set/blob/main/threadunsafe.go

## 1.3.6性能优化建议-atomic包

在工作中迟早会遇到多线程编程的场景，比如实现一个多线程共用的计数器，如何保证计数准确，线程安全，有不同的方式

**如何使用atomic包**

```go
type atomicCounter struct {
	i int32
}

func AtomicAddOne(c *atomicCounter) {
	atomic.AddInt32(&c.i, 1)
}
```

```go
type mutexCounter struct {
	i int32
	m sync.Mutex
}

func MutexAdd0ne(c *mutexCounter) {
	c.m.Lock()
	c.i++
	c.m.Unlock()
}
```

- 锁的实现是通过操作系统来实现，属于系统调用
- atomic操作是通过硬件实现，效率比锁高
- sync.Mutex应该用来保护一段逻辑，不仅仅用于保护一个变量
- 对于非数值操作，可以使用atomic.Value，能承载一个interfacef

## 1.3性能优化建议

### 小结

- 避免常见的性能陷阱可以保证大部分程序的性能
- 普通应用代码，不要一味地追求程序的性能
- 越高级的性能优化手段越容易出现问题
- 在满足正确可靠、简洁清晰的质量要求的前提下提高程序性能

# 02.性能调优实战

上面我们讲了高质量编程的原则和一些实践规范，同时给出了一些性能优化建议，那么在实际工作中，如何要针对某个应用程序进行性能调优，应该如何做呢?

## 2.1简介

**性能调优原则**

- 要依靠数据不是猜测
- 要定位最大瓶颈而不是细枝末节
- 不要过早优化
- 不要过度优化

## 2.2性能分析工具pprof

既然性能调优前提是对应用程序性能表现有实际的数据指标，那么有什么工具能够获得这种数据呢?

对于go程序，有一个很方便的工具就是pprof。接下来我们从三个方面来熟悉下pprof工具

1. pprof功能简介
2. pprof排查实战
3. pprof的采样过程和原理

## 2.2.1性能分析工具pprof-功能简介

具体pprof有哪些内容?可以看下图片

- 分析部分一有两种方式，网页和可视化终端
- 具体的工具-可以在runtime/pprof中找到源码，同时Golang的http标准库中也对pprof做了 一些封装，能让你在http服务中直接使用它
- 采样部分-它可以采样程序运行时的CPU、堆内存、goroutine、 锁竞争、 阻塞调用和系统线程的使用数据
- 展示-用户可以通过列表、调用图、火焰图、源码、反汇编等视图去展示采集到的性能指标。方便分析
  因为pprof的功能比较多，接下来通过一个实践项目来熟悉pprof工具的使用

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511194303.png)



## 2.2.2性能分析工具pprof-排查实战

**搭建pprof实践项目**

我们的目标是熟悉pprof工具，能够排查性能问题，那么首先我们需要构造一个有问题的程序， 看看如何用pprof来定位性能问题点
这里有个开源项目，已经制造了一些问题代码，需要我们进行排查。大家用pprof示例命令实验下能否正常打开pprof页面，是否缺少graphviz的组件

我们来看看「炸弹I程序是怎么做的。图中代码是main.go中初始化htp服务和pprof接口的代码，无关逻辑有所省略。可以看到，引入http pprof这个包，它会将pprof的入口注册到/debug/pprof这个路径下， 我们可通过浏览器打开这个路径，来查着些基本的性能统计。
运行「炸弹」， 并等待它运行稳定

```go
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	"github.com/wolfogre/go-pprof-practice/animal"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	for {
		for _, v := range animal.AllAnimals {
			v.Live()
		}
		time.Sleep(time.Second)
	}
}
```

---

**浏览器查看指标**

在浏览器中打开http://localhost:6060/debug/pprof,可以看到这样的页面，这就是我们刚刚引入的net/http/pprof注入的入口了。
页面上展示了可用的程序运行采样数据，下面也有简单说明， 分别是:

- allocs：内存分配情况
- blocks：阻塞操作情况
- cmdline：程序启动命令及
- goroutine：当前所有goroutine的堆栈信息
- heap：堆上内存使用情况(同alloc)
- mutex：锁竞争操作情况
- profile：CPU占用情况
- threadcreate：当前所有创建的系统线程的堆栈信息
- trace：程序运行跟踪信息

cmdline显示运行进程的命令，threadcreate比较复杂，不透明，trace需要另外的工具解析，暂不涉及。

炸弹在CPU，堆内存，goroutine, 锁竞争和阻塞操作上埋了炸弹，可以使用pprofI具进行分析

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511195018.png)

**allocs表示分配了12次内存，block表示有3次阻塞，goroutine表示有55个协程正在运行，heap为堆上内存使用，mutex表示有1个锁竞争，threadcreate表示有6个线程创建。**

浏览器访问：http://localhost:6060/debug/pprof/allocs?debug=1

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511194902.png)



---





**CPU**

**命令: topN**     查看占用资源最多的函数

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220512113456.png)

我们先从CPU问题排查开始，不同的操作系统具可能不同，我们首先使用自己熟悉的工具看看程序进程的资源占用，CPU占用了13.2%，显然这里是有问题的

pprof的采样结果是将一段时间内的信息汇 总输出到文件中，所以首先需要拿到这个profile文件。你可以直接使用暴露的接口链接下载文件后使用，也可以直接用pprofI具连接这个接口下载需要的数据。

这里我们使用go tool pprof +采样链接来启动采样。链接中就是刚刚「炸弹」 程序暴露出来的接口，链接结尾的profile代表采样的对象是CPU使用。 如果你在浏览器里直接打开这个链接，会启动一个60秒的采样，并在结束后下载文件。这里我们加上seconds=10的参数，让它采样+秒。稍等片刻，我们需要的采样数据已经记录和下载完成，并展示出pprof终端

go tool pprof "http://localhost:6060/debug/pprof/profile?seconds=10"

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511195948.png)

---

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511200117.png)

很明显`github.com/wolfogre/go-pprof-practice/animal/felidae/tiger.(*Tiger).Eat`的调用占用了绝大部分CPU。



输入top， 查看CPU占用最高的函数。这五列从左到右分别是：

| flat  | 当前函数本身的执行耗时               |
| ----- | ------------------------------------ |
| flat% | flat占CPU总时问的比例                |
| sum % | 上面每一行的flat%总和                |
| cum   | 指当前函数本身加上其调用函数的总耗时 |
| cum % | cum占CPU总时间的比例                 |

表格前面描述了采样的总体信息。默认会展示资源占用最高的10个函数，如果只需要查看最高的N个函数，可以输入topN, 例如查看最高的3个调用，输入top3
可以看到表格的第一行里， Tiger.Eat函数本身 占用3.56秒的CPU时间，占总时间的95.44%， 显然问题就是这里引起的

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511200626.png)

---

什么情况下Flat ==Cum?
什么情况下Flat ==O?

---

Flat == Cum，函数中没有调用其他函数
Flat == 0，函数中只有其他函数的调用

---

**命令：list       根据指定的正则表达式查找代码行**

接着，输入`list Eat`查找这个函数，看看具体是哪里出了问题

List命令会根据后面给定的正则表达式查找代码，并按行展示出每一行的占用。可以看到，第24行有 一个100亿次的空循环，占用了5.07秒的CPU时间，问题就在这儿了，定位成功。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220512122231.png)

代码注释一下

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220512122043.png)



---

**命令：web       调用关系可视化**

> 输入`web命令`,这个命令前提是需要下载`graphviz`  https://graphviz.org/download/

除了这两种视图之外，我们还可以输入web命令，生成一张调用关系图，默认会使用浏览器打开。图中除了每个节点的资源占用以外，还会将他们的调用关系穿起来。图中最明显的就是方框最红最大，线条最粗的 *Tiger.Eat函数， 是不是比top视图更直观些呢？到这里，CPU的炸弹已经定位完成，我们输入q退出终端。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220512115746.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220512121732.png)

为了方便后续的展示，我们先将这个「炸弹|拆除，注意这一部分我们的主 要目的是展示如何定位问题，所以不会花太多时间在看代码上，我们将问题代码直接注释掉，重新打开「炸弹」程序。

---

**Heap-堆内存**

注释CPU问题代码，重新运行后。打开活动监视器，可以发现进程的CPU已经降下来了。然而内存使用还是很高接我们来排查内存问题。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220512122827.png)

---

在刚刚排查CPU的过程中，我们使用的是pprof终端，这里我们介绍另一种展示方式。通过-http=:8080参数， 可以开启pprof自带的Web UI,性能指标会以网页的形式呈现。
再次启动pprof工具，注意这时的链接结尾是heap。等待采样完成后，浏览器会被自动打开，展示出熟悉的web视图，同时展示的资源使用从「CPU时间」变为了「内存占用」
可以明显看到，这里出问题的是*Mouse.Steal0函数， 它占用了1GB内存。在页面顶端的View菜单中，我们可以切换不同的视图。

go tool pprof -http=:8080 "http://localhost:6060/debug/pprof/heap“

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511204447.png)

命令输入之后，会自动跳转到http://localhost:8080/ui/ ，就可以看到如下画面了

![image-20220512123119455](https://cdn.jsdelivr.net/gh/nateshao/images/20220512123119.png)

---

我们再切换到Source视图，可以看到页面上展示出了刚刚看到的四个调用和具体的源码视图。如果觉得内容太多，也可以在顶部的搜索框中输入Steal来使用正则表达式过滤需要的代码。

根据源码我们会发现，在*Mouse.Steal()这 函数会向固定的Buffer中不断追加1MB内存，直到Buffer达到1GB大小为止， 和我们在Graph视图中发现的情况一致。
我们将这里的问题代码注释掉，至此，「炸弹 」已被拔除了两个。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511204703.png)

---

重新运行「炸弹」程序，发现内存占用已经降到了23.6MB,刚才的解决方案起效果了
在采样中，也没有出现异常节点了(实际上没有任何节点了)。不过，内存的问题真的被全部排除了吗?
大家在使用工具的过程中，有没有注意过右上角有个unknown jinuse. space

我们打开sample菜单，会发现堆内存实际上提供了四种指标：

在堆内存采样中，默认展示的是inuse_ space视图，只展示当前持有的内存， 但如果有的内存已经释放，这时inuse采样就不会展示了 。我们切换到alloc_ space指标。后续分析下alloc的内存问题

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511204821.png)

---

马上就定位到了*Dog.Run0中还藏着一个[炸弹」，它会 每次申请16MB大小的内存，并且已经累计申请了超过3.5GB内存。在Top视图中还能看到这个函数被内联了。

看看定位到的函数，果然，这个函数在反复申请16MB的内存，但因为是无意义的申请，分配结束之后会马上被GC,所以在inuse采样中并不会体现出来。

我们将这一行问题代码注释掉， 继续接下来的排查。至此，内存部分的「炸弹」已经被全部拆除。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511204903.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220512123352.png)

```go
github.com/wolfogre/go-pprof-practice/animal/muridae/mouse.(*Mouse).Steal
D:\Golang\GoWordPlace\pkg\mod\github.com\wolfogre\go-pprof-practice@v0.0.0-20190402114113-8ce266a210ee\animal\muridae\mouse\mouse.go

  Total:         1GB        1GB (flat, cum)   100%
     45            .          .            
     46            .          .           func (m *Mouse) Steal() { 
     47            .          .           	log.Println(m.Name(), "steal") 
     48            .          .           	max := constant.Gi 
     49            .          .           	for len(m.buffer) * constant.Mi < max { 
     50          1GB        1GB           		m.buffer = append(m.buffer, [constant.Mi]byte{}) 
     51            .          .           	} 
     52            .          .           } 
runtime.main
D:\Golang\Go\src\runtime\proc.go

  Total:           0        1GB (flat, cum)   100%
    198            .          .           		// A program compiled with -buildmode=c-archive or c-shared 
    199            .          .           		// has a main, but it is not executed. 
    200            .          .           		return 
    201            .          .           	} 
    202            .          .           	fn := main_main // make an indirect call, as the linker doesn't know the address of the main package when laying down the runtime 
    203            .        1GB           	fn() 
    204            .          .           	if raceenabled { 
    205            .          .           		racefini() 
    206            .          .           	} 
    207            .          .            
    208            .          .           	// Make racy client program work: if panicking on 
github.com/wolfogre/go-pprof-practice/animal/muridae/mouse.(*Mouse).Live
D:\Golang\GoWordPlace\pkg\mod\github.com\wolfogre\go-pprof-practice@v0.0.0-20190402114113-8ce266a210ee\animal\muridae\mouse\mouse.go

  Total:           0        1GB (flat, cum)   100%
     18            .          .           	m.Eat() 
     19            .          .           	m.Drink() 
     20            .          .           	m.Shit() 
     21            .          .           	m.Pee() 
     22            .          .           	m.Hole() 
     23            .        1GB           	m.Steal() 
     24            .          .           } 
     25            .          .            
     26            .          .           func (m *Mouse) Eat() { 
     27            .          .           	log.Println(m.Name(), "eat") 
     28            .          .           } 
main.main
F:\GitHub--Gitee\zijie\go-pprof-practice\main.go

  Total:           0        1GB (flat, cum)   100%
     26            .          .           		os.Exit(0) 
     27            .          .           	}() 
     28            .          .            
     29            .          .           	for { 
     30            .          .           		for _, v := range animal.AllAnimals { 
     31            .        1GB           			v.Live() 
     32            .          .           		} 
     33            .          .           		time.Sleep(time.Second) 
     34            .          .           	} 
     35            .          .            
     36            .          .           } 
```



---

**goroutine-协程**

`goroutine`泄露也会导致内存泄露

`Golang`是门自带垃圾回收的语言，一般情况 下内存泄露是没那么容易发生的
但是有种例外: `goroutine`是很容易泄露的，进而会导致内存泄露。所以接下来，我们来看看`goroutine`的使用情况。
打开http://ocalhost:6060/debug/pprof/   发现「炸弹」程序已经有54条`goroutine`在运行了，这个量级并不是很大，但对于一个简单的小程序来说还是很不正常的。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511205102.png)



---

大家对链接含义应该很熟悉了，把链接末尾换成`goroutine`, 可以看到，`pprof`生成了一张非常长的调用关系图，尽管最下方的问题用红色标了出来，不过这么多节点还是比较难以阅读的。这里我们介绍另一种更加直观的展示方式一火焰图。

go tool pprof -http=:8080   "http://localhost:6060/debug/pprof/goroutine”

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511205550.png)

----

打开`View`菜单，切换到Flame Graph视图。可以看到，刚才的节点被堆叠了起来
图中，自顶向下展示了各个调用，表示各个函数调用之间的层级关系
每一行中，条形越长代表消耗的资源占比越多。显然，那些[又平又长」的节点是占用资源多的节点
可以看到，*Wolf.Drink()这 个调用创建了超过90%的`goroutine`,问题可能在这里

火焰图是非常常用的性能分析工具，在程序逻辑复杂的情况下很有用，可以重点熟悉



![](https://cdn.jsdelivr.net/gh/nateshao/images/20220512123516.png)

- 由上到下表示调用顺序
- 每一块代表一个函数，越长代表占用`CPU`的时间更长
- 火焰图是动态的，支持点击块进行分析

---

这里为了模拟泄漏场景，只等待了30秒就退出了试想下，如果发起的`goroutine`没有退出，同时不断有新的`goroutine`被启动， 对应的内存占用持续增长，CPU调度压力也不断增大， 最终进程会被系统il掉

- 支持搜索，在`Source`视图下搜索`wolf`

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511205832.png)

---

**mutex-锁**

修改链接后缀，改成mutex，然后打开网页观察，发现存在1个锁操作

同样地，在`Graph`视图中定位到出问题的函数在*Wolf.HowI()。然后在`Source`视图中定位到具体哪一行发生了锁竞争

在这个函数中，`goroutine`足足等待 了1秒才解锁，在这里阻塞住了，显然不是什么业务需求，注释掉。

![image-20220511210024362](https://cdn.jsdelivr.net/gh/nateshao/images/20220511210024.png)



**block-阻塞**

浏览器访问：http://localhost:6060/debug/pprof/block

我们开始排查最后一个阻塞问题。在程序中，除了锁的竞争会导致阻塞之外，还有很多逻辑(例如读取一个`channel`) 也会导致阻塞，在页面中可以看到阻塞操作还剩两个强调)。链接地址末尾再换成`block`

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511210234.png)

---

和刚才一样，`Graph`到`Source`的视图切换。可以看到，在*Cat.Pee0函数中读取 了一个time.After()生成的`channel`,这就导致了这个`goroutine`实际上阻塞了1秒钟，而不是等待了1秒钟。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511210520.png)

---

两个Block为什么只展示了一个

go tool pprof "http://localhost:6060/debug/pprof/block"

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511210732.png)



不过通过上面的分析，我们只定位到一个block的问题
有同学可能会发现。刚刚的计数页面上有两个阻塞操作，但是实际上只有一个， 那么另一个为什么没有展示呢?
提示，可以关注一下`pprof Top`视图中表格之外的部分，从框住的地方可以发现，有4个节点因为cumulative小于1.41秒被drop掉了 ，这就是另一个阻塞操作的节点，但他因为总用时小于总时长的千分之5,所以被省略掉了。这样的过滤策略能够更加有效地突出问题所在，而省略相对没有问题的信息。如果不作任何过滤全部展示的话，对于-个复杂的程序可能内容就会非常庞大了，不利于我们的问题定位。

---

我们知道了另-个阻塞是确实存在的，接下来大家可以进一步猜测一下：没有展示的这个阻塞操作发生在哪里? 需要点击哪个链接查看?

提示：程序除了在跑性能「炸弹」之外， 还有什么工作呢?我们是怎样访问它的采样信息的?

尽管占用低的节点不会在`pprof`工具中展示出来，但实际上他是被记录下来的。我们还可以通过暴露出来的接口地址直接访问它。
所以，打开Block指标的页面，可以看到，第二个阻塞操作发生在了http handler中，这个阻塞操作是符合预期的。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511210911.png)

### 小结：

至此，我们已经发现了代码中存在的所有「炸弹J。以上我们介绍了五种使用`pprof`采集的常用性能指标: **CPU、堆内存、Goroutine、 锁竞争和阻塞;**

两种展示方式交 互式终端和网页:四种视图: Top、Graph、 源码和火焰图。
pprof除了http的获取方式之外，也可以直接在运行时调用runtime/pprof包将指标数据输出到本地文件。视图中还有一个更底层的[反汇编」视图。有兴趣的同学可以尝试一下它们的用法。

俗话说，知其然还要知其所以然，接下来我们来看看它们内部的实现

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511211007.png)

## 2.2.3性能分析工具pprof-采样过程和原理

首先来看`CPU`。`CPU`采样会记录所有的调用栈和它们的占用时间。
在采样时，进程会每秒暂停一百次， 每次会记录当前的调用栈信息。汇总之后，根据调用栈在采样中出现的次数来推断函数的运行时间。

你需要手动地启动和停止采样。每秒100次的暂停频率也不能更改。这个定时暂停机制在`unix`或类`unix`系统上是依赖信号机制实现的。
每次「暂停」都会接收到一个信号，通过系统计时器来保证这个信号是固定频率发送的。接下来看看具体的流程。

**CPU**

- 采样对象：函数调用和它们占用的时间
- 采样率: 100次/秒，固定值
- 采样时间:从手动启动到手动结束



![](https://cdn.jsdelivr.net/gh/nateshao/images/20220512125522.png)



----

一共有三个相关角色：**进程本身、操作系统和写缓冲。**

启动采样时，进程向OS注册一个定时器，OS会每隔10ms向进程发送一个SIGPROF信号， 进程接收到信号后就会对当前的调用栈进行记录。

与此同时，进程会启动一个写缓冲的`goroutine`,它会每隔100ms从进程中读取已经记录的堆栈信息，并写入到输出流。

当采样停止时，进程向OS取消定时器，不再接收信号，写缓冲读取不到新的堆栈时，结束输出。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511213524.png)

----

接下来看看堆内存采样。
提到内存指标的时候说的都是「堆内存」而不是「内存」， 这是因为pprof的内存采样是有局限性的。

内存采样在实现上依赖了内存分配器的记录，所以它只能记录在堆上分配，且会参与GC的内存，一些其他的内存分配， 例如调用结束就会回收的栈内存、一 些 更底层使用cgo调用分配的内存，是不会被内存采样记录的。

它的采样率是一个大小， 默认每分配512KB内存会采样一 次，采样率是可以在运行开头调整的， 设为1则为每次分配都会记录。

与CPU和`goroutine`都不同的是，内存的采样是一个持续的过程， 它会记录从程序运行起的所有 分配或释放的内存大小和对象数量，并在采样时遍 历这些结果进行汇总。

还记得刚才的例子中，堆内存采样的四种指标吗? alloc的两项指标是从程序运行开始的累计指标，而inuse的两项指标是通过累计分配减去累计释放得到的程序当前持有的指标。你也可以通过比较两次alloc的差值来得到某一段时间程序 分配的内存[大小和数量]



**Heap-堆内存**

- 采样程序通过内存分配器在堆上分配和释放的内存，记录分配/释放的大小和数量
- 采样率：每分配512KB记录一次，可在运行开头修改，1为每次分配均记录
- 采样时间：从程序运行开始到采样时
- 采样指标：`alloc space`, `alloc objects`,` inuse_ space`, `inuse_ objects`
- 计算方式：` inuse = alloc- free`

----

**Goroutine-协程 & ThreadCreate-线程创建**

接下来我们来看看`goroutine`和系统线程的采样。这两个采样指标在概念上和实现上都比较相似，所以在这里进行对比。

`goroutine`采样会记录所有用户发起，也就是入口不是runtime开头的`goroutine`, 以及main函数所在`goroutine`的信息和创建这些`goroutine`的调用栈:

他们在实现上非常的相似，都是会在STW之后，遍历所有goroutine/所有线程的列表(图中的m就是 GMP模型中的m,在golang中和线程一对应) 并输出堆栈。

最后Start The Word继续运行。这个采样是立刻触发的全量记录，你可以通过比较两个时间点的差值来得到某一时间段的指标。

- Goroutine
  - 记录所有用户发起且在运行中的`goroutine`(即入口非runtime开头的)
    runtime.main的调用栈信息

- ThreadCreate
  - 记录程序创建的所有系统线程的信息



![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511222400.png)

---

**Block-阻塞&Mutex-锁**

阻塞和锁竞争这两种采样

这两个指标在流程和原理上也非常相似。两个采样记录的都是对应操作发生的调用栈、次数和耗时，不过这两个指标的采样率含义并不相同。

- 阻塞操作的采样率是一个「阈值」，消耗超过阈值时间的阻塞操作才 会被记录，1为每次操作都会记录。记得炸弹程序的main代码吗?里面设置了rate= 1

- 锁竞争的采样率是一个「比例」，运行时会通过随机数来只记录固定比例的锁操作， 1为每次操作都会记录。

它们在实现上也是基本相同的。都是一个「主动上报」的过程。

在阻塞操作或锁操作发生时，会计算出消耗的时间，连同调用栈一起主动上报给采样器，采样器会根据采样率可能会丢弃些记录。

在采样时，采样器会遍历已经记录的信息，统计出具体操作的次数、调用栈和总耗时。和堆内存一样，你可以对比两 个时间点的差值计算出段时间内的操作指标。



**阻塞操作**

- 采样阻塞操作的次数和耗时
- 采样率：阻塞耗时超过闽值的才会被记录，1为每次阻塞均记录

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511223011.png)

**锁竞争**

- 采样争抢锁的次数和耗时
- 采样率：只记录固定比例的锁操作，1为每次加锁均记录

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511223328.png)





## 2.2性能分析工具pprof

小结

1. 掌握常用`pprof`工具功能
2. 灵活运用`pprof`工具分析解决性能问题
3. 了解`pprof`的采样过程和工作原理

**在实战过程中，虽然我们只是模拟了一个使用`pprof`分析性能问题的小场景，但是大部分排查思路和想法是通用的，熟悉`pprof`工具后， 分析真正的服务问题也会得心应手，对解决性能问题和性能调优很有帮助。**

## 2.3性能调优素例

简介

1. 介绍实际业务服务性能优化的案例
2. 对逻辑相对复杂的程序如何进行性能调优

- 业务服务优化
- 基础库优化
- Go语言优化

在实际工作中，当服务规模比较小的时候，可能不会触发很多性能问题，同时性能优化带来的效果也不明显，很难体会到性能调优带来的收益。而当业务量逐渐增大，比如一个服务使用了几千台机器的时候，性能优化一个百分点，就能节省数百台机器，成本降低是非常可观的

接下来我们来了解下工程中进行性能调优的实际案例

程序从不同的应用层次上看，可以分为**业务服务、基础库和Go语言本身**三类，对应优化的适用范围也是越来越广

- 业务服务一般指直 接提供功能的程序，比如专门处理用户评论操作的程序
- 基础库一般指提供通用功能的程序，主要是针对业务服务提供功能，比如监控组件，负责收集业务服务的运行指标
- 另外还有对Go语言本身进行的优化项

## 2.3.1性能调优案例-业务服务优化

**基本概念**

- 服务：能单独部署，承载一定功能的程序
- 依赖：ServiceA的功能实现依赖ServiceB的响应结果，称为Service A依赖Service B
- 调用链路：能支持一个接口请求的相关服务集合及其相互之间的依赖关系
- 基础库:公共的工具包、中间件

![image-20220511223920769](https://cdn.jsdelivr.net/gh/nateshao/images/20220511223920.png)



那么针对逻辑相对复杂的业务服务，它的性能调优流程是怎么样的呢?在介绍真正流程之前，可能有的同学对部分名词不太了解，先介绍一下

右边是系统部署的简单示意图，客户端请求经过网关转发，由不同的业务服务处理，业务服务可能依赖其他的服务，也可能会依赖存储、消息队列等组件。接下来我们以业务服务优化为例，说明性能调优的流程，图中的Service B被Service A依赖，同时也依赖了存储和Service D

---

**流程**

- 建立服务性能评估手段
- 分析性能数据，定位性能瓶颈
- 重点优化项改造
- 优化效果验证

那么接下来就来看一下业务服务优化的主要流程，主要分四步，这些流程也是性能调优相对通用的流程，可以适用其他场景和上面评估代码优化效果的`benchmarkI`具类似，对于服务的性能也需要一个评估手段和标准优化的核心是发现服务性能的瓶颈，这里主要也是用`ppro`采样性能数据，分析服务的表现发现瓶颈后需要进行服务改造，重构代码，使用更高效的组件等

最后一步是优化效果验证，通过压测对比和正确性验证之后，服务可以上线进行实际收益评估

整体的流程可以循环并行执行，每个优化点可能不同，可以分别评估验证

---

**建立服务性能评估手段**

- 服务性能评估方式
  - 单独`benchmark`无法满足复杂逻辑分析
  - 不同负载情况下性能表现差异

- 请求流量构造
  - 不同请求参数覆盖逻辑不同
  - 线上真实流量情况
- 压测范围
  - 单机器压测
  - 集群压测

性能数据采集

- 单机性能数据
- 集群性能数据

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511224325.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511224310.png)



之所以不用`benchmark`是因为实际服务逻辑比较复杂，希望从更高的层面分析服务的性能问题，同时机器在不同负载下的性能表现也会不同，右图是负载和单核`qps`的对应数据。

另外因为逻辑复杂，不同的请求参数会走不同的处理逻辑，对应的性能表现也不相同，需要尽量模拟线上真实情况，分析真正的性能瓶颈
压测会录制线上的请求流量，通过控制回放速度来对服务进行测试，测试范围可以是单个实例，也可以是整个集群，同样性能采集也会区分单机和集群



---

**建立服务性能评估手段**

评估手段建立后，它的产出是什么呢?实际是一个服务的性能指标分析报告
实际的压测报告截图，会统计压测期间服务的各项监控指标，包括qps，延迟等内容，同时在压测过程中，也可以采集服务的`ppro`数据，使用之前的方式分析性能问题

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511224417.png)



---

**分析性能数据，定位性能瓶颈**

使用库不规范

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511224503.png)

有了服务优化前的性能报告和一些性能采样数据， 我们可以进行性能瓶颈分析了业务服务常见的性能问题可能是使用基础组件不规范

比如这里通过火焰图看出`json`的解析部分占用了较多的CPU资源，那么我们就能定位到具体的逻辑代码，是在每次使用配置时都会进行`json`解析， 拿到配置项，实际组件内部提供了缓存机制，只有数据变更的时候才需要重新解析`json`

---

还有是类似日志使用不规范，一部分是调试日志发布到线上，一部分是线上服务在不同的调用链路上数据有差别，测试场景日志量还好，但是到了真实线上全量场景，会导致日志量增加，影响性能

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511224553.png)

---

高并发场景优化不足

另外常见的性能问题就是高并发场景的优化不足，左边是服务高峰期的火焰图，右边是低峰期的火焰图，可以发现metrics, 即监控组件的CPU资源占用变化较大，主要原因是监控数据上报是同步请求，在请求量上涨，监控打点数据量增加时，达到性能瓶颈，造成阻塞，影响业务逻辑的处理，后续是改成异步上报的机制提升了性能

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511224634.png)



---

**重点优化项改造**

- 正确性是基础
- 响应数据diff
  - 线上请求数据录制回放
  - 新旧逻辑接口数据diff

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511224802.png)

定位到性能瓶颈后，我们也有了对应的修复手段，但是修改完后能直接发布上线吗?
性能优化的前提是保证正确性，所以在变动较大的性能优化上线之前，还需要进行正确性验证，因为线上的场景和流程太多，所以要借助自动化手段来保证优化后程序的正确性

同样是线上请求的录制，不过这里不仅包含请求参数录制，还会录制线上的返回内容，重放时对比线上的返回内容和优化后服务的返回内容进行正确性验证。比如图中作者信息相关的字段值在优化有有变化，需要进一步排查原因

---

**优化效果验证**

- 重复压测验证
- 上线评估优化效果
  - 关注服务监控
  - 逐步放量
  - 收集性能数据

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511224934.png)

改造完成后，可以进行优化效果验证了

验证分两部分，首先依然是用同样的数据对优化后的服务进行压测，可以看到现在的数据比优化前好很多，能够支持更多的qps
正式上线的时候会逐步放量，记录真正的优化效果。同时压测并不能保证和线上表现完全一致， 有时还要通过线上的表现再进行分析改进，是个长期的过程

---

**进一步优化，服务整体链路分析**

- 规范上游服务调用接口
- 明确场景需求，分析链路，通过业务流程优化提升服务性能

以上的内容是针对单个服务的优化过程，从更高的视角看，性能是不是还有优化空间?
在熟悉服务的整体部署情况后，可以针对具体的接口链路进行分析调优。比如Service A调用Service B是否存在重复调用的情况，调用Service B服务时，是否更小的结果数据集就能满足需求，接口是否一定要实时数据，能否在Service A层进行缓存，减轻调用压力

这种优化只使用与特定业务场景，适用范围窄 ，不过能更合理的利用资源

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511225121.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511225140.png)

## 2.3.2性能调优素例-基础库优化

**AB实验SDK的优化**

- 分析基础库核心逻辑和性能瓶颈
  - 设计完善改造方案
  - 数据按需获取
  - 数据序列化协议优化

- 内部压测验证
- 推广业务服务落地验证

适用范围更广的就是基础库的优化

比如在实际的业务服务中，为了评估某些功能上线后的效果，经常需要进行AB实验，看看不同策略对核心指标的影响，所以公司内部多数服务都会使用AB实验的SDK，如果能优化AB组件库的性能，所有用到的服务都会有性能提升

类似业务服务的优化流程，也会先统计下各个服务中AB组件的资源占用情况，看看AB组件的哪些逻辑更耗费资源，提取公共问题进行重点优化

图中看到有部分性能耗费在序列化上，因为AB相关的数据量较大，所以在制定优化方案时会考虑优化数据序列化协议，同时进行按需加载，只处理服务需要的数据。完成改造和内部压测验证后，会逐步选择线上服务进行试点放量，发现潜在的正确性和使用上的问题，不断迭代后推广到更多服务







---

**编译器&运行时优化**

- 优化内存分配策略
- 优化代码编译流程，生成更高效的程序
- 内部压测验证
- 推广业务服务落地验证
- 优点
  - 接入简单，只需要调整编译配置
  - 通用性强

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220511225645.png)

接下来是适用范围最广的优化，就是针对`Go`本身进行的优化，会优化编译器和运行时的内存分配策略，构建更高效的`go`发行版本

这样的优化业务服务接入非常简单，只要调整编译配置即可，通用性很强，几乎对所有go的程序都会生效， 比如右图中服务只是换用新的发行版本进行编译，`CPU`占用降低8%

## 总结

性能调优原则

- 要依靠数据不是猜测
  - 性能分析工具`pprof`
- 熟练使用`pprof`工具排查性能问题并了解其基本原理
- 性能调优
  - 保证正确性
  - 定位主要瓶颈

**性能调优的流程很长，这里总结下重要的点**

1. 我们性能评估要依靠数据，用实际的结果做决策
2. 对于`pprof`工具，可以通过分析实际的程序熟悉相关功能，理解基本原理，后续能够更好地解决性能问题
3. 在真正的服务性能调优流程中，链路会很长，重点是要保证正确性，不影响功能，同时定位主要问题
4. 期待大家真正掌握实操课程，项目中进行应用，项目作业中也会对服务性能进行评估

