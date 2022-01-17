package main

import (
	"fmt"
	"github.com/containerd/cgroups"
)

func main() {
	cgroup, err := cgroups.Load(cgroups.V1, cgroups.StaticPath("/"))
	if err != nil {
		fmt.Println(err)
	}
	stats, err := cgroup.Stat(cgroups.IgnoreNotExist)
	if err != nil {
		fmt.Println(stats)
	}
}
