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
	GetPublishSubject() string
	PublishEvent([]byte) error
	ListenEvent(chan error, chan struct{})
	CloseBroker() error
}
