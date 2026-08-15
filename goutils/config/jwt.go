package config

// import (
// 	"crypto/rsa"
// 	"os"

// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/bvv_utils/goutils/mistake"
// )

// type JwtConfig struct {
// 	PublicKey string `conf:"default:./key/fflare.test.pub,env:JWT_PUBLIC_KEY"`
// }

// func (j *JwtConfig) OutPubKey() (*rsa.PublicKey, error) {
// 	publicKey, err := os.ReadFile(j.PublicKey)
// 	if err != nil {
// 		return nil, mistake.ErrReadPubKey
// 	}
// 	k, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
// 	if err != nil {
// 		return nil, mistake.ErrLoadPubKey
// 	}
// 	return k, nil
// }
