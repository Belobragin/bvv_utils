package pbroker

import (
	"context"

	"github.com/belobragin/bvv_utils/goutils/mistake"
)

type BrokerI interface {
	Subscribe(uri, name, user, psw string) error
	Unsubscribe() error
}

type MessageProcessFoo func(context.Context, EventMessageI) mistake.Outerror

type EventListenerI interface {
	ListenEvent() error
}

type EventMessageI interface {
	GetEventMessageKey() *uint8
	GetValue() map[string]interface{}
	// GetValue() interface{}
}

type HandleEventI interface {
	ProcessEvent(EventMessageI) (MessageProcessFoo, error)
}
