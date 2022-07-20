package main

import (
	"fmt"

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
	istioPodVersion := ic.GetIstioPod("default")
	fmt.Printf("Istiond version: %s, IstioPod version:%s ", istioControlVersion, istioPodVersion)

	nsutil := util.KubeNamespace{}
	namespaces, nserr := nsutil.ListAllNamespace(clientset)

	if nserr != nil {
		fmt.Println("")
	}

	for _, n := range namespaces.Items {
		fmt.Println(n.Name)
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	if deploymentsClient != nil {
		fmt.Println("")
	}
}
