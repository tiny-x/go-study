package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_get_container(t *testing.T) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Print(err.Error())
	}

	if list, err := cli.ContainerList(context.TODO(), types.ContainerListOptions{
		Filters: filters.NewArgs(
			filters.Arg("name", "nginx"),
		),
	}); err != nil {
		fmt.Print(err.Error())
	} else {
		container := list[0]
		fmt.Printf("容器ID：%s \n", container.ID)
	}
}

func Test_copy(t *testing.T) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Print(err.Error())
	}

	if list, err := cli.ContainerList(context.TODO(), types.ContainerListOptions{
		Filters: filters.NewArgs(
			filters.Arg("name", "nginx"),
		),
	}); err != nil {
		fmt.Print(err.Error())
	} else {
		container := list[0]
		fmt.Printf("容器ID：%s \n", container.ID)

		buffer := new(bytes.Buffer)

		writer := tar.NewWriter(buffer)
		defer writer.Close()

		if err := ArchiveTar("/opt/chaosblade-1.3.1", writer); err != nil {
			fmt.Print(err)
		}

		if err := cli.CopyToContainer(context.TODO(), container.ID, "/opt/", buffer, types.CopyToContainerOptions{
			AllowOverwriteDirWithFile: true,
			CopyUIDGID:                true,
		}); err != nil {
			fmt.Print(err)
		}
	}
}

func Test_exec(t *testing.T) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Print(err.Error())
	}

	if list, err := cli.ContainerList(context.TODO(), types.ContainerListOptions{
		Filters: filters.NewArgs(
			filters.Arg("name", "nginx"),
		),
	}); err != nil {
		fmt.Print(err.Error())
	} else {
		container := list[0]
		fmt.Printf("容器ID：%s \n", container.ID)

		ctx := context.Background()
		id, err := cli.ContainerExecCreate(ctx, container.ID, types.ExecConfig{
			AttachStderr: true,
			AttachStdout: true,
			Cmd:          []string{"sh", "-c", "ls -l /opt/chaosblade-1.3.1/blade"},
			Privileged:   true,
			User:         "root",
		})
		if err != nil {
			fmt.Print(err.Error())
		}
		resp, err := cli.ContainerExecAttach(ctx, id.ID, types.ExecStartCheck{})
		if err != nil {
			fmt.Print(err.Error())
		}
		defer resp.Close()
		stdout := new(bytes.Buffer)
		stderr := new(bytes.Buffer)
		_, err = stdcopy.StdCopy(stdout, stderr, resp.Reader)
		if err != nil {
			fmt.Print(err.Error())
		}
		result := stdout.String()
		errorMsg := stderr.String()

		fmt.Printf("success: %s, \n", result)
		fmt.Printf("error: %s, \n", errorMsg)
	}
}

func Test_get_blade_version(t *testing.T) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Print(err.Error())
	}

	if list, err := cli.ContainerList(context.TODO(), types.ContainerListOptions{
		Filters: filters.NewArgs(
			filters.Arg("name", "nginx"),
		),
	}); err != nil {
		fmt.Print(err.Error())
	} else {
		container := list[0]
		fmt.Printf("容器ID：%s \n", container.ID)

		ctx := context.Background()
		id, err := cli.ContainerExecCreate(ctx, container.ID, types.ExecConfig{
			AttachStderr: true,
			AttachStdout: true,
			Cmd:          []string{"sh", "-c", "/opt/chaosblade-1.3.1/blade -v"},
			Privileged:   true,
			User:         "root",
		})
		if err != nil {
			fmt.Print(err.Error())
		}
		resp, err := cli.ContainerExecAttach(ctx, id.ID, types.ExecStartCheck{})
		if err != nil {
			fmt.Print(err.Error())
		}
		defer resp.Close()
		stdout := new(bytes.Buffer)
		stderr := new(bytes.Buffer)
		_, err = stdcopy.StdCopy(stdout, stderr, resp.Reader)
		if err != nil {
			fmt.Print(err.Error())
		}
		result := stdout.String()
		errorMsg := stderr.String()

		fmt.Printf("success: %s, \n", result)
		fmt.Printf("error: %s, \n", errorMsg)
	}
}

func ArchiveTar(file string, writer *tar.Writer) error {

	return filepath.Walk(file, func(path string, fileInfo fs.FileInfo, err error) error {
		if fileInfo == nil {
			return err
		}
		if fileInfo.IsDir() {
			if path == path {
				return nil
			}
			header, err := tar.FileInfoHeader(fileInfo, "")
			if err != nil {
				return err
			}
			header.Name = filepath.Join(path, strings.TrimPrefix(path, path))
			if err = writer.WriteHeader(header); err != nil {
				return err
			}
			os.Mkdir(strings.TrimPrefix(path, fileInfo.Name()), os.ModeDir)
			return ArchiveTar(path, writer)
		}
		return func(originFile, path string, fileInfo fs.FileInfo, writer *tar.Writer) error {
			if file, err := os.Open(path); err != nil {
				return err
			} else {

				if header, err := tar.FileInfoHeader(fileInfo, ""); err != nil {
					return err
				} else {

					index := strings.LastIndex(originFile, "/")
					header.Name = strings.ReplaceAll(path, originFile[0:index+1], "")

					if err := writer.WriteHeader(header); err != nil {
						return err
					}

					if _, err = io.Copy(writer, file); err != nil {
						return err
					}
				}

			}
			return nil
		}(file, path, fileInfo, writer)
	})
}
