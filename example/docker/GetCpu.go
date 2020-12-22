package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/client"
	"io/ioutil"
	"time"
)

func main() {
	opts, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.24"))
	if err != nil {
		fmt.Print(err.Error())
	}
	stats, err := opts.ContainerStats(context.Background(), "5b508ee3a0d1", true)

	defer stats.Body.Close()

	decoder := json.NewDecoder(stats.Body)
	for decoder.More() {
		bytes, err := ioutil.ReadAll(decoder.Buffered())
		if err != nil {
			fmt.Print(err.Error())
		}
		time.Sleep(2)
		fmt.Print(string(bytes[:]))
	}

}
