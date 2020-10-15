package main

import (
	"context"
	"github.com/chaosblade-io/chaosblade-exec-docker/exec"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

func main() {

	executor := exec.NewNetWorkSidecarExecutor()

	executor.Exec("abc", context.Background(), &spec.ExpModel{
		Target:     "network",
		ActionName: "delay",
		ActionFlags: map[string]string{
			"time":         "3000",
			"interface":    "eth0",
			"local-port":   "8080",
			"container-id": "a500c599e284",
			"debug":        "true",
		},
	})

}
