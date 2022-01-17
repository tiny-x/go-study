package main

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"testing"
	"time"
)

func Test_speed(t *testing.T) {

	var preBytesSent uint64
	var preBytesRecv uint64

	for true {
		time.Sleep(time.Second)
		counters, _ := net.IOCounters(false)
		for _, c := range counters {
			fmt.Println(fmt.Sprintf("网卡: %s, 发送字节数：%d, 接受字节数：%d, 发送速度: %dKB, 接受速度: %dKB",
				c.Name,
				c.BytesSent,
				c.BytesRecv,
				(c.BytesSent-preBytesSent)/1024,
				(c.BytesRecv-preBytesRecv)/1024,
			))

			preBytesSent = c.BytesSent
			preBytesRecv = c.BytesRecv
		}
	}

}
