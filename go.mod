module github.com/tiny-x/go-study

go 1.14

require (
	github.com/chaosblade-io/chaosblade v1.5.0
	github.com/chaosblade-io/chaosblade-exec-cri v1.5.0 // indirect
	github.com/chaosblade-io/chaosblade-exec-docker v1.5.0 // indirect
	github.com/containerd/cgroups v1.0.2-0.20210605143700-23b51209bf7b
	github.com/docker/docker v20.10.12+incompatible
	k8s.io/apimachinery v0.20.6
    k8s.io/client-go v12.0.0+incompatible
    sigs.k8s.io/controller-runtime v0.6.0
)

replace k8s.io/client-go => k8s.io/client-go v0.20.6