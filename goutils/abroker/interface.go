/*
abroker is general interface and instruments for event-driving architecture
*/
package abroker

import "github.com/nats-io/nats.go"

type AllMessageI interface {
	GetData() []byte
	GetKey() string
	GetSubject() string
}

type AllBrokerI interface {
	GetConn() *nats.Conn
	GetPublishSubject() string
	PublishEvent([]byte) error
	ListenEvent(chan error, chan struct{})
	CloseBroker() error
}
