package main

import (
	"fmt"
	"github.com/geges1101/l0/streaming"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"log"
	"net/http"
	"time"
)

type Config struct {
	NatsAddress string `envconfig:"NATS_ADDRESS"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	//Connect to Nats
	setup := newSetup()
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		ob, err := streaming.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))
		if err != nil {
			log.Println(err)
			return err
		}

		err = ob.OnCreate(func(message streaming.NewMessage) {
			log.Printf("message received: %v\n", message)
			setup.broadcast(newMessage(message.ID, message.Body, message.CreatedAt), nil)
		})
		if err != nil {
			log.Println(err)
			return err
		}

		streaming.SetNats(ob)
		return nil
	})
	defer streaming.Close()

	//run server
	go setup.run()
	http.HandleFunc("/sender", setup.handleWebSocket)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
