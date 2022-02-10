package main

import (
	"errors"
	"fmt"
	"github.com/containerd/cgroups"
	"os/exec"
)

func main() {
	command := exec.Command("nsenter", "-t", "3033", "-p", "-m", "--",
		"/opt/chaosblade/bin/chaos_burncpu", "--nohup", "--cpu-count=4", "--cpu-percent=100", "--climb-time=0")
	
	control, err := cgroups.Load(cgroups.V1, PidPath(3033))
	if err != nil{
		fmt.Println(err)
	}

	if err := command.Start(); err != nil {
		fmt.Println(err)
	}

	if err = control.Add(cgroups.Process{Pid: command.Process.Pid}); err != nil {
		if err := command.Process.Kill(); err != nil {
			fmt.Println(err)
		}
	}

	if err := command.Wait(); err != nil {
		fmt.Println(err)
	}

}

func PidPath(pid int) cgroups.Path {
	p := fmt.Sprintf("/proc/%d/cgroup", pid)
	paths, err := cgroups.ParseCgroupFile(p)
	if err != nil {
		return func(_ cgroups.Name) (string, error) {
			return "", fmt.Errorf("failed to parse cgroup file %s: %s", p, err.Error())
		}
	}

	return func(name cgroups.Name) (string, error) {
		root, ok := paths[string(name)]
		if !ok {
			if root, ok = paths["name="+string(name)]; !ok {
				return "", errors.New("controller is not supported")
			}
		}

		return root, nil
	}
}