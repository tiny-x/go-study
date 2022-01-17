package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/client"
	"io/ioutil"
	"time"
)

func main() {
	opts, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Print(err.Error())
	}
	stats, err := opts.ContainerStats(context.Background(), "094e2320e1a4", true)

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
