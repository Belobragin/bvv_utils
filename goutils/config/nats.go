package config

type BrokerConfigI interface {
	GetUri() string
	GetName() string
	GetUser() string
	GetPassw() string
}

type BrokerConfig struct {
	Broker struct {
		URI  string `conf:"env:NATS_ADDRESS"`
		Name string `conf:"env:NATS_CLIENT_NAME"`
		User string `default:,conf:"env:BROKER_AUTH_NAME"`
		Psw  string `default:,conf:"env:BROKER_PSW"`
	}
}

func (r *BrokerConfig) GetUri() string {
	return r.Broker.URI
}
func (r *BrokerConfig) GetName() string {
	return r.Broker.Name
}
func (r *BrokerConfig) GetPassw() string {
	return r.Broker.Psw
}
func (r *BrokerConfig) GetUser() string {
	return r.Broker.User
}
func (r *BrokerConfig) ValidateBrokerConfig() error {
	return nil
}

// func ConnectToNats(n BrokerConfigI) (*nats.Conn, error) {
// 	var (
// 		uri  = n.GetUri()
// 		name = n.GetName()
// 	)
// 	if uri == "" {
// 		return nil, mistake.ErrNatsUriAbsent
// 	}
// 	if name == "" {
// 		return nil, mistake.ErrNatsClientNameAbsent
// 	}
// 	if n.GetUser() != "" && n.GetPassw() != "" {
// 		return nats.Connect(uri, nats.Name(name), nats.UserInfo(n.GetUser(), n.GetPassw()))
// 	}
// 	return nats.Connect(uri, nats.Name(name))
// }

type NatsConfig struct {
	Nats struct {
		URI  string `conf:"env:NATS_ADDRESS"`
		Name string `conf:"env:NATS_CLIENT_NAME"`
		User string `default:,conf: "env:NATS_AUTH_NAME"`
		Psw  string `default:,conf:"env:NATS_PSW"`
	}
}

func (r *NatsConfig) GetUri() string {
	return r.Nats.URI
}
func (r *NatsConfig) GetName() string {
	return r.Nats.Name
}
func (r *NatsConfig) GetPassw() string {
	return r.Nats.Psw
}
func (r *NatsConfig) GetUser() string {
	return r.Nats.User
}
