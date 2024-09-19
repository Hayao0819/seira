package bundle

import (
	"os"
	"path"

	"github.com/Hayao0819/seira/bundler"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func withInternal(_ *cobra.Command, input string, output string, _ bool) error {
	var (
		ifile *os.File
		ofile *os.File
		err   error
	)

	ifile, err = os.Open(input)
	if err != nil {
		return err
	}
	defer ifile.Close()

	if output != "" {
		ofile, err = os.Create(output)
		if err != nil {
			return err
		}
	}

	_, err = bundler.Bundle(ifile, input, ofile, bundler.Base(path.Dir(input)))
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
