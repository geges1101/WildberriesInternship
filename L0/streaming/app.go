package streaming

import "github.com/geges1101/l0/model"

type Server interface {
	Close()
	Publish(data model.Data) error
	Subscribe() (<-chan NewMessage, error)
	OnCreate(f func(NewMessage)) error
}

var serv Server

func SetNats(ob Server) {
	serv = ob
}

func Close() {
	serv.Close()
}

func Publish(data model.Data) error {
	return serv.Publish(data)
}

func Subscribe() (<-chan NewMessage, error) {
	return serv.Subscribe()
}

func OnCreate(f func(message NewMessage)) error {
	return serv.OnCreate(f)
}
