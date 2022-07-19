package util

import (
	"fmt"
	"log"

	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/rest"
)

type IstioClient struct {
	istioClient *versionedclient.Clientset
	ns          string
}

func (i *IstioClient) New(config *rest.Config, ns string) {

	ic, err := versionedclient.NewForConfig(config)

	if err != nil {
		log.Fatalf("Failed to create istio client: %s", err)
	}

	i.istioClient = ic
	i.ns = ns
}

func (i *IstioClient) GetVersionInfo() {

	fmt.Printf("Getting version: %s \n", i.istioClient.NetworkingV1alpha3().RESTClient().APIVersion().Version)

	sv, err := i.istioClient.ServerVersion()

	if err != nil {
		fmt.Println("error getting istio ServerVersion")
	}

	fmt.Printf("Istio version: %s.%s \n", sv.Major, sv.Minor)
}

func (i *IstioClient) Expore() {

}

func (i *IstioClient) GetGateways() {

}

func (i *IstioClient) GetVirtualServices() {

}

func (i *IstioClient) GetDesinationRules() {

}
