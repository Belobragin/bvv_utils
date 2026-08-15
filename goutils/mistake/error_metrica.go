package mistake

import "errors"

var (
	ErrInvalidMetricRegister = errors.New("can not tegister prometheus metric vector")
	ErrNoSuchMetrica         = errors.New("metrica is not initialized ")
	ErrMetricaLabel          = errors.New("metrica label incorrect")
)
