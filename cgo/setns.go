package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	_ "github.com/tiny-x/go-study/cgo/ns"
	"os"
	"os/exec"
)

const ENV_EXEC_PID = "mydocker_pid"
const ENV_EXEC_CMD = "mydocker_cmd"

func main() {
	fmt.Println("hell ns")

	cmd := exec.Command("/proc/self/exe", "exec")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	os.Setenv(ENV_EXEC_PID, "20701")
	os.Setenv(ENV_EXEC_CMD, "ls -la")

	output, err := cmd.CombinedOutput()
	s := string(output)
	if err != nil {
		logrus.Infof("exec error, result: [%s] err: %v", s, err)
	}
	logrus.Infof("exec success, result: [%s]", s)
}
