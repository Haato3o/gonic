package main

import (
	"context"
	"github.com/Haato3o/gonic"
	"testing"
)

func Test_Connection(t *testing.T) {
	searchable, err := gonic.NewSearchableWithContext(context.TODO(), func(options *gonic.Options) {
		options.Address = "localhost"
		options.Password = "testing"
	})

}
