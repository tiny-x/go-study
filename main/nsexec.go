package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tiny-x/go-study/ns"
	_ "github.com/tiny-x/go-study/nsexec"
	"os"
)

func main() {

	process := &ns.Process{Cmd: "ls",
		Args:   nil,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err := process.Exec("3514", "date")

	if err != nil {
		logrus.Errorf("exec err: %v", err)
	}
}
