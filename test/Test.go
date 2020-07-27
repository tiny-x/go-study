package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second * 5)
	go func() {
		fmt.Printf("ticked at %v", time.Now())
		for _ = range ticker.C {
			fmt.Printf("ticked at %v", time.Now())
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGKILL)
	for s := range ch {
		switch s {
		case syscall.SIGHUP, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt:
			fmt.Println("caught interrupt, exit")
			return
		}
	}
}
