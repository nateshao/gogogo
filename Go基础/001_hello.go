package main

// 单个包可以这样
// import "fmt"
// 多个包
import (
	"fmt"
	"time"
)

func main() {
	//fmt.Print()// 不换行
	//fmt.Println()// 换行
	//fmt.Printf()// 格式化输出
	// fmt.Println("你好，世界！");
	// fmt.Println("hello world!");
	fmt.Println("hello world")
	time.Sleep(100)
	fmt.Println("你好，世界！")
	year := 2023
	offer := 1
	happy := 1
	money := 1
	if year == 2023 {
		money++ // 钱兔无量
		offer++ // offer收割机
		happy++ // 幸福多多
	}

}

//func RemoteFunc(ctx context.Context, x int) (int, error) {
//	ctx2, defer_func := context.WithTimeout(ctx, time.Second)
//	defer defer_func()
//	res, err := grpc_client.Calculate(ctx2, x*2)
//	return res, err
//}
//
//func RemoteFuncRetry(ctx context.Context, x int) (res int, err error) {
//	for i := 0; i < 3; i++ {
//		if res, err = RemoteFunc(ctx, x); err == nil {
//			return
//		}
//	}
//	return
//}
