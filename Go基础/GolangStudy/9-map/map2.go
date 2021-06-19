package main

import (
	"fmt"
)

func prinyMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key = ", key, "value = ", value)
	}
}

func changeMap(cityMap map[string]string) {
	cityMap["Japan"] = "BBB"
}

func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "ShenZhen"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"
	// 遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key, ",value = ", value)
		//fmt.Println("value = ", value)
	}
	// 删除
	fmt.Println("-------删除-------")
	delete(cityMap, "China")

	prinyMap(cityMap)

	changeMap(cityMap)
	// 修改
	cityMap["USA"] = "AAAA"

	for key, value := range cityMap {
		fmt.Println("key = ", key, ",value = ", value)

	}
}
