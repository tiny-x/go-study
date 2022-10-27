package main

import (
	"fmt"
	"os"
	"os/user"
	"time"
)

func main() {

	user, err := user.Lookup("xf.yefei1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(user.Uid)
	fmt.Println(user.Username)

	c := make(chan int64)
	for {
		select {
		case <-c:
			fmt.Println("yy")
		case <-time.After(time.Second):

		default:
			//fmt.Println("xx")
		}
	}
}
