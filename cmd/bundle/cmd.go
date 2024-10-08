package bundle

import (
	"log/slog"
	"path"

	"github.com/Hayao0819/seira/bundler"
	"github.com/cockroachdb/errors"
	"github.com/spf13/cobra"
)

func withInternal(_ *cobra.Command, input string, output string, minify bool) error {
	baseDir := path.Dir(input)
	slog.Info("options", "input", input, "output", output, "base", baseDir, "minify", minify)

	err := bundler.Bundle(
		bundler.InputFile(input),
		bundler.OutputFile(output),
		bundler.Base(baseDir),
		bundler.Minify(minify),
		bundler.WorkDir("work"),
	)

	if err != nil {
		return errors.Wrap(err, "failed to bundle")
	}
	return err
}

func Cmd() *cobra.Command {
	output := ""
	minify := false

	cmd := cobra.Command{
		Use:  "bundle",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return withInternal(cmd, args[0], output, minify)
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "output.sh", "output file")
	cmd.Flags().BoolVarP(&minify, "minify", "m", false, "minify output")

	return &cmd
}
