package main

import (
	"github.com/geges1101/l0/db"
	"github.com/geges1101/l0/model"
	"github.com/geges1101/l0/streaming"
	"github.com/geges1101/l0/util"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
)

func createMessageHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		ID string `json:"id"`
	}

	ctx := r.Context()

	// Read parameters
	body := template.HTMLEscapeString(r.FormValue("body"))
	if len(body) < 1 || len(body) > 140 {
		util.ResponseError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	// Create message
	createdAt := time.Now().UTC()
	id, err := ksuid.NewRandomWithTime(createdAt)
	if err != nil {
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create meow")
		return
	}
	data := model.Data{
		ID:        id.String(),
		Body:      body,
		CreatedAt: createdAt,
	}
	if err := db.InsertMessage(ctx, data); err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create meow")
		return
	}

	// Publish event
	if err := streaming.Publish(data); err != nil {
		log.Println(err)
	}

	// Return new data
	util.ResponseOk(w, response{ID: data.ID})
}
