package util

import (
	"context"
	"log"

	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	"istio.io/istio/pkg/kube"
	"istio.io/pkg/version"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func (i *IstioClient) GetIstioControlVersion() (*version.MeshInfo, error) {
	cc, _ := kube.NewExtendedClient(kube.BuildClientCmd("", ""), "")
	mvi, err := cc.GetIstioVersions(context.TODO(), "istio-system")
	return mvi, err
}

func (i *IstioClient) GetGateways() (*v1alpha3.GatewayList, error) {
	return i.istioClient.NetworkingV1alpha3().Gateways(i.ns).List(context.TODO(), v1.ListOptions{})
}

func (i *IstioClient) GetVirtualServices() {
}

func (i *IstioClient) GetDesinationRules() {
}
