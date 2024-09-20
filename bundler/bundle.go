package bundler

import (
	"log/slog"
	"os"
	"path"

	cp "github.com/otiai10/copy"
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
	lo.ForEach(*targetFiles, func(file string, i int) {
		dir := path.Base(file)
		if err := os.MkdirAll(path.Join(o.work, dir), 0755); err != nil {
			slog.Error("failed to create directory", "dir", dir, "error", err)
			return
		}

		if err := cp.Copy(file, path.Join(o.work, dir)); err != nil {
			slog.Error("failed to copy file", "file", file, "error", err)
		}
	})

	// Create tarball
	// TODO: Implement this

	// Minify the files
	// TODO: Implement this

	// Create the output file
	// TODO: Implement this

	return nil
}
