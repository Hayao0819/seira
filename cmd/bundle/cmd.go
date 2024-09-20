package bundle

import (
	"path"

	"github.com/Hayao0819/seira/bundler"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func withInternal(_ *cobra.Command, input string, output string, minify bool) error {
	baseDir := path.Dir(input)

	err := bundler.Bundle(
		bundler.InputFile(input),
		bundler.OutputFile(output),
		bundler.Base(baseDir),
		bundler.Minify(minify),
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

	cmd.Flags().StringVarP(&output, "output", "o", "", "output file")
	cmd.Flags().BoolVarP(&minify, "minify", "m", false, "minify output")

	return &cmd
}
