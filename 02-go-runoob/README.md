


Go配置Linux环境
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


