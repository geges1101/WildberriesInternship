package search

import (
	"context"
	"github.com/geges1101/l0/model"
)

type Repository interface {
	Close()
	InsertMessage(ctx context.Context, data model.Data) error
	SearchMessage(ctx context.Context, query string, skip uint64, take uint64) ([]model.Data, error)
}

var ob Repository

func SetRepository(repository Repository) {
	ob = repository
}

func Close() {
	ob.Close()
}

func InsertMessage(ctx context.Context, data model.Data) error {
	return ob.InsertMessage(ctx, data)
}

func SearchMessage(ctx context.Context, query string, skip uint64, take uint64) ([]model.Data, error) {
	return ob.SearchMessage(ctx, query, skip, take)
}
