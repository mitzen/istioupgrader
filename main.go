package main

import (
	"fmt"

	"github.com/mitzen/istioupgrader/pkg/config"
	apiv1 "k8s.io/api/core/v1"
)

func main() {

	// var kubeconfig *string
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()

	// config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	// if err != nil {
	// 	panic(err)
	// }

	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	panic(err)
	// }

	cfg := config.ClientConfig{}
	clientset := cfg.New()

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	if deploymentsClient != nil {
		fmt.Println("ok all good")
	}
}
