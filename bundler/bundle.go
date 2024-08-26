package bundler

import (
	"errors"
	"io"
	"os"
	"path"

	"github.com/Hayao0819/seira/utils"
	"github.com/samber/lo"
	"mvdan.cc/sh/v3/syntax"
)

func Bundle(input io.Reader, iname string, output io.Writer, opts ...Option) (any, error) {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}

	info, err := getScriptInfo(input, iname)
	if err != nil {
		return nil, err
	}

	if mainFunc := lo.Filter(info.TopLevelFuncs, func(fn syntax.FuncDecl, index int) bool {
		return fn.Name.Value == "main"
	}); len(mainFunc) == 0 {
		return nil, errors.New("main function not found")
	}

	for _, ipt := range info.Imports {
		srcFile, err := os.Open(path.Join(o.base, ipt.FilePath))
		if err != nil {
			return nil, err
		}
		defer srcFile.Close()

		iptScript, err := getScriptInfo(srcFile, ipt.FilePath)
		if err != nil {
			return nil, err
		}

		for _, f := range iptScript.TopLevelFuncs {
			utils.PrintAsJSON(f.Name)
		}
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
