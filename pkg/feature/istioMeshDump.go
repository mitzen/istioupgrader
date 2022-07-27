package feature

import (
	"github.com/mitzen/istioupgrader/pkg/kube/config"
	"github.com/mitzen/istioupgrader/pkg/kube/util"
)

type IstioMeshDump struct {
	namespace string
}

func (i *IstioMeshDump) Execute(namespace string) {

	cfg := config.ClientConfig{}
	restConfig := cfg.NewRestConfig()

	ic := util.IstioClient{}
	ic.New(restConfig, namespace)

	// dump istio routing information //
	// gs, vs, dr
	// how to get conflicting vs
	// tls info ???
}
