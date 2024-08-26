package bundler

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/Hayao0819/seira/utils"
	"mvdan.cc/sh/v3/syntax"
)

func Bundle(input io.Reader, output io.Writer, opts ...Option) (any, error) {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}

	imports, err := getImportInfosFromReader(input)
	utils.PrintAsJSON(imports)
	if err != nil {
		return nil, err
	}
	for _, ipt := range *imports {
		srcFile, err := os.Open(path.Join(o.base, ipt.FilePath))
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
	base   string
}

type Option func(*options)

func Minify() Option {
	return func(o *options) {
		o.minify = true
	}
}

func Base(base string) Option {
	return func(o *options) {
		o.base = base
	}
}
