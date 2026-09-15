package mistake

import "fmt"

type OutErrI interface {
	Err() error
	ErrCode() int
}

type OutErr struct {
	err  error
	code int
}

func (o OutErr) Err() error {
	return o.err
}

func (o OutErr) ErrCode() int {
	if o != (OutErr{}) {
		return o.code
	}
	return CodeOutErrNull
}

func NewOutErr(e error, c int) OutErr {
	return OutErr{
		err:  e,
		code: c,
	}
}

func NewErrS(e error, s string) error {
	return fmt.Errorf("%s: %w", s, e)
}
func NewAddErr(eNew error, ePrev error) error {
	switch {
	case eNew == nil && ePrev == nil:
		return nil
	case eNew == nil:
		return ePrev
	case ePrev == nil:
		return eNew
	default:
		return fmt.Errorf("%w: %w", eNew, ePrev)
	}
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
