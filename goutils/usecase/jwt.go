package usecase

import (
	"github.com/belobragin/bvv_utils/goutils/config"
	"github.com/belobragin/bvv_utils/goutils/token"
)

type JwtConfigI interface {
	OutPubKey() (any, error)
	GetToken() token.TokenDataI
}

type StandardJwtRealization struct {
	key   any
	token token.TokenDataI
}

func (s *StandardJwtRealization) OutPubKey() (any, error) {
	return s.key, nil
}

func (s *StandardJwtRealization) GetToken() token.TokenDataI {
	return s.token
}

// argument foo is a function which provides key as a byte sequence and a read key error
func (s *StandardJwtRealization) NewStandardJwtRealization(
	c config.JwtConfigI,
	t token.TokenDataI,
	foo func(string) ([]byte, error),
) error {
	s.token = t
	var err error
	s.key, err = c.GetKey(foo)
	return err
}

type JwtTokenConfigI interface {
	OutPubKey() (any, error)
	OutSecretKey() (any, error)
	GetToken() token.TokenDataI
}

type TokenJwtRealization struct {
	pubKey    any
	secretKey any
	token     token.TokenDataI
}

func (s *TokenJwtRealization) OutPubKey() (any, error) {
	return s.pubKey, nil
}
func (s *TokenJwtRealization) OutSecretKey() (any, error) {
	return s.secretKey, nil
}

func (s *TokenJwtRealization) GetToken() token.TokenDataI {
	return s.token
}

// argument foo is a function which provides key as a byte sequence and a read key error
func (s *TokenJwtRealization) NewStandardJwtRealization(
	c config.JwtTokenConfigI,
	t token.TokenDataI,
	foo func(string) ([]byte, error),
) error {
	s.token = t
	var err error
	s.pubKey, err = c.GetPubKey(foo)
	if err != nil {
		return err
	}
	s.secretKey, err = c.GetSecretKey(foo)
	if err != nil {
		return err
	}
	return nil
}
