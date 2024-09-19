package bundler

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

func getOpts(o []Option) *options {
	opts := &options{}
	for _, f := range o {
		f(opts)
	}
	return opts
}
