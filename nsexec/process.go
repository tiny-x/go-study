package nsexec

import (
	"fmt"
	_ "github.com/tiny-x/go-study/nsexec/nsenter"
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

func (p *Process) Exec(target string, cmd string) (*os.Process, error) {
	fmt.Print("exec.....\n")

	command := exec.Command("/proc/self/exe")

	err := os.Setenv(ENV_EXEC_PID, target)
	if err != nil {
		return nil, err
	}

	err = os.Setenv(ENV_EXEC_CMD, cmd)
	if err != nil {
		return nil, err
	}

	//containerEnvs, err := getEnvsByPid(target)
	if err != nil {
		return nil, err
	} else {
		//command.Env = append(os.Environ(), containerEnvs...)

		command.Stdin = p.Stdin

		command.Stdout = p.Stdout
		command.Stderr = p.Stderr

		if err := command.Start(); err != nil {
			return nil, err
		}
		pid := command.Process.Pid
		fmt.Println(pid)
		fmt.Println("----------------------------")
		command.Wait()
	}
	return command.Process, nil
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
