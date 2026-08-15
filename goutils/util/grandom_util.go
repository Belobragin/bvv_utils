package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"

	"github.com/google/uuid"
)

const (
	fingerPrintLen = 64
)

func GenerateRandomContext(inpRaw *string) (string, string, error) {
	var h [32]byte
	b := make([]byte, fingerPrintLen)
	if inpRaw != nil {
		h = sha256.Sum256([]byte(*inpRaw))
	} else {
		_, err := rand.Read(b)
		if err != nil {
			return "", "", err
		}
		h = sha256.Sum256(b)
	}
	/* return:
	 	- b64 encoded hash of random value => token 'context_id' field
		- b64 encoded random value (not-hashed) => auth_cookie
	*/
	return base64.StdEncoding.EncodeToString([]byte(h[:])),
		base64.StdEncoding.EncodeToString(b),
		nil
}

func GenerateRandomEmailCode() (string, error) {
	return uuid.NewString()[1:7], nil
	// return "v76tF2", nil
}

func GenerateRandomSmsCode() (string, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(9999-1000))
	if err != nil {
		return "", err
	}
	return fmt.Sprint(1000 + nBig.Int64()), nil
	// return "5599", nil
}
