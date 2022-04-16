


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


