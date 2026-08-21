package mistake

import "errors"

// metrica errors
var (
	ErrNulMetrica            = errors.New("metrica object null")
	ErrApiMetricVectorNull   = errors.New("no api call metrica counter_vec")
	ErrEventMetricVectorNull = errors.New("event metrica counter_vec null")
	ErrMetricPort            = errors.New("metric port unset")
	ErrInvalidMetricRegister = errors.New("can not tegister prometheus metric vector")
	ErrNoSuchMetrica         = errors.New("metrica is not initialized ")
	ErrMetricaLabel          = errors.New("metrica label incorrect")
)
