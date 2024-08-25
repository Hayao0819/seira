package bundler

import (
	"fmt"
	"io"
	"os"

	"mvdan.cc/sh/v3/syntax"
)

func Bundle(input io.Reader, output io.Writer, opts ...Option) (any, error) {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}

	srcs := getImportedFilePathsFromReader(input)
	for _, src := range *srcs {
		srcFile, err := os.Open(src)
		if err != nil {
			return nil, err
		}
		defer srcFile.Close()

		walkReader(srcFile, func(s *syntax.Stmt) bool {
			funcs := getDefinedFuncList(s)
			fmt.Println(funcs)
			return true
		})
	}

	return nil, nil
}

type options struct {
	minify bool
}

type Option func(*options)

func Minify() Option {
	return func(o *options) {
		o.minify = true
	}
}
