package bundler

import (
	"log/slog"
	"os"

	"github.com/samber/lo"
)

func Bundle(opts ...Option) error {
	o, err := getOpts(opts)
	if err != nil {
		return err
	}
	defer lo.ForEach(o.deferFn, func(fn func(), i int) {
		fn()
	})
	slog.Debug("options", "minify", o.minify, "base", o.base)

	// Get the list of target files
	targetFiles, err := getTargetFileList(o.input, o.name)
	if err != nil {
		return err
	}
	slog.Debug("target files", "files", *targetFiles)

	// Create the work directory
	if err := os.MkdirAll(o.work, 0755); err != nil {
		return err
	}

	// Copy the target files to the work directory
	// TODO: Implement this

	// Create tarball
	// TODO: Implement this

	// Minify the files
	// TODO: Implement this

	// Create the output file
	// TODO: Implement this

	return nil
}
