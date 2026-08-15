package pbroker

import (
	"github.com/belobragin/bvv_utils/goutils/mistake"
	"github.com/nats-io/nats.go"
)

type NatsBroker struct {
	conn       *nats.Conn
	listenChan chan *nats.Msg
	subscriber *nats.Subscription
}

func (p *NatsBroker) Unsubscribe() error {
	return p.subscriber.Unsubscribe()
}

func (p *NatsBroker) GetNatsConn() *nats.Conn {
	return p.conn
}

func (p *NatsBroker) GetListenChan() chan *nats.Msg {
	return p.listenChan
}

func (p *NatsBroker) GetNatsSubscriber() *nats.Subscription {
	return p.subscriber
}

func (n *NatsBroker) Subscribe(uri, name, user, psw string) error {
	var (
		err error
	)
	n.listenChan = make(chan *nats.Msg, 64)
	if uri == "" {
		return mistake.ErrNatsUriAbsent
	}
	if name == "" {
		return mistake.ErrNatsClientNameAbsent
	}
	if user != "" && psw != "" {
		n.conn, err = nats.Connect(
			uri, nats.Name(name), nats.UserInfo(user, psw))
		if err != nil {
			return nil
		}
	} else {
		n.conn, err = nats.Connect(uri, nats.Name(name))
		if err != nil {
			return err
		}
	}
	n.subscriber, err = n.conn.ChanSubscribe(name, n.listenChan)
	if err != nil {
		return err
	}

	return nil
}
