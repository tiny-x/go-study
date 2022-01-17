package ns

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Process struct {
	Cmd    string    `json:"cmd,omitempty"`
	Args   []string  `json:"args,omitempty"`
	Stdin  io.Reader `json:"stdin,omitempty"`
	Stdout io.Writer `json:"stdout,omitempty"`
	Stderr io.Writer `json:"stderr,omitempty"`
}

const ENV_EXEC_PID = "pid"
const ENV_EXEC_CMD = "cmd"

func (p *Process) Exec(target string, cmd string) error {

	command := exec.Command("/proc/self/exe", "exec")

	err := os.Setenv(ENV_EXEC_PID, target)
	if err != nil {
		return err
	}

	err = os.Setenv(ENV_EXEC_CMD, cmd)
	if err != nil {
		return err
	}

	containerEnvs, err := getEnvsByPid(target)
	if err != nil {
		return err
	} else {
		command.Env = append(os.Environ(), containerEnvs...)

		command.Stdin = p.Stdin

		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		if err := command.Run(); err != nil {
			return err
		}
	}
	return nil
}

func getEnvsByPid(pid string) ([]string, error) {
	path := fmt.Sprintf("/proc/%s/environ", pid)
	contentBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	//env split by \u0000
	envs := strings.Split(string(contentBytes), "\u0000")
	return envs, nil
}
