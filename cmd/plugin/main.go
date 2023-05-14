package main

import (
	"github.com/sphr2k/kubectl-unused-volumes/cmd/plugin/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth/exec"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {
	cli.InitAndExecute()
}
