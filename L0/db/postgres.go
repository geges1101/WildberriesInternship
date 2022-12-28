package db

import (
	"context"
	"database/sql"
	"github.com/geges1101/l0/model"
	"log"

	_ "github.com/lib/pq"
)

type myPSQL struct {
	db *sql.DB
}

func NewPostgres(url string) (*myPSQL, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &myPSQL{
		db,
	}, nil
}

func (r *myPSQL) Close() {
	if err := r.db.Close(); err != nil {
		log.Fatal(err)
	}
}

func (r *myPSQL) InsertMessage(ctx context.Context, data model.Data) error {
	_, err := r.db.Exec("INSERT INTO meows(id, body, created_at) VALUES($1, $2, $3)", data.ID, data.Body, data.CreatedAt)
	return err
}
