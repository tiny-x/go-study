package main

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-operator/channel"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {

	kubeconfig := config.GetConfigOrDie()
	c := channel.Client{
		Interface: kubernetes.NewForConfigOrDie(kubeconfig),
		Client:    nil,
		Config:    kubeconfig,
	}
	//expStatuses := result.Status.ExpStatuses
	//execCommand(expStatuses, c)
	list, err := c.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	} else {
		if len(list.Items) > 0 {
			fmt.Println(list.Items[1].Namespace)
			fmt.Println(list.Items[1].Name)
			fmt.Println(list.Items[1].Labels)
		}
	}

}
