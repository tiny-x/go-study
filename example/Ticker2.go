package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
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
