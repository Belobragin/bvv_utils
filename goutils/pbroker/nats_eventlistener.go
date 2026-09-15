package pbroker

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/belobragin/bvv_utils/goutils/mistake"
	"github.com/nats-io/nats.go"

	"go.uber.org/zap"
)

type NatsEventListener struct {
	listenNatChan  chan *nats.Msg
	eventProcessor HandleEventI
	logger         *zap.Logger
	stopChan       chan struct{}
}

func (w *NatsEventListener) getEventProcessor() (HandleEventI, error) {
	if eventProcessor := w.eventProcessor; eventProcessor == nil {
		return nil, mistake.ErrMessageTypeUnknown
	} else {
		return eventProcessor, nil
	}
}

func NewNatsEventListener(natsChan chan *nats.Msg, ep HandleEventI, logger *zap.Logger, stopChan chan struct{}) *NatsEventListener {
	return &NatsEventListener{
		listenNatChan:  natsChan,
		eventProcessor: ep,
		logger:         logger,
		stopChan:       stopChan,
	}
}

func (w *NatsEventListener) ListenEvent() error {
	zapstruct := w.logger
	var (
		eventProcessor HandleEventI
		wg             sync.WaitGroup
		err            error
	)
	if eventProcessor, err = w.getEventProcessor(); err != nil {
		zapstruct.Error(err.Error())
		return err
	}
	for {
		select {
		case <-w.stopChan:
			zapstruct.Info("stop on overall stop signal")
			wg.Wait()
			return nil
		case msg := <-w.listenNatChan:
			// m.Value = make(map[string]interface{})
			if msg == nil {
				continue
			}
			var m InputEvent
			err := json.Unmarshal(msg.Data, &m)
			if err != nil {
				zapstruct.Error(fmt.Sprintf(`can not parse data from nats message with subject %s, reply %s,
					header %+v and subscription %+v`,
					msg.Subject, msg.Reply, msg.Header, msg.Sub), zap.Error(err))
				continue
			}
			eventKey := m.Key
			if eventKey == nil {
				zapstruct.Error("nats message key invalid format: ", zap.Error(err))
				continue
			}
			zapstruct.Info(fmt.Sprintf("Received event message %+v from %s with key: %d",
				m, msg.Subject, *eventKey))
			processF, err := eventProcessor.ProcessEvent(&m)
			if err != nil {
				zapstruct.Error(fmt.Sprintf("get process foo error for nats message key %d",
					m.Key), zap.Error(err))
				continue
			}
			wg.Add(1)
			go func(inp *InputEvent, f MessageProcessFoo) {
				defer wg.Done()
				e := f(context.TODO(), inp)
				if e.Err() != nil {
					zapstruct.Error(fmt.Sprintf("processed nats key %d",
						m.Key), zap.Error(e.Err()))
				}
			}(&m, processF)
		}
	}
}
