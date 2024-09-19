package bundler

import (
	"io"
	"log/slog"
)

func Bundle(input io.Reader, iname string, output io.Writer, opts ...Option) (any, error) {
	o := getOpts(opts)
	slog.Debug("options", "minify", o.minify, "base", o.base)

	targetFiles, err := getTargetFileList(input, iname)
	if err != nil {
		return nil, err
	}

	slog.Debug("target files", "files", *targetFiles)

	return nil, nil
}
