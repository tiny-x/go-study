package main

import (
	"fmt"
	"strconv"
	"time"
)

func task() (bool, error)  {
	return true, nil
}

func main() {
	strings := make(chan string, 1)
	defer close(strings)

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			time.Sleep(1000)
			strings <- strconv.Itoa(i)
		}()
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-strings)
	}
}
