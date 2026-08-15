package mistake

import "errors"

// metrica errors
var (
	ErrNulMetrica            = errors.New("metrica object null")
	ErrApiMetricVectorNull   = errors.New("no api call metrica counter_vec")
	ErrMetricPort            = errors.New("metric port unset")
	ErrInvalidMetricRegister = errors.New("can not tegister prometheus metric vector")
	ErrNoSuchMetrica         = errors.New("metrica is not initialized ")
	ErrMetricaLabel          = errors.New("metrica label incorrect")
)
