package model

import "time"

type Data struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
