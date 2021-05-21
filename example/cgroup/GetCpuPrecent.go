package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
	"os"
	"time"
)

func main() {
	totalCpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Print(err.Error())
	}

	curProcess, err := process.NewProcess(int32(os.Getpid()))

	fmt.Println(totalCpuPercent[0])
	fmt.Print(curProcess.CPUPercent())
}
