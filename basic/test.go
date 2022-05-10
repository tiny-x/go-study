package main

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"path"
)

func main() {

	fmt.Println(path.Join("/", "/a"))
	ctx := context.Background()
	add(ctx)
	fmt.Println(ctx.Value("a"))

	process, err := process.NewProcess(36578)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(process.Cwd())
	fmt.Println(process.Name())
}

func add(ctx context.Context) {
	ctx = context.WithValue(ctx, "a", "b")
	fmt.Println(ctx.Value("a"))
}
