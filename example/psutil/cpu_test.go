package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"testing"
	"time"
)

func Test_cpu(t *testing.T) {

	for true {
		p, _ := cpu.Percent(time.Second, false)
		time.Sleep(1)
		fmt.Printf(time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println(p[0])
	}
}
