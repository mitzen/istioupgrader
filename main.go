package main

import (
	"fmt"

	"github.com/mitzen/istioupgrader/pkg/kube/config"
	"github.com/mitzen/istioupgrader/pkg/kube/util"
	apiv1 "k8s.io/api/core/v1"
)

func main() {

	cfg := config.ClientConfig{}
	clientset := cfg.New()

	nsutil := util.KubeNamespace{}

	namespaces, nserr := nsutil.ListAllNamespace(clientset)

	if nserr != nil {
		fmt.Println("")
	}

	fmt.Println("Outputing namespace")

	for _, n := range namespaces.Items {
		fmt.Println(n.Name)
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	if deploymentsClient != nil {
		fmt.Println("ok all good")
	}
}
