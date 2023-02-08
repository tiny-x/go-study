module github.com/tiny-x/go-study

go 1.14

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/alecthomas/kong v0.5.0
	github.com/containerd/cgroups v1.0.4
	github.com/docker/docker v20.10.21+incompatible
	github.com/ethercflow/hookfs v0.3.0
	github.com/google/gopacket v1.1.19
	github.com/jasonlvhit/gocron v0.0.1
	github.com/ncw/directio v1.0.5
	github.com/prometheus/client_golang v1.14.0
	github.com/safchain/ethtool v0.0.0-20210803160452-9aa261dae9b1
	github.com/shirou/gopsutil v3.21.6+incompatible
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.6.1
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	go.uber.org/automaxprocs v1.3.0
	golang.org/x/crypto v0.5.0
	helm.sh/helm/v3 v3.11.1
	k8s.io/client-go v12.0.0+incompatible // indirect
)

replace k8s.io/client-go => k8s.io/client-go v0.21.1
