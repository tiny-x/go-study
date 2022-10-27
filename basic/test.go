package main

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"path"
	"strconv"
	"time"
)

func main() {
	ax(nil)

	atoi, _ := strconv.Atoi("1")
	fmt.Println(atoi * -1)
	m, _ := time.ParseDuration(fmt.Sprintf("%sms", "-1000"))
	fmt.Println(m)

	fmt.Println(path.Join("/", "/a"))
	ctx := context.Background()
	add(&ctx)
	fmt.Println(ctx.Value("a"))

	process, err := process.NewProcess(36578)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(process.Cwd())
	fmt.Println(process.Name())
}

func add(ctx *context.Context) {
	*ctx = context.WithValue(*ctx, "a", "b")
	fmt.Println((*ctx).Value("a"))
	fmt.Println("------------")
}

func ax(a interface{}) {
	s, ok := a.(string)
	fmt.Println(s)
	fmt.Println(ok)
}
