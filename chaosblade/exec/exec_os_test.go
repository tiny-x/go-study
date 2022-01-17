package chaosblade

import (
	"context"
	"github.com/chaosblade-io/chaosblade-exec-os/exec"
	"github.com/chaosblade-io/chaosblade-spec-go/channel"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"testing"
)

type testChannel struct {
	channel.LocalChannel
}

func (l *testChannel) GetScriptPath() string {
	return "/opt/chaosblade-0.9.0/bin"
}

func Test_exec_os(t *testing.T) {

	t.Log("test os cpu load executor")
	cpuModel := exec.NewCpuCommandModelSpec()

	action := cpuModel.Actions()[0]
	executor := action.Executor()
	executor.SetChannel(new(testChannel))

	response := executor.Exec("", context.Background(), &spec.ExpModel{
		Target:     "cpu",
		ActionName: "load",
		ActionFlags: map[string]string{
			"cpu-percent": "60",
		},
	})
	t.Log(response.ToString())
}
