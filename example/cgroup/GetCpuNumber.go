package main

import (
	"fmt"
	"github.com/containerd/cgroups"
)

func main() {
	// Your application logic here.
	cgroup, err := cgroups.Load(cgroups.V1, cgroups.StaticPath("/"))
	if err != nil {
		fmt.Println(fmt.Errorf("load cgroup error, %v", err))
	}
	stats, err := cgroup.Stat(cgroups.IgnoreNotExist)
	if err != nil {
		fmt.Println(fmt.Errorf("load cgroup stat error, %v", err))
	}
	cpu := stats.CPU
	fmt.Println(fmt.Printf("size: %d", cpu.XXX_Size()))
}
