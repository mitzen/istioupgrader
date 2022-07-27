package main

import (
	"github.com/mitzen/istioupgrader/pkg/feature"
)

func main() {

	istioUpgrader := feature.IstioUpgrade{}
	istioUpgrader.Execute()
}
