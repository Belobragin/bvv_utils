package token

import (
	"errors"
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
)

func (d *TokenData) IssueRsaToken(jc interface{}) (string, error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodRS256, d)
	tokenString, e := newToken.SignedString(jc)
	if e != nil {
		return "", e
	}
	return tokenString, nil
}

func (d *TokenData) IssueECDSAToken(jc interface{}) (string, error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodES512, d)
	tokenString, e := newToken.SignedString(jc)
	if e != nil {
		return "", e
	}
	return tokenString, nil
}

func (d *TokenData) ParseValidateRsaToken(inputKey, jc interface{}) error {
	var (
		e, err error
		// j = new(jwt.Token)
	)
	t := fmt.Sprintf("%v", inputKey)
	_, e = jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		return jc, nil
	}, jwt.WithValidMethods([]string{"RS256"}))
	switch {
	case errors.Is(e, jwt.ErrTokenExpired):
		// d, ok := j.Claims.(TokenData)
		// fmt.Printf("%+v\n", d)
		// fmt.Println(ok)
		err = jwt.ErrTokenExpired
	case e != nil:
		return e
	default:
	}
	jwt.ParseWithClaims(t, d, func(t *jwt.Token) (interface{}, error) {
		return jc, nil
	})
	return err
}

func (d *TokenData) ParseValidateECDSAToken(t string, jc interface{}) error {
	var (
		err error
		// j = new(jwt.Token)
	)
	return err
}
