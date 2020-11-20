package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func main() {
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.24"))
	ctx := context.Background()
	id, _ := cli.ContainerExecCreate(ctx, "86f4dfea83583f4d900b74817aeba34996ea882b36a8cc8b1de0a4264ab6ac2e", types.ExecConfig{
		AttachStderr: true,
		AttachStdout: true,
		Cmd:          []string{"sh", "-c", "/opt/chaosblade/bladea destroy network loss --percent=100 --local-port=80 --interface=eth0 --remote-port=90"},
	})
	resp, _ := cli.ContainerExecAttach(ctx, id.ID, types.ExecStartCheck{})

	//text, _ := resp.Reader.ReadString('\n')
	//fmt.Println(text)

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	stdcopy.StdCopy(stdout, stderr, resp.Reader)
	fmt.Print("out: " + stdout.String())
	fmt.Print("error: " + stderr.String())

}
