package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/bvv_utils/goutils/mistake"
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

// inpCookie, inpHash are b64 encoded [32]byte and []byte
func CompareCookieWithHash(inpCookie, inpHash string) bool {
	d1, e := base64.StdEncoding.DecodeString(inpCookie)
	// fmt.Println("d1: ", d1)
	if e != nil {
		return false
	}
	h1 := sha256.Sum256(d1)
	// fmt.Println("h1: ", h1)
	d2, e := base64.StdEncoding.DecodeString(inpHash)
	if e != nil {
		return false
	}
	// fmt.Println("d2: ", d2)
	a := hex.EncodeToString(h1[:])
	b := hex.EncodeToString(d2)
	// fmt.Println("processed input cookie to compare:", a)
	// fmt.Println("processed input hash to compare:", b)
	return a == b
}

// make hash from password:
func ProduceHashFromPsw(psw *string) (string, error) {
	if psw == nil {
		return "", mistake.ErrNilPassword
	}
	d1, e := base64.StdEncoding.DecodeString(*psw)
	// fmt.Println("d1: ", d1)
	if e != nil {
		return "", e
	}
	h1 := sha256.Sum256(d1)
	return hex.EncodeToString(h1[:]), nil
}

// compare password and password hash
func PswHashCorrect(hash, psw string) bool {
	d1, e := base64.StdEncoding.DecodeString(psw)
	// fmt.Println("d1: ", d1)
	if e != nil {
		return false
	}
	h1 := sha256.Sum256(d1)
	// fmt.Println("h1: ", h1)
	d2, e := base64.StdEncoding.DecodeString(hash)
	if e != nil {
		return false
	}
	// fmt.Println("d2: ", d2)
	a := hex.EncodeToString(h1[:])
	b := hex.EncodeToString(d2)
	// fmt.Println("processed input cookie to compare:", a)
	// fmt.Println("processed input hash to compare:", b)
	return a == b
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
