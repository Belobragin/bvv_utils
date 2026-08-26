package mistake

import "fmt"

type OuterrorI interface {
	Err() error
	ErrCode() int
}

type Outerror struct {
	err  error
	code int
}

func (o Outerror) Err() error {
	return o.err
}

func (o Outerror) ErrCode() int {
	if o != (Outerror{}) {
		return o.code
	}
	return CodeOutErrNull
}

func NewOutErr(e error, c int) Outerror {
	return Outerror{
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
