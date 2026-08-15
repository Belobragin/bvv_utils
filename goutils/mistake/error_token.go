package mistake

import (
	"errors"
)

// tokens mistakes:
var (
	ErrInvalidToken          = errors.New("token is invalid")
	ErrInvalidTokenData      = errors.New("token data invalid")
	ErrIllegalInputTokenData = errors.New("invalid input data for token")
	ErrTokenContext          = errors.New("token context is invalid")
	ErrNullAuthCtxValues     = errors.New("context does not contain auth data or role boundary exclude api invocation")
	ErrNoTokenServicePanel   = errors.New("No token service panel")
	ErrOldTokenDataNull      = errors.New("old token data are null on refresh")
	ErrOldTokenDataInvalid   = errors.New("old token data invalid on refresh")
	ErrOldTokenDataCorrupt   = errors.New("old token data corrupt on refresh")
)

// key input-output mistakes:
var (
	ErrReadPubKey    = errors.New("failed to read jwt pub key")
	ErrLoadPubKey    = errors.New("failed to load pem format of public key")
	ErrReadSecretKey = errors.New("failed to read key")
	ErrLoadSecretKey = errors.New("failed to load pem format of secret key")
	ErrReadInput     = errors.New("failed to read other token data")
)

// input-output mistakes:
var (
	ErrTokenNoRoleID       = errors.New("no parsed role ID")
	ErrTokenNoUserUUID     = errors.New("no parsed user uuid")
	ErrTokenNoSiteID       = errors.New("no parsed site id")
	ErrTokenNoAlien        = errors.New("no parsed alien")
	ErrTokenNoIsAdminID    = errors.New("no parsed is_admin ")
	ErrTokenInvalidIsAdmin = errors.New("isAdmin field must be bool")
	ErrConfirmSiteID       = errors.New("invalid site id")
	ErrNoSuchAlien         = errors.New("input alen value does not exist")
)
