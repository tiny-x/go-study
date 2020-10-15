package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"os"
)

func main() {

	username, _ := getUsername(os.Getuid())
	fmt.Print(username)
}

func getUsername(pid int) (string, error) {

	javaProcess, err := process.NewProcess(int32(pid))
	if err != nil {
		return "", err
	}
	return javaProcess.Username()
}
