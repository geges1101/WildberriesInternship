package db

import (
	"context"
	"github.com/geges1101/l0/model"
)

type PSQL interface {
	Close()
	InsertMessage(ctx context.Context, data model.Data) error
}

var ob PSQL

func SetRepository(repository PSQL) {
	ob = repository
}

func Close() {
	ob.Close()
}

func InsertMessage(ctx context.Context, data model.Data) error {
	return ob.InsertMessage(ctx, data)
}
