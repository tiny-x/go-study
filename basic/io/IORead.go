package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("/dev/zero")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}
