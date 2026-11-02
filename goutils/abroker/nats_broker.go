package abroker

import (
	"time"

	"github.com/belobragin/bvv_utils/goutils/mistake"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

const publishFlushTimeout = 3 * time.Second

type NatsMsg struct {
	msg *nats.Msg
}

func (nm *NatsMsg) GetData() []byte {
	return nm.msg.Data
}
func (nm *NatsMsg) GetKey() string {
	return nm.msg.Reply
}
func (nm *NatsMsg) GetSubject() string {
	return nm.msg.Subject
}

func NewNatsMsg(m *nats.Msg) (*NatsMsg, error) {
	nm := new(NatsMsg)
	if m != nil {
		nm.msg = m
	} else {
		nm.msg = new(nats.Msg)
	}
	return nm, nil
}

type NatsBroker struct {
	conn           *nats.Conn
	listenSubject  string
	publishSubject string
	logger         *zap.Logger
	// InputMsgFoo process input messages i.e. subscriptions
	InputMsgFoo nats.MsgHandler //func(msg *nats.Msg)
}

func (n *NatsBroker) GetConn() *nats.Conn {
	return n.conn
}
func (n *NatsBroker) GetPublishSubject() string {
	return n.publishSubject
}
func (n *NatsBroker) PublishEvent(data []byte) error {
	conn := n.GetConn()
	if err := conn.Publish(n.GetPublishSubject(), data); err != nil {
		return err
	}
	return conn.FlushTimeout(publishFlushTimeout)
}
func (n *NatsBroker) CloseBroker() error {
	n.conn.Drain()
	return nil
}
func (nb *NatsBroker) ListenEvent(errC chan error, sch chan struct{}) {
	ch := make(chan *nats.Msg, 64)
	sub, err := nb.conn.ChanSubscribe(nb.listenSubject, ch)
	defer sub.Unsubscribe()
	if err != nil {
		errC <- err
		return
	}
	for {
		select {
		case _ = <-sch:
			return
		case m := <-ch:
			nb.InputMsgFoo(m)
		}
	}
}

// initialize nats broker structure
func NewNatsBroker(uri, name, portListen, portMetrics, user, psw string,
	listenSubject, publishSubject string,
	inputMsgFoo nats.MsgHandler,
	w *zap.Logger) (*NatsBroker, error) {
	var (
		nb  = new(NatsBroker)
		err error
	)

	if len(uri) == 0 || len(portListen) == 0 || len(portMetrics) == 0 {
		return nil, mistake.ErrSubscribeData
	}
	if user != "" && psw != "" {
		nb.conn, err = nats.Connect(
			uri, nats.Name(name), nats.UserInfo(user, psw))
		if err != nil {
			return nil, err
		}
	} else {
		nb.conn, err = nats.Connect(uri, nats.Name(name))
		if err != nil {
			return nil, err
		}
	}
	nb.logger = w
	if err != nil {
		return nil, err
	}
	nb.listenSubject = listenSubject
	nb.publishSubject = publishSubject
	nb.InputMsgFoo = inputMsgFoo

	return nb, nil
}
