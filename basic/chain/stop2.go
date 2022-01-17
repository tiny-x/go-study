package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	quit := make(chan bool, 1)

	go func(quit <-chan bool) {
		for x := range quit {
			if x {
				fmt.Printf("closing")
				runtime.Goexit()
			} else {
				fmt.Printf("start")
			}
		}
	}(quit)

	quit <- false

	time.Sleep(time.Second * 10)
	quit <- true
	quit <- false

	// 等待两秒后看看有没有输出 start
	time.Sleep(time.Second * 2)
}
