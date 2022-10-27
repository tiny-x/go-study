module github.com/tiny-x/go-study

go 1.14

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/alecthomas/kong v0.5.0
	github.com/containerd/cgroups v1.0.2-0.20210605143700-23b51209bf7b
	github.com/docker/docker v20.10.12+incompatible
	github.com/ethercflow/hookfs v0.3.0
	github.com/google/gopacket v1.1.19
	github.com/google/uuid v1.2.0 // indirect
	github.com/jasonlvhit/gocron v0.0.1
	github.com/ncw/directio v1.0.5
	github.com/onsi/gomega v1.8.1 // indirect
	github.com/prometheus/client_golang v1.7.1
	github.com/prometheus/common v0.10.0
	github.com/safchain/ethtool v0.0.0-20190326074333-42ed695e3de8
	github.com/shirou/gopsutil v3.21.6+incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	go.uber.org/automaxprocs v1.3.0
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	helm.sh/helm/v3 v3.6.2
	k8s.io/client-go v12.0.0+incompatible // indirect
	rsc.io/letsencrypt v0.0.3 // indirect
)

replace k8s.io/client-go => k8s.io/client-go v0.21.1
