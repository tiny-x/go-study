package main

import (
	"container/list"
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"time"
)

func main() {

	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}

	readBytesQueue := make(map[string]*list.List)
	preReadBytes := make(map[string]uint64)

	writeBytesQueue := make(map[string]*list.List)
	preWriteBytes := make(map[string]uint64)
	for true {
		ioStat, _ := disk.IOCounters()
		for k, v := range ioStat {
			time.Sleep(time.Second * 1)

			r := readBytesQueue[k]
			if r == nil {
				readBytesQueue[k] = list.New()
			}
			if preReadBytes[k] != 0 {
				diff := uint((v.ReadBytes - preReadBytes[k]) / 1024)
				fmt.Printf("差值: %d", diff)
				fmt.Println("----------------------------")
				readBytesQueue[k].PushBack(diff)
				if readBytesQueue[k].Len() > 5 {
					readBytesQueue[k].Remove(readBytesQueue[k].Front())
				}
			}
			var rTemp uint
			for e := readBytesQueue[k].Front(); e != nil; e = e.Next() {
				rTemp = rTemp + e.Value.(uint)
				fmt.Printf("%v, ", e.Value)
			}
			fmt.Println()

			w := writeBytesQueue[k]
			if w == nil {
				writeBytesQueue[k] = list.New()
			}
			if preWriteBytes[k] != 0 {
				diff := uint((v.WriteBytes - preWriteBytes[k]) / 1024)
				fmt.Printf("差值: %d", diff)
				fmt.Println("----------------------------")
				writeBytesQueue[k].PushBack(diff)
				if writeBytesQueue[k].Len() > 5 {
					writeBytesQueue[k].Remove(writeBytesQueue[k].Front())
				}
			}
			var wTemp uint
			for x := writeBytesQueue[k].Front(); x != nil; x = x.Next() {
				wTemp = wTemp + x.Value.(uint)
				fmt.Printf("%v, ", x.Value)
			}
			fmt.Println()

			fmt.Printf("%v: read: %d, write: %d, c_read: %dKB, c_write: %dKB ",
				k,
				v.ReadBytes/1024,
				v.WriteBytes/1024,
				func() uint {
					if readBytesQueue[k].Len() == 0 {
						return 0
					}
					return rTemp / uint(readBytesQueue[k].Len())
				}(),
				func() uint {
					if writeBytesQueue[k].Len() == 0 {
						return 0
					}
					return wTemp / uint(writeBytesQueue[k].Len())
				}(),
			)
			preReadBytes[k] = v.ReadBytes
			preWriteBytes[k] = v.WriteBytes
			fmt.Println()
		}
	}

}
