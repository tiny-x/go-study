package chaosblade

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-exec-os/exec"
	"github.com/chaosblade-io/chaosblade-exec-os/version"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"testing"
)

func Test_exec_ssh(t *testing.T) {
	executor := exec.NewSSHExecutor()
	level, _ := logrus.ParseLevel("debug")
	version.BladeVersion = "0.7.0"
	logrus.SetLevel(level)
	// value := context.WithValue(context.Background(), "suid", "abcde")
	value := context.Background()
	uid, _ := uuid.NewV4()
	response := executor.Exec(
		uid.String(),
		value,
		&spec.ExpModel{
			Target:     "cpu",
			ActionName: "load",
			ActionFlags: map[string]string{
				"channel":              "ssh",
				"ssh-host":             "101.37.30.161",
				"ssh-user":             "root",
				"ssh-key":              "/Users/yefei/.ssh/hg.pem",
				"debug":                "true",
				exec.BladeRelease.Name: "https://chaosblade.oss-cn-hangzhou.aliyuncs.com/agent/github/0.6.0/chaosblade-0.6.0-linux-amd64.tar.gzz",
				exec.InstallPath.Name:  "/opt/abcd",
			},
		},
	)
	fmt.Print(response.ToString())
}
