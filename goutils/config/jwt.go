package config

import (
	"github.com/belobragin/bvv_utils/goutils/mistake"
	"github.com/golang-jwt/jwt/v5"
)

type JwtConfigI interface {
	GetKey(func(string) ([]byte, error)) (any, error)
}

type JwtConfig struct {
	PublicKey string `conf:"env:JWT_PUBLIC_KEY"`
}

func (c *JwtConfig) ValidateJwtConfig() error {
	return nil
}

// argument foo is a function which provides key as a byte sequence and a read key error
func (j *JwtConfig) GetKey(foo func(string) ([]byte, error)) (any, error) {
	publicKey, err := foo(j.PublicKey)
	if err != nil {
		return nil, mistake.ErrReadPubKey
	}
	k, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, mistake.ErrLoadPubKey
	}
	return any(k), nil
}

type JwtTokenConfigI interface {
	GetPubKey(func(string) ([]byte, error)) (any, error)
	GetSecretKey(func(string) ([]byte, error)) (any, error)
}

type JwtTokenConfig struct {
	PublicKey string `conf:"env:JWT_PUBLIC_KEY"`
	SecretKey string `conf:"env:JWT_KEY"`
}

func (c *JwtTokenConfig) ValidateJwtConfig() error {
	return nil
}

// argument foo is a function which provides key as a byte sequence and a read key error
func (j *JwtTokenConfig) GetPubKey(foo func(string) ([]byte, error)) (any, error) {
	publicKey, err := foo(j.PublicKey)
	if err != nil {
		return nil, mistake.ErrReadPubKey
	}
	k, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, mistake.ErrLoadPubKey
	}
	return any(k), nil
}

// argument foo is a function which provides key as a byte sequence and a read key error
func (j *JwtTokenConfig) GetSecretKey(foo func(string) ([]byte, error)) (any, error) {
	secretKey, err := foo(j.SecretKey)
	if err != nil {
		return nil, mistake.ErrReadSecretKey
	}
	k, err := jwt.ParseRSAPrivateKeyFromPEM(secretKey)
	if err != nil {
		return nil, mistake.ErrLoadSecretKey
	}
	return any(k), nil
}

// import (
// 	"crypto/rsa"
// 	"os"

// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/belobragin/bvv_utils/goutils/mistake"
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
