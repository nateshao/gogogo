


Go配置Linux环境
```go
[root@iZwz97sm51bf4kxapt7xzvZ go]# ll
total 138384
-rw-r--r-- 1 root root 141699677 Apr 16 11:30 go1.18.1.linux-amd64.tar.gz
-rw-r--r-- 1 root root        75 Apr 16 11:33 hello.go
[root@iZwz97sm51bf4kxapt7xzvZ go]#  tar -C /usr/local -xzf go1.18.1.linux-amd64.tar.gz 
[root@iZwz97sm51bf4kxapt7xzvZ go]# export PATH=$PATH:/usr/local/go/bin
[root@iZwz97sm51bf4kxapt7xzvZ go]# export PATH=$PATH:/usr/local/go/bin
[root@iZwz97sm51bf4kxapt7xzvZ go]# source ~/.bash_profile
[root@iZwz97sm51bf4kxapt7xzvZ go]# go run hello.go 
你好，千羽[root@iZwz97sm51bf4kxapt7xzvZ go]# 
```


