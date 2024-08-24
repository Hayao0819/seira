package cmd

import (
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

var cmdReg = cobrautils.Registory{}

func rootCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "seira",
		Short: "A shell parser",
	}

	cmdReg.Bind(&cmd)
	return &cmd
}
