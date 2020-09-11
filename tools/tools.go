// +build tools

package tools

import (
	_ "github.com/go-delve/delve/cmd/dlv"
	_ "github.com/onsi/ginkgo/ginkgo"
	_ "k8s.io/code-generator/cmd/conversion-gen"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
	_ "sigs.k8s.io/kind"
)