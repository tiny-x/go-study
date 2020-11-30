package chaosblade

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-exec-docker/exec"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"testing"
)

func Test_main(t *testing.T) {
	executor := exec.NewNetWorkSidecarExecutor()
	c := context.WithValue(context.Background(), "suid", "abc")
	response := executor.Exec("", c, &spec.ExpModel{
		Target:     "network",
		ActionName: "loss",
		ActionFlags: map[string]string{
			"image-version": "0.7.0",
			"image-repo":    "registry.cn-hangzhou.aliyuncs.com/chaosblade/chaosblade-tool",
			"container-id":  "751a4568af32",
			"interface":     "eth0",
			"percent":       "100",
			"local-port":    "80",
		},
	})
	fmt.Print(response)
}
