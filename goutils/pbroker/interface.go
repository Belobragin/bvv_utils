package pbroker

import (
	"context"

	"github.com/belobragin/bvv_utils/goutils/mistake"
)

type BrokerI interface {
	Subscribe(uri, name, user, psw string) error
	Unsubscribe() error
	NewEventListener() EventListenerI
}

type MessageProcessFoo func(context.Context, EventMessageI) mistake.OutErr

type EventListenerI interface {
	ListenEvent() error
}

type HandleEventI interface {
	ProcessEvent(EventMessageI) (MessageProcessFoo, error)
}
