package token

import (
	"github.com/golang-jwt/jwt/v5"
)

type TokenData struct {
	UserUUID  string `json:"user_uuid"`
	Role      string `json:"role"`
	ContextID string `json:"context_id"`
	jwt.RegisteredClaims
}
