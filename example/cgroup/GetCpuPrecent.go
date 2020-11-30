package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

func main() {
	totalCpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Print(err.Error())
	}

	format := totalCpuPercent[0]
	fmt.Print(format)
}
