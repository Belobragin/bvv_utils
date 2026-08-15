package token

import (
	"github.com/golang-jwt/jwt/v5"
)

const (
	JwtOldClaimKey = "old_token"
)

// TokenDataI :: general interface - any token
type TokenDataI interface {
	GetStandardClaims() jwt.RegisteredClaims
	ParseValidateRsaToken(string, any) error
}

// MinimalisticTokenDataI :: minimalistic token (no autorenew etc.)
type MinimalisticTokenDataI interface {
	GetStandardClaims() jwt.RegisteredClaims
	ParseValidateRsaToken(string, any) error
	GetRoleID() (string, error)
}

// AutoRenewTokenDataI :: automatic renew under correct session cookie
type AutoRenewTokenDataI interface {
	GetSessionID() *string
	GetStandardClaims() jwt.RegisteredClaims
	ParseValidateRsaToken(string, any) error
	GetRoleID() (string, error)
}
