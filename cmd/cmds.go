package cmd

import (
	"github.com/Hayao0819/seira/cmd/ast"
	"github.com/Hayao0819/seira/cmd/bundle"
)

func init() {
	astCmd := ast.Cmd()
	bundleCmd := bundle.Cmd()

	cmdReg.Add(astCmd, bundleCmd)
}
