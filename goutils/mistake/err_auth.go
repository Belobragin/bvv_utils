package mistake

import "errors"

var (
	ErrNilPassword = errors.New("nil password, nothing to hash")
	ErrNulUser     = errors.New("user uuid must not be nil")
)
