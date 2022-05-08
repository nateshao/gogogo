


## Go配置Linux环境
```go
[root@iZwz97sm51bf4kxapt7xzvZ go]# ll
total 138384
-rw-r--r-- 1 root root 141699677 Apr 16 11:30 go1.18.1.linux-amd64.tar.gz
-rw-r--r-- 1 root root        75 Apr 16 11:33 hello.go
[root@iZwz97sm51bf4kxapt7xzvZ go]#  tar -C /usr/local -xzf go1.18.1.linux-amd64.tar.gz 
[root@iZwz98qdx9tvkknmzwtbl9Z ~]# vi  ~/.bash_profile  # 文件末尾添加 export PATH=$PATH:/usr/local/go/bin
[root@iZwz98qdx9tvkknmzwtbl9Z ~]# source ~/.bash_profile
[root@iZwz98qdx9tvkknmzwtbl9Z ~]# cd go/
[root@iZwz98qdx9tvkknmzwtbl9Z go]# go run hello.go 
你好，千羽[root@iZwz98qdx9tvkknmzwtbl9Z go]#
```


## 标识符
标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。

以下是有效的标识符：

mahesh   kumar   abc   move_name   a_123
myname50   _temp   j   a23b9   retVal
以下是无效的标识符：

1ab（以数字开头）
case（Go 语言的关键字）
a+b（运算符是不允许的）
--- 

在go语言中，单纯地给 a 赋值也是不够的，这个值必须被使用，所以使用

## Go 语言常量
常量是一个简单值的标识符，在程序运行时，不会被修改的量。

常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

- 显式类型定义： const b string = "abc"

- 隐式类型定义： const b = "abc"

多个相同类型的声明可以简写为：
```go
const c_name1, c_name2 = value1, value2
```
## Go 语言切片(Slice)

Go 语言切片是对数组的抽象。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

定义切片
你可以声明一个未指定大小的数组来定义切片：

var identifier []type
切片不需要说明长度。

或使用 make() 函数来创建切片:

var slice1 []type = make([]type, len)

也可以简写为

slice1 := make([]type, len)
也可以指定容量，其中 capacity 为可选参数。

make([]T, length, capacity)
这里 len 是数组的长度并且也是切片的初始长度。

切片初始化
s :=[] int {1,2,3 } 
直接初始化切片，[] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。

s := arr[:] 
初始化切片 s，是数组 arr 的引用。

s := arr[startIndex:endIndex] 
将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片。

s := arr[startIndex:] 
默认 endIndex 时将表示一直到arr的最后一个元素。

s := arr[:endIndex] 
默认 startIndex 时将表示从 arr 的第一个元素开始。

s1 := s[startIndex:endIndex] 
通过切片 s 初始化切片 s1。

s :=make([]int,len,cap) 
通过内置函数 make() 初始化切片s，[]int 标识为其元素类型为 int 的切片。

len() 和 cap() 函数
切片是可索引的，并且可以由 len() 方法获取长度。

切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。




