package main

import (
	"fmt"
	"strconv"
	"time"
)

func go1(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- "hello" + strconv.Itoa(i)
		//time.Sleep(1*time.Second)
	}
}
func go2(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		//time.Sleep(2*time.Second)
	}
}

func main() {
	chan1 := make(chan string, 3)
	chan2 := make(chan int, 5)

	go go1(chan1)
	go go2(chan2)
	time.Sleep(1 * time.Second)
Loop:
	for {
		select {
		case str, ok := <-chan1:
			if !ok {
				fmt.Println("ch1 failed")
				continue
			}
			fmt.Println(str)

		case p, oke := <-chan2:
			if !oke {
				fmt.Println("ch2 failed")
				continue
			}
			fmt.Println(p)
		default:
			fmt.Println("ds")
			break Loop
		}
	}

}
