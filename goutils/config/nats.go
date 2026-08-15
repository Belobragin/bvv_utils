package config

import (
	"github.com/bvv_utils/goutils/mistake"
	"github.com/nats-io/nats.go"
)

type NatsI interface {
	GetNatsUri() string
	GetNatsName() string
	GetUser() string
	GetPassw() string
}

func ConnectToNats(n NatsI) (*nats.Conn, error) {
	var (
		uri  = n.GetNatsUri()
		name = n.GetNatsName()
	)
	if uri == "" {
		return nil, mistake.ErrNatsUriAbsent
	}
	if name == "" {
		return nil, mistake.ErrNatsClientNameAbsent
	}
	if n.GetUser() != "" && n.GetPassw() != "" {
		return nats.Connect(uri, nats.Name(name), nats.UserInfo(n.GetUser(), n.GetPassw()))
	}
	return nats.Connect(uri, nats.Name(name))
}

type NatsConfig struct {
	Nats struct {
		URI  string `conf:"env:NATS_ADDRESS"`
		Name string `conf:"env:NATS_CLIENT_NAME"`
		User string `default:, conf: "env:NATS_AUTH_NAME"`
		Psw  string `default:, conf:"env:NATS_PSW"`
	}
}

func (r *NatsConfig) GetNatsUri() string {
	return r.Nats.URI
}
func (r *NatsConfig) GetNatsName() string {
	return r.Nats.Name
}
func (r *NatsConfig) GetPassw() string {
	return r.Nats.Psw
}
func (r *NatsConfig) GetUser() string {
	return r.Nats.User
}
