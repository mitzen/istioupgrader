package main

import (
	"fmt"

	"github.com/hashicorp/go-version"
	"github.com/mitzen/istioupgrader/pkg/kube/config"
	"github.com/mitzen/istioupgrader/pkg/kube/util"
	apiv1 "k8s.io/api/core/v1"
)

func main() {

	cfg := config.ClientConfig{}
	restConfig := cfg.NewConfig()
	clientset := cfg.NewClient(restConfig)

	ic := util.IstioClient{}
	ic.New(restConfig, apiv1.NamespaceAll)

	istioControlVersion := ic.GetIstioControlVersion()
	istiodVersion, err := version.NewVersion(istioControlVersion)

	if err != nil || istioControlVersion == "" {
		fmt.Printf("Unable to get istiod version from istio-system")
	}

	nsutil := util.KubeNamespace{}
	namespaces, nserr := nsutil.ListAllNamespace(clientset)

	if nserr != nil {
		fmt.Println("")
	}

	for _, n := range namespaces.Items {

		fmt.Println(n.Name)

		if n.Name == "istio-system" || n.Name == "kube-system" {
			continue
		}

		istioPodVersion := ic.GetIstioPod(n.Name)

		if istioPodVersion != "" {

			fmt.Printf("Istiond version: %s, IstioPod version:%s ", istioControlVersion, istioPodVersion)

			podIstioVersion, err := version.NewVersion(istioPodVersion)

			var isRestartPodRequired bool = false

			if err != nil {
				fmt.Printf("Unable to istio version from pods")
			} else {

				if !istiodVersion.Equal(podIstioVersion) {
					fmt.Printf("We need to restart pods in namespace: %s", n.Name)
					isRestartPodRequired = true
				} else {
					fmt.Printf("No pods restart is required for namespace: %s", n.Name)
				}

				if isRestartPodRequired {

					// restart pod in a namespace //

				}
			}
		}
	}

	// deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	// if deploymentsClient != nil {
	// 	fmt.Println("")
	// }
}
