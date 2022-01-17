package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	_ "github.com/tiny-x/go-study/cgo/ns"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const ENV_EXEC_PID = "mydocker_pid"
const ENV_EXEC_CMD = "mydocker_cmd"

func main() {
	pid := "20701"
	fmt.Println("hell ns")

	cmd := exec.Command("/proc/self/exe", "exec")

	os.Setenv(ENV_EXEC_PID, pid)
	os.Setenv(ENV_EXEC_CMD, "/opt/chaosblade/blade version")

	containerEnvs := getEnvsByPid(pid)
	cmd.Env = append(os.Environ(), containerEnvs...)

	open, err := os.Open("/root/node_exporter-1.3.1.linux-amd64.tar.gz")
	if err != nil {
		logrus.Errorf("open file err, %v", err)
	}
	cmd.Stdin = open

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		logrus.Errorf("Exec container %v error", err)
	}
}

func getEnvsByPid(pid string) []string {
	path := fmt.Sprintf("/proc/%s/environ", pid)
	contentBytes, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Errorf("Read file %s error %v", path, err)
		return nil
	}
	//env split by \u0000
	envs := strings.Split(string(contentBytes), "\u0000")
	return envs
}
