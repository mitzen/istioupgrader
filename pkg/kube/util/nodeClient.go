package util

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubectl/pkg/drain"
)

type NodeClient struct {
	clientset kubernetes.Interface
}

func (nc *NodeClient) NewNodeClient(clientset kubernetes.Interface) {
	nc.clientset = clientset
}

func (nc *NodeClient) Cordon(nodeList *v1.NodeList) {

	helper := drain.Helper{
		Ctx:                             context.TODO(),
		Client:                          nc.clientset,
		Force:                           false,
		GracePeriodSeconds:              0,
		IgnoreAllDaemonSets:             false,
		Timeout:                         0,
		DeleteEmptyDirData:              false,
		Selector:                        "",
		PodSelector:                     "",
		ChunkSize:                       0,
		DisableEviction:                 false,
		SkipWaitForDeleteTimeoutSeconds: 0,
		AdditionalFilters:               []drain.PodFilter{},
		Out:                             nil,
		ErrOut:                          nil,
		DryRunStrategy:                  0,
		//DryRunVerifier:                  &resource.QueryParamVerifier{},
		//OnPodDeletedOrEvicted: func(pod *v1.Pod, usingEviction bool) {
	}

	for _, v := range nodeList.Items {
		drain.RunCordonOrUncordon(&helper, &v, false)
	}
}
