package main

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-operator/pkg/apis/chaosblade/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {
	k8sClient, err := newClient2("")
	if err != nil {
		fmt.Print(err.Error())
	}

	results := &v1alpha1.ChaosBladeList{}
	err = k8sClient.List(context.TODO(), results, &client.ListOptions{
		FieldSelector: client.MatchingFieldsSelector{
			Selector: fields.OneTermEqualSelector("status.phase", "Error"),
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	for _, item := range results.Items {
		fmt.Println(fmt.Sprintf("blade: %s , status: %s", item.Name, item.Status.Phase))
	}
}

func get(cli client.Client) {

	result := &v1alpha1.ChaosBlade{}
	err := cli.Get(context.TODO(), types.NamespacedName{Name: "1ac399c2aa79dc0d"}, result)
	result.TypeMeta = metav1.TypeMeta{
		APIVersion: "chaosblade.io/v1alpha1",
		Kind:       "ChaosBlade",
	}

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result.Status.ExpStatuses)
}

func newClient2(kubeConfig string) (client.Client, error) {
	var clusterConfig *rest.Config
	var err error
	if kubeConfig == "" {
		clusterConfig, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	} else {
		clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			&clientcmd.ClientConfigLoadingRules{
				ExplicitPath: kubeConfig,
			},
			&clientcmd.ConfigOverrides{},
		)
		clusterConfig, err = clientConfig.ClientConfig()
	}
	if err != nil {
		return nil, err
	}
	clusterConfig.ContentConfig.GroupVersion = &v1alpha1.SchemeGroupVersion
	clusterConfig.APIPath = "/apis"
	clusterConfig.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: scheme.Codecs}
	clusterConfig.UserAgent = rest.DefaultKubernetesUserAgent()
	scheme, err := v1alpha1.SchemeBuilder.Build()
	if err != nil {
		return nil, err
	}
	return client.New(clusterConfig, client.Options{Scheme: scheme})
}
