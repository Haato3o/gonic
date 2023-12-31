package gonic

import (
	"context"
	"github.com/Haato3o/gonic/internal/driver"
	"github.com/Haato3o/gonic/internal/search"
)

type SearchOptions struct {
	Limit  *int
	Offset *int
	Locale *Lang
}

type SuggestOptions struct {
	Limit *int
}

type ListOptions struct {
	Limit  *int
	Offset *int
}

type Searchable interface {
	QueryWithContext(ctx context.Context, collection, bucket, query string, opts ...func(options *SearchOptions)) ([]string, error)
	SuggestWithContext(ctx context.Context, collection, bucket, word string, opts ...func(options *SuggestOptions)) ([]string, error)
	ListWithContext(ctx context.Context, collection, bucket string, opts ...func(options *ListOptions)) ([]string, error)
	PingWithContext(ctx context.Context) error
}

func NewSearchableWithContext(ctx context.Context, options ...func(options *Options)) (Searchable, error) {
	opts := &Options{
		Port: 1491,
	}
	for _, fun := range options {
		fun(opts)
	}

	dvr, err := driver.New(opts.Address, opts.Port, driver.ModeSearch)
	if err != nil {
		return nil, err
	}

	if err = dvr.ConnectWithContext(ctx, opts.Password); err != nil {
		return nil, err
	}

	return &search.Engine{
		Driver: dvr,
	}, nil
}
