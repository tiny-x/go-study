package main

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-exec-os/exec"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

func main() {
	executor := exec.NewSSHExecutor()
	response := executor.Exec(
		"abcde",
		context.Background(),
		&spec.ExpModel{
			Target:     "cpu",
			ActionName: "load",
			ActionFlags: map[string]string{
				"channel":            "ssh",
				"ssh-host":           "101.37.30.161",
				"ssh-user":           "root",
				"ssh-key":            "/Users/yefei/.ssh/hg.pem",
				"ssh-key-passphrase": "true",
				"debug":              "true",
			},
		},
	)
	fmt.Print(response.ToString())
}
