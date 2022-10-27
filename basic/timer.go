package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(1 * time.Second)
	for i := 0; i < 1; i++ {
		<-timer.C
	}

	time.AfterFunc(time.Second, func() {
		fmt.Println("xxx")
	})
}
