package feature

import (
	"fmt"

	"github.com/hashicorp/go-version"
	"github.com/mitzen/istioupgrader/pkg/kube/config"
	"github.com/mitzen/istioupgrader/pkg/kube/util"
	apiv1 "k8s.io/api/core/v1"
)

type IstioUpgrade struct {
}

// gateway + vs => no downtime migration
// go to all namespace and check version again after drain operation

func (i *IstioUpgrade) Execute() {

	var dryRun bool = true

	cfg := config.ClientConfig{}
	restConfig := cfg.NewRestConfig()
	clientset := cfg.NewClientSet(restConfig)

	ic := util.IstioClient{}
	ic.New(restConfig, apiv1.NamespaceAll)

	istioControlVersion := ic.GetIstioControlVersion()
	istiodVersion, err := version.NewVersion(istioControlVersion)

	if err != nil || istioControlVersion == "" {
		fmt.Printf("Unable to get istiod version from istio-system \n")
	}

	nsutil := util.KubeObject{}
	nsutil.NewKubeObject(clientset)

	namespaces, nserr := nsutil.ListAllNamespace()
	if nserr != nil {
		panic("Unable to get namespace(s) from kubernetes")
	} else {

		var isRestartPodRequired bool = false

		for _, n := range namespaces.Items {

			fmt.Println(n.Name)

			if n.Name == config.IstioSystem || n.Name == config.KubeSystem {
				continue
			}

			istioPodVersion := ic.GetIstioPod(n.Name)

			if istioPodVersion != "" {

				fmt.Printf("Istiond version: %s, IstioPod version:%s ", istioControlVersion, istioPodVersion)

				podIstioVersion, err := version.NewVersion(istioPodVersion)

				if err != nil {
					fmt.Printf("Unable to istio version from pods")
				} else {

					if !istiodVersion.Equal(podIstioVersion) {
						fmt.Printf("Istio control plane and data plan version mismatch in namespace: %s", n.Name)
						isRestartPodRequired = true
					} else {
						fmt.Printf("No pods restart is required for namespace: %s", n.Name)
					}
				}
			}
		}

		if !dryRun {

			if isRestartPodRequired {

				nodeLists, err := nsutil.ListAllNodes()

				if err != nil {
					fmt.Println("Error listing node.")
				}

				nc := &util.NodeClient{}
				nc.NewNodeClient(clientset)

				for _, v := range nodeLists.Items {

					nc.Cordon(&v)
					nc.DrainNode(&v)
					nc.UnCordon(&v)
				}
			}
		}
	}
}
