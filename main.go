package main

import (
	"fmt"

	"github.com/mitzen/istioupgrader/pkg/kube/config"
	"github.com/mitzen/istioupgrader/pkg/kube/util"
	"istio.io/istio/pkg/kube"
	apiv1 "k8s.io/api/core/v1"
)

func main() {

	cfg := config.ClientConfig{}
	restConfig := cfg.NewConfig()

	clientset := cfg.NewClient(restConfig)

	ic := util.IstioClient{}
	ic.New(restConfig, apiv1.NamespaceAll)

	cc, _ := kube.NewExtendedClient(kube.BuildClientCmd("", ""), "")

	iv, e := ic.GetIstioControlVersion(cc, "istio-system")

	if e != nil {
		fmt.Printf("unable to get version istiod")
	}

	fmt.Printf("%s", iv)

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
