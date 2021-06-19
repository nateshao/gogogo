package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//resp, _ := http.Get("http://www.baidu.com")
	//fmt.Println(resp)
	resp, _ := http.Get("http://www.baidu.com")
	defer resp.Body.Close()
	// 200 OK
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	buf := make([]byte, 1024)
	for {
		// 接收服务端信息
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println(res)
			break
		}
	}
}
