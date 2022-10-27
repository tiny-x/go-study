package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 4; i++ {
		go func() {
			for {
				stress_cpu(time.Second, 90)
			}
		}()
	}
	select {}
}

func stress_cpu(interval time.Duration, cpuPercent float64) {
	bias := 0.0
	startTime := time.Now().UnixNano()
	nanoInterval := int64(interval / time.Nanosecond)
	fmt.Printf("[%d]=========nanoInterval\n", nanoInterval)
	for {
		if time.Now().UnixNano()-startTime > nanoInterval {
			break
		}
		startTime1 := time.Now().UnixNano()
		// Loops and methods may be specified later.
		ackermann(3, 7)
		endTime1 := time.Now().UnixNano()
		fmt.Println(startTime1, endTime1, cpuPercent)
		delay := ((100 - cpuPercent) * float64(endTime1-startTime1) / cpuPercent)
		fmt.Printf("delay : [%f], bias : [%f]\n", delay, bias)
		delay -= bias
		if delay <= 0.0 {
			bias = 0.0
		} else {
			startTime2 := time.Now().UnixNano()
			time.Sleep(time.Duration(delay) * time.Nanosecond)
			endTime2 := time.Now().UnixNano()
			bias = float64(endTime2-startTime2) - delay
		}
	}
}

func ackermann(m uint32, n uint32) uint32 {
	if m == 0 {
		return n + 1
	} else if n == 0 {
		return ackermann(m-1, 1)
	} else {
		return ackermann(m-1, ackermann(m, n-1))
	}
}
