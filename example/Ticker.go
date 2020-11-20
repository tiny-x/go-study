package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	go func() {
		duration, err := time.ParseDuration("24h")
		if err != nil {
			fmt.Println(err.Error())
		}
		ticker := time.NewTicker(time.Second * time.Duration(duration.Seconds()))
		for range ticker.C {
			duration, err := time.ParseDuration("24x")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(duration.Seconds())
		}
	}()
	fmt.Println("start Ticker")

	// Wait for exit signals
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGKILL)
	for s := range ch {
		switch s {
		case syscall.SIGHUP, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt:
			fmt.Println("caught interrupt, exit")
			return
		}
	}
	fmt.Println("end Ticker")
}

func funcName() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	ticker := time.NewTicker(time.Second)
TickerLoop:
	for range ticker.C {
		select {
		case <-ctx.Done():
			ticker.Stop()
			break TickerLoop
		default:
			fmt.Println("abc")
		}
	}
}
