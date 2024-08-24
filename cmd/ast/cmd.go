package ast

import (
	"github.com/Hayao0819/seira/utils"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "ast",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			parsed, err := utils.ParseFile(args[0])
			if err != nil {
				return err
			}
			utils.PrintAsJSON(parsed)
			return nil
		},
	}

	return &cmd
}
