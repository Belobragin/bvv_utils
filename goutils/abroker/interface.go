/*
abroker is general interface and instruments for event-driving architecture
*/
package abroker

type AllMessageI interface {
	GetData() []byte
	GetKey() string
	GetSubject() string
}

type AllBrokerI interface {
	PublishEvent([]byte)
	ListenEvent(chan error, chan struct{})
	CloseBroker() error
}

// type MessageProcessFoo func(context.Context, EventMessageI) mistake.OutErr

// type EventListenerI interface {
// 	ListenEvent() error
// }

// type HandleEventI interface {
// 	ProcessEvent(EventMessageI) (MessageProcessFoo, error)
// }
