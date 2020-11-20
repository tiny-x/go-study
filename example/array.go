package main

import (
	"fmt"
	"github.com/chaosblade-io/chaosblade-operator/pkg/apis/chaosblade/v1alpha1"
)

func main() {

	results := &v1alpha1.ChaosBladeList{}
	for _, item := range results.Items {
		fmt.Print(item)
	}
}
