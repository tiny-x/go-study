package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tiny-x/go-study/nsexec"
)

func main() {

	process := &nsexec.Process{Cmd: "",
		Args:   nil,
		Stdin:  nil,
		Stdout: nil,
		Stderr: nil,
	}
	err := process.Exec("", "")

	if err != nil {
		logrus.Errorf("exec err: %v", err)
	}
}
