package main

import "io/ioutil"

func main() {
	var bytes = []byte("he")
	ioutil.WriteFile("/dev/null", bytes, 777)
}
