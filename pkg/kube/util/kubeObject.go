package util

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type KubeObject struct {
}

func (n *KubeObject) ListAllNamespace(c *kubernetes.Clientset) (*v1.NamespaceList, error) {
	return c.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
}

func (n *KubeObject) ListAllNodes(c *kubernetes.Clientset) (*v1.NodeList, error) {
	return c.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
}
