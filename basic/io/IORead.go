package main

import (
	"fmt"
	"github.com/ncw/directio"
	"io"
	"os"
)

func main() {
	var block = 1024 * 1024
	blocks := make([]byte, block)
	f, err := directio.OpenFile("/tmp/1.log", os.O_CREATE|os.O_WRONLY|os.O_RDONLY, 0666)
	if err != nil {
		fmt.Print(err)
	}
	for i := 0; i < 1024; i++ {
		f.Write(blocks)
	}
	f.Close()

	f, err = directio.OpenFile("/tmp/1.log", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	for {
		_, err := io.ReadFull(f, blocks)
		if err != nil {
			if err == io.EOF {
				f, err = directio.OpenFile("/tmp/1.log", os.O_RDONLY, 0666)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		}
	}
}
