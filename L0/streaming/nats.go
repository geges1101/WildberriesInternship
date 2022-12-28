package streaming

import (
	"bytes"
	"encoding/gob"
	"github.com/geges1101/l0/model"
	"github.com/nats-io/nats.go"
	"log"
)

type Connection struct {
	nc              *nats.Conn
	NewSubscription *nats.Subscription
	NewChannel      chan NewMessage
}

func NewNats(url string) (*Connection, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &Connection{nc: nc}, nil
}

func (ob *Connection) Subscribe() (<-chan NewMessage, error) {
	m := NewMessage{}
	ob.NewChannel = make(chan NewMessage, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	ob.NewSubscription, err = ob.nc.ChanSubscribe(m.Key(), ch)

	if err != nil {
		return nil, err
	}

	//Parse message
	go func() {
		for {
			select {
			case msg := <-ch:
				if err := ob.readMessage(msg.Data, &m); err != nil {
					log.Fatal(err)
				}
				ob.NewChannel <- m
			}
		}
	}()
	return (<-chan NewMessage)(ob.NewChannel), nil
}

func (ob *Connection) OnCreate(f func(message NewMessage)) (err error) {
	m := NewMessage{}
	ob.NewSubscription, err = ob.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		if err := ob.readMessage(msg.Data, &m); err != nil {
			log.Fatal(err)
		}
		f(m)
	})
	return
}

func (ob *Connection) Close() {
	if ob.nc != nil {
		ob.nc.Close()
	}
	if ob.NewSubscription != nil {
		if err := ob.NewSubscription.Unsubscribe(); err != nil {
			log.Fatal(err)
		}
	}
	close(ob.NewChannel)
}

func (ob *Connection) Publish(data model.Data) error {
	m := NewMessage{data.ID, data.Body, data.CreatedAt}
	msg, err := ob.writeMessage(&m)
	if err != nil {
		return err
	}
	return ob.nc.Publish(m.Key(), msg)
}

func (ob *Connection) writeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (ob *Connection) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}
