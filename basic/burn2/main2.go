/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"

	"github.com/chaosblade-io/chaosblade-exec-os/exec/bin"

	_ "go.uber.org/automaxprocs/maxprocs"
)

var (
	burnCpuStart, burnCpuStop, burnCpuNohup bool
	cpuCount, cpuPercent, climbTime         int
	slopePercent                            float64
	cpuList                                 string
	cpuProcessor                            string
)

func main() {

	burnCpu2()
}

func burnCpu2() {
	cpuPercent = 60
	cpuCount = runtime.GOMAXPROCS(cpuCount)

	var totalCpuPercent []float64
	var curProcess *process.Process
	var curCpuPercent float64
	var err error

	totalCpuPercent, err = cpu.Percent(time.Second, false)
	if err != nil {
		bin.PrintErrAndExit(err.Error())
	}

	curProcess, err = process.NewProcess(int32(os.Getpid()))
	if err != nil {
		bin.PrintErrAndExit(err.Error())
	}

	curCpuPercent, err = curProcess.CPUPercent()
	if err != nil {
		bin.PrintErrAndExit(err.Error())
	}

	otherCpuPercent := (100.0 - (totalCpuPercent[0] - curCpuPercent)) / 100.0
	go func() {
		t := time.NewTicker(3 * time.Second)
		for {
			select {
			// timer 3s
			case <-t.C:
				totalCpuPercent, err = cpu.Percent(time.Second, false)
				if err != nil {
					bin.PrintErrAndExit(err.Error())
				}

				curCpuPercent, err = curProcess.CPUPercent()
				if err != nil {
					bin.PrintErrAndExit(err.Error())
				}
				otherCpuPercent = (100.0 - (totalCpuPercent[0] - curCpuPercent)) / 100.0
				fmt.Println(slopePercent, totalCpuPercent[0], otherCpuPercent, curCpuPercent)
			}
		}
	}()

	if climbTime == 0 {
		slopePercent = float64(cpuPercent)
	} else {
		var ticker *time.Ticker = time.NewTicker(1 * time.Second)
		slopePercent = totalCpuPercent[0]
		var startPercent = float64(cpuPercent) - slopePercent
		go func() {
			for range ticker.C {
				if slopePercent < float64(cpuPercent) {
					slopePercent += startPercent / float64(climbTime)
				} else if slopePercent > float64(cpuPercent) {
					slopePercent -= startPercent / float64(climbTime)
				}
			}
		}()
	}

	for i := 0; i < cpuCount; i++ {
		go func() {
			busy := int64(0)
			idle := int64(0)
			all := int64(10000000)
			dx := 0.0
			ds := time.Duration(0)
			for i := 0; ; i = (i + 1) % 1000 {
				startTime := time.Now().UnixNano()
				if i == 0 {
					dx = (slopePercent - totalCpuPercent[0]) / otherCpuPercent
					fmt.Println("abc: ", slopePercent, totalCpuPercent[0], otherCpuPercent, dx)
					busy = busy + int64(dx*100000)
					if busy < 0 {
						busy = 0
					}
					idle = all - busy
					if idle < 0 {
						idle = 0
					}
					ds, _ = time.ParseDuration(strconv.FormatInt(idle, 10) + "ns")
				}
				for time.Now().UnixNano()-startTime < busy {
				}
				time.Sleep(ds)
				runtime.Gosched()
			}
		}()
	}
	select {}
}
