package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func main() {
	name, i := GetRunFuncName()
	logrus.WithFields(logrus.Fields{"uid": "abc", "file": name, "line": i}).Infoln("aa", name)

	cmd := exec.Command("blade", "version")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := os.Setenv("PATH", "/tmp")
	fmt.Printf("set env err, %v", err)
	environ := os.Environ()
	fmt.Println(environ)

	fmt.Println(cmd.Env)

	cmd.SysProcAttr = &syscall.SysProcAttr{}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}

}

func GetRunFuncName() (string, int) {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.FileLine(pc[0])
}
