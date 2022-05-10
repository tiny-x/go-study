package main

import "fmt"

func main() {

	c := make(chan int64)
	for {
		select {
		case <-c:
			fmt.Println("yy")
		default:
			fmt.Println("xx")
		}
	}
}
