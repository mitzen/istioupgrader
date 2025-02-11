package util

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/mitzen/istioupgrader/pkg/kube/config"
	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	"istio.io/istio/pkg/kube"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

const (
	istioSystemNamespace = "istio-system"
)

type IstioClient struct {
	istioClient         *versionedclient.Clientset
	namespace           string
	istioExtendedClient kube.ExtendedClient
}

func (i *IstioClient) New(config *rest.Config, namespace string) {

	ic, err := versionedclient.NewForConfig(config)

	i.initializeIstioClient()

	if err != nil {
		log.Fatalf("Failed to create istio client: %s", err)
	}

	i.istioClient = ic
	i.namespace = namespace
}

func (i *IstioClient) GetIstioControlVersion() string {

	mvi, err := i.istioExtendedClient.GetIstioVersions(context.TODO(), istioSystemNamespace)

	if err != nil {
		fmt.Printf("Unable to get version istiod")
	}

	for _, v := range *mvi {
		if v.Info.Version != "" {
			return v.Info.Version
		}
	}
	return ""
}

func (i *IstioClient) GetIstioPod(namespace string) string {

	mvi, err := i.istioExtendedClient.GetIstioPods(context.TODO(), namespace, map[string]string{})

	if err != nil {
		fmt.Println("error getting pods")
	}

	for _, v := range mvi {

		for _, a := range v.Spec.Containers {
			if strings.Contains(a.Name, config.IstioProxyImage) {
				ss := strings.Split(a.Image, ":")
				istioProxyVersion := ss[len(ss)-1]
				return istioProxyVersion
			}
		}
	}

	return ""
}

func (i *IstioClient) GetGateways() (*v1alpha3.GatewayList, error) {
	return i.istioClient.NetworkingV1alpha3().Gateways(i.namespace).List(context.TODO(), v1.ListOptions{})
}

func (i *IstioClient) GetVirtualServices() {
	i.istioClient.NetworkingV1alpha3().VirtualServices(i.namespace)
}

func (i *IstioClient) GetDesinationRules() {
	i.istioClient.NetworkingV1alpha3().DestinationRules(i.namespace)
}

func (i *IstioClient) initializeIstioClient() {

	client, err := kube.NewExtendedClient(kube.BuildClientCmd("", ""), "")

	if err != nil {
		fmt.Println("unable to create istio extended client")
	}

	i.istioExtendedClient = client
}
