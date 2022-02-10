package main

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade/exec/kubernetes"
	"github.com/google/uuid"
	"time"
)

func main() {
	newUUID, _ := uuid.NewUUID()

	uid := newUUID.String()
	response := inject(uid)
	fmt.Print(response)

	time.Sleep(30 * time.Second)

	res := recover(uid)
	fmt.Print(res)
}

func inject(uid string) *spec.Response {
	executor := kubernetes.NewExecutor()
	return executor.Exec(uid, context.Background(), &spec.ExpModel{
		Scope:      "pod",
		Target:     "pod",
		ActionName: "delete",
		ActionFlags: map[string]string{
			"kubeconfig":   "/Users/xf.yefei/.kube/config",
			"namespace":    "default",
			"names":        "chaosblade-box-56c4875976-67tl4",
			"waiting-time": "60s",
		},
	})
}

func recover(suid string) *spec.Response {
	c := context.WithValue(context.Background(), "suid", suid)
	executor := kubernetes.NewExecutor()
	return executor.Exec("", c, &spec.ExpModel{
		Scope:      "pod",
		Target:     "pod",
		ActionName: "delete",
		ActionFlags: map[string]string{
			"kubeconfig":   "/Users/xf.yefei/.kube/config",
			"namespace":    "default",
			"names":        "chaosblade-box-56c4875976-67tl4",
			"waiting-time": "60s",
		},
	})
}
