package main

import (
	"context"
	"github.com/geges1101/l0/model"
	"github.com/geges1101/l0/search"
	"github.com/geges1101/l0/streaming"
	"github.com/geges1101/l0/util"
	"log"
	"net/http"
	"strconv"
)

func onCreated(m streaming.NewMessage) {
	message := model.Data{
		ID:        m.ID,
		Body:      m.Body,
		CreatedAt: m.CreatedAt,
	}
	if err := search.InsertMessage(context.Background(), message); err != nil {
		log.Println(err)
	}
}

func searchMessagesHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()

	query := r.FormValue("query")
	if len(query) == 0 {
		util.ResponseError(w, http.StatusBadRequest, "Missing query parameter")
		return
	}
	next := uint64(0)
	nextStr := r.FormValue("next")
	take := uint64(100)
	takeStr := r.FormValue("take")
	if len(nextStr) != 0 {
		next, err = strconv.ParseUint(nextStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid skip parameter")
			return
		}
	}
	if len(takeStr) != 0 {
		take, err = strconv.ParseUint(takeStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid take parameter")
			return
		}
	}

	messages, err := search.SearchMessage(ctx, query, next, take)
	if err != nil {
		log.Println(err)
		util.ResponseOk(w, []model.Data{})
	}

	util.ResponseOk(w, messages)
}
