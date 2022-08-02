module github.com/mitzen/istioupgrader

go 1.16

require (
	github.com/envoyproxy/go-control-plane v0.10.3-0.20220630171008-2ce288a42860
	github.com/fatih/color v1.13.0
	github.com/hashicorp/go-version v1.5.0
	github.com/spf13/cobra v1.5.0
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f
	google.golang.org/grpc v1.48.0 // indirect
	istio.io/client-go v1.14.2
	istio.io/istio v0.0.0-20220719142312-772978915dfe
	istio.io/pkg v0.0.0-20220713144817-9a0434494868 // indirect
	k8s.io/api v0.24.3
	k8s.io/apiextensions-apiserver v0.24.3 // indirect
	k8s.io/apimachinery v0.24.3
	k8s.io/client-go v0.24.3
	k8s.io/kubectl v0.24.3
	sigs.k8s.io/controller-runtime v0.12.2
)
