package bundler

import (
	"fmt"
	"io"

	"github.com/samber/lo"
)

func Bundle(input io.Reader, output io.Writer, opts ...Option) (any, error) {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}

	srcs := getImportedFilePaths(input)
	if srcs != nil {
		lo.ForEach(*srcs, func(src string, index int) {
			fmt.Println(src)
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
