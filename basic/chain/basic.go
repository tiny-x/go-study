package main

import (
	"fmt"
	"time"
)

var message = make(chan string)

func go1() {
	message <- "hello1"
	message <- "hello2"
	message <- "hello3"
	//message <- "hello4"
}
func go2() {
	time.Sleep(2 * time.Second)
	str := <-message
	str = str + "go"
	message <- str
}
func main() {
	go go1()
	go go2()
	time.Sleep(3 * time.Second)
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
