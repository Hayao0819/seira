package bundler

import (
	"io"
	"log/slog"
	"os"
	"path"
	"path/filepath"
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
		slog.Debug("minify", "enable", enable)
		return nil
	}
}

func Base(base string) Option {

	return func(o *options) error {
		var err error
		o.base, err = filepath.Abs(base)
		if err != nil {
			return err
		}

		slog.Debug("base", "dir", o.base)
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
		slog.Debug("input file", "path", path)
		return nil
	}
}

func OutputFile(f string) Option {
	return func(o *options) error {
		if err := os.MkdirAll(path.Dir(f), 0755); err != nil {
			return err
		}
		
		f, err := os.Create(f)
		if err != nil {
			return err
		}
		o.out = f
		o.deferFn = append(o.deferFn, func() {
			f.Close()
		})
		slog.Debug("output file", "path", f)
		return nil
	}
}

func WorkDir(path string) Option {
	return func(o *options) error {
		o.work = path
		slog.Debug("work dir", "path", path)
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
