module github.com/tiny-x/go-study

go 1.14

require (
	github.com/alecthomas/kong v0.2.12
	github.com/chaosblade-io/chaosblade-exec-docker v0.7.0
	github.com/chaosblade-io/chaosblade-exec-os v0.9.0
	github.com/chaosblade-io/chaosblade-operator v0.7.0
	github.com/chaosblade-io/chaosblade-spec-go v0.9.0
	github.com/docker/docker v1.4.2-0.20200203170920-46ec8731fbce
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	github.com/shirou/gopsutil v2.20.5+incompatible
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/cobra v1.0.0
	go.uber.org/automaxprocs v1.3.0
	golang.org/x/crypto v0.0.0-20200220183623-bac4c82f6975
	helm.sh/helm/v3 v3.1.2
	k8s.io/apimachinery v0.17.5
	k8s.io/client-go v12.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.5.3
)

replace k8s.io/client-go => k8s.io/client-go v0.17.5
