package bundler

import (
	"io"
	"os"
)

type options struct {
	minify  bool
	base    string
	name    string
	input   io.Reader
	out     io.Writer
	work    string
	deferFn []func()
}

type Option func(*options) error

func Minify(enable bool) Option {
	return func(o *options) error {
		o.minify = enable
		return nil
	}
}

func Base(base string) Option {
	return func(o *options) error {
		o.base = base
		return nil
	}
}

func InputFile(path string) Option {
	return func(o *options) error {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		o.name = path
		o.input = f
		o.deferFn = append(o.deferFn, func() {
			f.Close()
		})
		return nil
	}
}

func OutputFile(path string) Option {
	return func(o *options) error {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		o.out = f
		o.deferFn = append(o.deferFn, func() {
			f.Close()
		})
		return nil
	}
}

func WorkDir(path string) Option {
	return func(o *options) error {
		o.work = path
		return nil
	}
}

func getOpts(o []Option) (*options, error) {
	opts := &options{}
	for _, f := range o {
		if err := f(opts); err != nil {
			return nil, err
		}
	}
	return opts, nil
}
