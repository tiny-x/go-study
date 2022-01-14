package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"os"
	"testing"
)

func Test_run_container(t *testing.T) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "xchaos-tool:0.0.1",
		Cmd:   []string{"/bin/sh", "-c", "/opt/chaosblade/blade version"},
		Labels: map[string]string{
			"xchaos": "xchaos-sidecar",
		},
	}, &container.HostConfig{
		NetworkMode: container.NetworkMode(fmt.Sprintf("container:%s", "79a276832ed8")),
		CapAdd:      []string{"NET_ADMIN"},
		Binds:       []string{"/sys/fs/cgroup:/sys/fs/cgroup"},
	}, nil, nil, fmt.Sprintf("xchaos-sidecar-%s", "79a276832ed8"))

	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	//cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})
}
