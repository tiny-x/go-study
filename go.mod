module github.com/tiny-x/go-study

go 1.14

require (
	github.com/alecthomas/kong v0.2.12
	github.com/chaosblade-io/chaosblade-exec-docker v0.7.0
	github.com/chaosblade-io/chaosblade-exec-os v0.9.0
	github.com/chaosblade-io/chaosblade-operator v0.7.0
	github.com/chaosblade-io/chaosblade-spec-go v0.9.0
	github.com/containerd/cgroups v0.0.0-20191011165608-5fbad35c2a7e
	github.com/docker/docker v20.10.12+incompatible
	github.com/google/gopacket v1.1.19
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/jasonlvhit/gocron v0.0.1
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/prometheus/client_golang v1.5.1
	github.com/prometheus/common v0.9.1
	github.com/safchain/ethtool v0.0.0-20210803160452-9aa261dae9b1
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	github.com/shirou/gopsutil v3.21.8-0.20210816101416-f86a04298073+incompatible
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/cobra v1.0.0
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	go.uber.org/automaxprocs v1.3.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gotest.tools/v3 v3.1.0 // indirect
	helm.sh/helm/v3 v3.1.2
	k8s.io/apimachinery v0.17.5
	k8s.io/client-go v12.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.5.3
)

replace k8s.io/client-go => k8s.io/client-go v0.17.5
