package mistake

import "errors"

// user procesiing errors:
var (
	ErrUserLoginNil = errors.New("user login null")
	ErrUserLoginLen = errors.New("user login length invalid")
	ErrUserPswNil   = errors.New("user password null")
	ErrUserPswLen   = errors.New("user password length invalid")
)
