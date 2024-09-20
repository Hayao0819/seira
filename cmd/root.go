package cmd

import (
	"log/slog"

	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/m-mizutani/clog"
	"github.com/spf13/cobra"
)

var cmdReg = cobrautils.Registory{}

func rootCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:           "seira",
		Short:         "A shell parser",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmdReg.Bind(&cmd)
	return &cmd
}

func init() {
	handler := clog.New(clog.WithColor(true), clog.WithLevel(slog.LevelDebug))
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
