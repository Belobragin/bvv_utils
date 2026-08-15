package util

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"github.com/bvv_utils/goutils/mistake"
	"golang.org/x/crypto/bcrypt"
)

func MakePswHash(p string) string {
	b := []byte(p)
	h := sha1.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
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
func PswHashCorrectBC(hash, psw string) bool {
	d1, e := base64.StdEncoding.DecodeString(psw)
	// fmt.Println("d1: ", d1)
	if e != nil {
		return false
	}
	d2, e := base64.StdEncoding.DecodeString(hash)
	if e != nil {
		return false
	}
	// fmt.Println("d2: ", d2)
	err := bcrypt.CompareHashAndPassword(d2, d1)
	return err == nil

	// h1 := sha256.Sum256(d1)
	// fmt.Println("h1: ", h1)

	// fmt.Println("d2: ", d2)
	// a := hex.EncodeToString(h1[:])
	// b := hex.EncodeToString(d2)
	// fmt.Println("processed input cookie to compare:", a)
	// fmt.Println("processed input hash to compare:", b)
	// return a == b
}

func PswHashCorrect(hash, psw string) bool {
	d1, e := base64.StdEncoding.DecodeString(psw)
	if e != nil {
		return false
	}
	h1 := sha256.Sum256(d1)
	a := hex.EncodeToString(h1[:])
	return a == hash
}
