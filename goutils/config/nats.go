package config

type BrokerConfigI interface {
	GetNatsUri() string
	GetNatsName() string
	GetNatsUser() string
	GetNatsPassw() string
	ValidateBrokerConfig() error
}

type BrokerConfig struct {
	URI        string `conf:"env:NATS_ADDRESS"`
	Name       string `conf:"env:NATS_CLIENT_NAME"`
	Port       string `conf:"env:NATS_PORT"`
	MetricPort string `conf:"env:NATS_METRIC_PORT"`
	User       string `conf:"env:BROKER_AUTH_NAME"`
	Psw        string `conf:"env:BROKER_PSW"`
}

func (r *BrokerConfig) GetNatsUri() string {
	return r.URI
}
func (r *BrokerConfig) GetNatsPort() string {
	return r.Port
}
func (r *BrokerConfig) GetNatsMetricPort() string {
	return r.Port
}
func (r *BrokerConfig) GetNatsName() string {
	return r.Name
}
func (r *BrokerConfig) GetNatsPsw() string {
	return r.Psw
}
func (r *BrokerConfig) GetNatsUser() string {
	return r.User
}
func (r *BrokerConfig) ValidateBrokerConfig() error {
	return nil
}

type BrokerConfigStaticI interface {
	GetNatsUri() string
	GetNatsName() string
	GetNatsUser() string
	GetNatsPsw() string
	ValidateBrokerConfig() error
	GetNatsPort() string
	GetNatsMetricPort() string
	GetListenSubject() string
	GetPublishSubject() string
}

type BrokerConfigStatic struct {
	URI            string `conf:"env:NATS_ADDRESS"`
	Name           string `conf:"env:NATS_CLIENT_NAME"`
	Port           string `conf:"env:NATS_PORT"`
	MetricPort     string `conf:"env:NATS_METRIC_PORT"`
	User           string `conf:"env:BROKER_AUTH_NAME"`
	Psw            string `conf:"env:BROKER_PSW"`
	ListenSubject  string `conf:"env:NATS_FROM_ALTERED_CHAN"`
	PublishSubject string `conf:"env:NATS_FOR_PROCESSING_CHAN"`
}

func (r *BrokerConfigStatic) GetNatsUri() string {
	return r.URI
}
func (r *BrokerConfigStatic) GetNatsPort() string {
	return r.Port
}
func (r *BrokerConfigStatic) GetNatsMetricPort() string {
	return r.MetricPort
}
func (r *BrokerConfigStatic) GetNatsName() string {
	return r.Name
}
func (r *BrokerConfigStatic) GetNatsPsw() string {
	return r.Psw
}
func (r *BrokerConfigStatic) GetNatsUser() string {
	return r.User
}
func (r *BrokerConfigStatic) ValidateBrokerConfig() error {
	return nil
}
func (r *BrokerConfigStatic) GetListenSubject() string {
	return r.ListenSubject
}
func (r *BrokerConfigStatic) GetPublishSubject() string {
	return r.PublishSubject
}

type NatsConfig struct {
	Nats struct {
		URI  string `conf:"env:NATS_ADDRESS"`
		Name string `conf:"env:NATS_CLIENT_NAME"`
		User string `conf:"env:NATS_USER"`
		Psw  string `conf:"env:NATS_PSW"`
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
