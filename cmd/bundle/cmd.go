package bundle

import (
	"os"
	"path"

	"github.com/Hayao0819/seira/bundler"
	sbb "github.com/malscent/bash_bundler/pkg/bundler"
	"github.com/spf13/cobra"
)

func withSbb(_ *cobra.Command, input string, output string, minify bool) error {
	s, err := sbb.Bundle(input, true)
	if err != nil {
		return err
	}

	if minify {
		s, err = sbb.Minify(s)
		if err != nil {
			return err
		}
	}

	if output != "" {
		err = sbb.WriteToFile(output, s)
		if err != nil {
			return err
		}
	}

	return nil
}

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
	return err
}

func Cmd() *cobra.Command {
	output := ""
	minify := false
	useSbb := false

	cmd := cobra.Command{
		Use:  "bundle",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if useSbb {
				return withSbb(cmd, args[0], output, minify)
			}
			return withInternal(cmd, args[0], output, minify)
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "output file")
	cmd.Flags().BoolVarP(&minify, "minify", "m", false, "minify output")
	cmd.Flags().BoolVarP(&useSbb, "sbb", "s", false, "use sbb")
	return &cmd
}
