package main

import (
	"io/ioutil"
)

func main() {

	ioutil.ReadFile("/dev/zero")
}
