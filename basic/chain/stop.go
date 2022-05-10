package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	quit := make(chan bool, 1)
	defer close(quit)

	var httpServer http.Server
	go func() {
		httpServer = http.Server{Addr: fmt.Sprintf(`:%d`, 9000)}
		httpServer.ListenAndServe()
	}()

	go func() {
		<-quit
		fmt.Println("aaaaa")
	}()

	go func() {
		quit := <-quit
		if quit {
			httpServer.Close()
			fmt.Println("http server stop")
		}
	}()

	go func() {
		<-quit
		fmt.Println("bbbbb")
	}()

	time.Sleep(time.Second * 5)
	quit <- true

	time.Sleep(time.Second * 10)
}
