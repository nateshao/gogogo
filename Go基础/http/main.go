package main

import (
	"fmt"
	"net/http"
)

// handler函数
func myHander(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功。。")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method: ", r.Method)
	fmt.Println("url: ", r.URL)
	fmt.Println("host:", r.Host)
	fmt.Println("form", r.Form)
	fmt.Println("header", r.Header)

	// 回复
	w.Write([]byte("www.baidu.com"))
}

func main() {
	//http://127.0.0.1:8080/go
	http.HandleFunc("/go", myHander)
	// 单独写回调函数
	//http.HandleFunc("/ungo",myHandler2 )
	// addr：监听的地址
	// handler：回调函数
	http.ListenAndServe("127.0.0.1:8080", nil)

}
