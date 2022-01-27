package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tiny-x/go-study/nsexec"
	"os"
)

func main() {

	process := &nsexec.Process{Cmd: "ls",
		Args:   nil,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	//p, err := process.Exec("6676", "/opt/chaosblade/blade create cpu load --cpu-percent 60 -d")
	p, err := process.Exec("6676", "top")
	fmt.Println(p.Pid)

	if err != nil {
		logrus.Infoln("exec err: %v", err)
	}
}
