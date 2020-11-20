package main

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-operator/pkg/apis/chaosblade/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {
	k8sClient, err := newClient("/Users/yefei/.kube/config")
	if err != nil {
		fmt.Print(err.Error())
	}

	results := &v1alpha1.ChaosBladeList{}
	k8sClient.List(context.TODO(), results, &client.ListOptions{})
	for _, item := range results.Items {
		fmt.Println(item.Name)
		fmt.Println(item.Status.Phase == v1alpha1.ClusterPhaseDestroying)

		// patch blade
		err := k8sClient.Patch(context.TODO(),
			&v1alpha1.ChaosBlade{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "chaosblade.io/v1alpha1",
					Kind:       "ChaosBlade",
				},
				ObjectMeta: metav1.ObjectMeta{Name: item.Name},
			},
			client.RawPatch(types.MergePatchType, []byte(`{"metadata":{"finalizers":[]}}`)),
		)

		if err != nil {
			fmt.Println(err.Error())
		}
	}

}
