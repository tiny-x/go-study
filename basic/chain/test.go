package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

func task() (bool, error) {
	return true, nil
}

func main() {
	strings := make(chan string, 1)
	defer close(strings)

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			strings <- strconv.Itoa(i)
			time.Sleep(1000)
			runtime.Goexit()
		}()
	}

	s := &student{}
	value := reflect.ValueOf(s).Elem()

	conn := value.FieldByName("age")
	if conn.CanAddr() {

	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-strings)
	}
	httpServer := &http.Server{Addr: fmt.Sprintf(`:%d`, 9000)}

	if err := httpServer.ListenAndServe(); err != nil {
	}
}

type student struct {
	age int
}
