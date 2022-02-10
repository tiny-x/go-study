package main

import "fmt"

func main() {
	ifx("x", "y")
	fmt.Println("------------")
	ifx("x", "")
	fmt.Println("------------")
	ifx("", "y")
	fmt.Println("------------")
	ifx("", "")
}

func ifx(x, y string) {
	if x != "" && y != "" {
		fmt.Println(x, y)
	} else if x != "" {
		fmt.Println(x)
	} else if y != "" {
		fmt.Println(y)
	} else {
		fmt.Println("no")
	}
}
