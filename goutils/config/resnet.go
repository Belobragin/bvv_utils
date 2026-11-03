package config

import "errors"

type ResnetConfigI interface {
	GetResnetPort() string
	GetResnetHost() string
	GetResnetPredictPath() string
	ValidateResnetConfig() error
}

type ResnetConfig struct {
	ResnetPort        string `conf:"env:RESNET_SERVER_PORT"`
	ResnetHost        string `conf:"env:RESNET_SERVER_HOST"`
	ResnetPredictPath string `conf:"env:RESNET_PREDICT_PATH"`
}

func (r *ResnetConfig) GetResnetPort() string {
	return r.ResnetPort
}
func (r *ResnetConfig) GetResnetHost() string {
	return r.ResnetHost
}
func (r *ResnetConfig) GetResnetPredictPath() string {
	return r.ResnetPredictPath
}
func (r *ResnetConfig) ValidateResnetConfig() error {
	if len(r.GetResnetPort()) == 0 || len(r.GetResnetHost()) == 0 || len(r.GetResnetPredictPath()) == 0 {
		return errors.New("resnet config data imvalid")
	}
	return nil
}
