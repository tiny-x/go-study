package chaosblade

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-exec-docker/exec"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"testing"
)

func Test_main(t *testing.T) {

	logrus.SetReportCaller(true)

	logrus.SetOutput(io.MultiWriter(&lumberjack.Logger{}, os.Stdout))
	logrus.SetLevel(logrus.DebugLevel)
	executor := exec.NewRunCmdInContainerExecutorByCP()
	//c := context.WithValue(context.Background(), "suid", "")
	response := executor.Exec("abc", context.Background(), &spec.ExpModel{
		Target:     "jvm",
		ActionName: "full-gc",
		ActionFlags: map[string]string{
			"container-name":     "zoo_test",
			"pid":                "1",
			"chaosblade-release": "/opt/chaosblade-1.5.0-linux-arm64.tar.gz",
		},
	})
	fmt.Print(response)

	executor2 := exec.NewNetWorkSidecarExecutor()
	c := context.WithValue(context.Background(), "suid", "abc")
	response = executor2.Exec("", c, &spec.ExpModel{
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
