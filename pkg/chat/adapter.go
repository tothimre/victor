package chat

import (
	"fmt"
	"github.com/brettbuddin/victor/pkg/store"
	"github.com/gorilla/mux"
)

const (
	MSG_TYPE_TEXT = "text"
)

var adapters = map[string]InitFunc{}

func Register(name string, init InitFunc) {
	adapters[name] = init
}

func Load(name string) (InitFunc, error) {
	a, ok := adapters[name]

	if !ok {
		return nil, fmt.Errorf("unkown adapter: %s", name)
	}

	return a, nil
}

type InitFunc func(Robot) Adapter

type Adapter interface {
	Run()
	Send(string, string, string)
	Stop()
}

type Robot interface {
	Name() string
	HTTP() *mux.Router
	Store() store.Adapter
	Chat() Adapter
	Receive(Message)
}

type Message interface {
	UserID() string
	UserName() string
	ChannelID() string
	ChannelName() string
	Text() string
}
