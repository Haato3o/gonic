package search

import (
	"context"
	"github.com/Haato3o/gonic"
	"github.com/Haato3o/gonic/internal/driver"
)

type Engine struct {
	*driver.Driver
}

func (e Engine) QueryWithContext(
	ctx context.Context,
	collection,
	bucket,
	query string,
	opts ...func(options *gonic.SearchOptions),
) ([]string, error) {
	// TODO implement me
	panic("implement me")
}

func (e Engine) SuggestWithContext(
	ctx context.Context,
	collection,
	bucket,
	word string,
	opts ...func(options *gonic.SuggestOptions),
) ([]string, error) {
	// TODO implement me
	panic("implement me")
}

func (e Engine) ListWithContext(ctx context.Context, collection, bucket string, opts ...func(options *gonic.ListOptions)) ([]string, error) {
	// TODO implement me
	panic("implement me")
}

func (e Engine) PingWithContext(ctx context.Context) error {
	e.Driver
}
