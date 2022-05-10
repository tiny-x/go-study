package main

import (
	"fmt"
	"github.com/ncw/directio"
	"os"
)

// 1M
var block = 1024 * 1024

func main() {
	blocks := make([]byte, block)
	for {
		f, err := directio.OpenFile("/tmp/1.log", os.O_CREATE|os.O_WRONLY, 0666)
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
