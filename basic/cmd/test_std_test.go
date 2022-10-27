package cmd

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"testing"
	"time"
)

func Test_std(t *testing.T) {

	ctx := context.Background()
	command := exec.CommandContext(ctx, "/tmp/GoLand/bin/chaos_os", "create", "cpu", "fullload", "--cpu-percent=80")

	var outb bytes.Buffer
	var errb bytes.Buffer
	command.Stdout = &outb
	command.Stderr = &errb
	if err := command.Start(); err != nil {
		fmt.Printf("start err: %v", err)
	}
	errc := make(chan string, 1)

	ctx, cancelFunc := context.WithTimeout(ctx, time.Millisecond*100)

	go func() {
		if err := command.Wait(); err != nil {
			fmt.Printf("wait err: %v, %s", err, errb.String())
			errc <- errb.String()
		}
	}()

	<-ctx.Done()
	cancelFunc()
	select {
	case err := <-errc:
		t.Errorf(err)
		break
	default:
		fmt.Println("success")
	}

}
