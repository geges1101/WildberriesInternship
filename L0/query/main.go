package main

import (
	"fmt"
	"github.com/geges1101/l0/db"
	"github.com/geges1101/l0/search"
	"github.com/geges1101/l0/streaming"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"log"
	"net/http"
	"time"
)

type Config struct {
	PostgresDB           string `envconfig:"POSTGRES_DB"`
	PostgresUser         string `envconfig:"POSTGRES_USER"`
	PostgresPassword     string `envconfig:"POSTGRES_PASSWORD"`
	NatsAddress          string `envconfig:"NATS_ADDRESS"`
	ElasticsearchAddress string `envconfig:"ELASTICSEARCH_ADDRESS"`
}

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/search", searchMessagesHandler).
		Methods(http.MethodGet)
	router.Use(mux.CORSMethodMiddleware(router))
	return
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to PostgreSQL
	retry.ForeverSleep(2*time.Second, func(attempt int) error {
		addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
		repo, err := db.NewPostgres(addr)
		if err != nil {
			log.Println(err)
			return err
		}
		db.SetRepository(repo)
		return nil
	})
	defer db.Close()

	//Connect to Elastic
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		es, err := search.NewElastic(fmt.Sprintf("http://%s", cfg.ElasticsearchAddress))
		if err != nil {
			log.Println(err)
			return err
		}
		search.SetRepository(es)
		return nil
	})
	defer search.Close()

	// Connect to Nats
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		es, err := streaming.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))
		if err != nil {
			log.Println(err)
			return err
		}
		err = es.OnCreate(onCreated)
		if err != nil {
			log.Println(err)
			return err
		}
		streaming.SetNats(es)
		return nil
	})
	defer streaming.Close()

	// Start http server
	router := newRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
