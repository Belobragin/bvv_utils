package mistake

import "errors"

var (
	ErrNilPassword       = errors.New("nil password, nothing to hash")
	ErrEmptyAuthCookie   = errors.New("authentication cookie empty")
	ErrNillAuthCookie    = errors.New("authentication cookie not set")
	ErrInvalidAuthCookie = errors.New("authentication cookie invalid")

	ErrEmptySessionCookie   = errors.New("session cookie empty")
	ErrNillSessionCookie    = errors.New("session cookie not set")
	ErrInvalidSessionCookie = errors.New("session cookie invalid")

	ErrJwtTokenInvalid        = errors.New("jwt token not valid")
	ErrTokenExpired           = errors.New("token expired")
	ErrInvalidJAuthHeader     = errors.New("authorization header invalid")
	ErrNoContextID            = errors.New("token structure invalid: no context id")
	ErrCookieSignatureInvalid = errors.New("cookie contradict new token signature")

	ErrIncorrectPassword = errors.New("incorrect password")
)
