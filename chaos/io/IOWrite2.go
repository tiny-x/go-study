package main

import (
	"fmt"
	"os"
)

var b = 1024 * 1024

func main() {
	blocks := make([]byte, b)
	for {
		f, err := os.OpenFile("/tmp/1.log", os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Print(err)
		}
		for i := 0; i < 1024; i++ {
			f.Write(blocks)
		}
		f.Close()
		os.Remove("/tmp/1.log")
	}
}
