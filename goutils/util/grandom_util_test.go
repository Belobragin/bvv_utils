package util

import (
	"testing"
)

func Test_CompareCookieWithHash(t *testing.T) {
	// example hash from token 'context_id' field (b64 encoded string)
	ih := "DX1+xnh4FLf0YBSgNYMv8Oa/oWQY2ZGaUlXiwAotFPM="
	// example b64 encoded auth_cookie value:
	ic := "/RHBUVwUeVuqDIjBZETqhEr8kwr+/IoI3y2MCDl92r/HAztmvKKeFZntIeUkFWcxT60DbEP247PWrcY6IBORdw=="
	if !CompareCookieWithHash(ic, ih) {
		t.Errorf("cookie hash and hash from cookie are different; must be equal")
	}
}

// func Test_GenerateRandomSmsCode(t *testing.T) {
// 	p, e := GenerateRandomSmsCode()
// 	fmt.Println(p, e)
// }

// func Test_GenerateRandomEmailCode(t *testing.T) {
// 	p, e := GenerateRandomEmailCode()
// 	fmt.Println(p, e)
// }
