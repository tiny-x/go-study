package main

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-operator/channel"
	"github.com/chaosblade-io/chaosblade-operator/exec/model"
	"github.com/chaosblade-io/chaosblade-operator/pkg/apis/chaosblade/v1alpha1"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
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
		data := []byte(`{"metadata":{"finalizers":[]}}`)
		if err != nil {
			fmt.Println(err.Error())
		}
		obj := &v1alpha1.ChaosBlade{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "chaosblade.io/v1alpha1",
				Kind:       "ChaosBlade",
			},
			ObjectMeta: metav1.ObjectMeta{Name: item.Name},
		}
		patch := client.RawPatch(types.MergePatchType, data)

		err = k8sClient.Patch(context.TODO(), obj, patch)
		if err != nil {
			fmt.Println(err.Error())
		}

		patch = client.MergeFrom(&v1alpha1.ChaosBlade{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "chaosblade.io/v1alpha1",
				Kind:       "ChaosBlade",
			},
			ObjectMeta: metav1.ObjectMeta{
				Finalizers: []string{""},
			},
		})
	}

}

func execCommand(expStatuses []v1alpha1.ExperimentStatus, c channel.Client) {
	for _, expStatuses := range expStatuses {
		for _, status := range expStatuses.ResStatuses {
			if !status.Success {
				// does not need to destroy
				continue
			}
			containerObjectMeta := model.ParseIdentifier(status.Identifier)
			response := c.Exec(&channel.ExecOptions{
				StreamOptions: channel.StreamOptions{
					ErrDecoder: func(bytes []byte) interface{} {
						content := string(bytes)
						return fmt.Errorf(content)
						//return spec.Decode(content, spec.ReturnFail(spec.Code[spec.K8sInvokeError], content))
					},
					OutDecoder: func(bytes []byte) interface{} {
						//content := string(bytes)
						//return spec.Decode(content, spec.ReturnFail(spec.Code[spec.K8sInvokeError], content))
						return nil
					},
				},
				PodName:       containerObjectMeta.PodName,
				PodNamespace:  containerObjectMeta.Namespace,
				ContainerName: containerObjectMeta.ContainerName,
				Command:       []string{"/opt/chaosblade/blade", "status", ""},
				IgnoreOutput:  false,
			}).(*spec.Response)

			fmt.Println(response.Result)
		}
	}
}

func newClient(kubeConfig string) (client.Client, error) {
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
