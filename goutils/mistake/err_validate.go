package mistake

import "errors"

var (
	ErrIncorrectApiPortValue    = errors.New("incorrect api port value: must be among 1001 and 64000")
	ErrIncorrectMetricPortValue = errors.New("incorrect metric port value: must be among 1001 and 64000")
	ErrIncorrectLogLevel        = errors.New("incorrect log level, value must be in 'DEBUG', 'WARN', 'ERROR', 'PANIC', 'DPANIC', 'FATAL'")
)

var (
	ErrNilID = errors.New("id must be not null")
)
