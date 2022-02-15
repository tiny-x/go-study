package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("top")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}

}
