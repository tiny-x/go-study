package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	command := exec.Command("nsenter", "-t", "3033", "-p", "-m", "--", "cat", "/root/a.tar.gz")

	open, err := os.Open("")
	if err != nil {
		fmt.Println(err)
	}

	command.Stdin = open

	if err := command.Start(); err != nil {
		fmt.Println(err)
	}

	if err := command.Wait(); err != nil {
		fmt.Println(err)
	}
}
