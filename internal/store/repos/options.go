package repos

import "github.com/jackc/pgx/v4"

type Option func(*Options)

type Options struct {
	Tx pgx.Tx
}

func getOptions(opts ...Option) Options {
	options := Options{}
	for _, opt := range opts {
		opt(&options)
	}

	return options
}

func WithTx(tx pgx.Tx) Option {
	return func(o *Options) {
		o.Tx = tx
	}
}
