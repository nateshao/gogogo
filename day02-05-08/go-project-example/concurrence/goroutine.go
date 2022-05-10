package concurrence

import (
	"fmt"
	"sync"
	"time"
)

func hello(i int) {
	println("hello world : " + fmt.Sprint(i))
}

func HelloGoRoutine() {
	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}
	time.Sleep(time.Second)
}

func ManyGo() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	wg.Wait()
}
