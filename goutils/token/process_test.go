package token

// import (
// 	"testing"

// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/stretchr/testify/assert"
// )

// func Test_RenewToken_P(t *testing.T) {
// 	var (
// 		arrayTests = []struct {
// 			description string
// 			expTime     time.Time
// 			inKey       interface{}
// 			outS        string
// 			outE        error
// 		}{
// 			{"positive_test-1", time.Time{}, MockPrivateKey, "", nil},
// 			{"positive_test-2", MockFakeExp, MockPrivateKey, MockFakeToken, nil},
// 			{"positive_test-2", MockExp, MockPrivateKey, MockToken, nil},
// 		}
// 	)

// 	for i, tt := range arrayTests {
// 		tt := tt
// 		t.Run(tt.description, func(t *testing.T) {
// 			s, e := MockFakeTokenData.RenewToken(tt.inKey, tt.expTime)
// 			assert.Nil(t, e, "expected nil, get %v", e)
// 			if i >= 1 {
// 				assert.Equal(t, tt.outS, s)
// 			}
// 		})
// 	}
// }

// func Test_ParseValidateToken_N(t *testing.T) {
// 	var (
// 		arrayTests = []struct {
// 			description string
// 			tData       TokenData
// 			token       string
// 			inKey       interface{}
// 			outE        error
// 		}{
// 			{"negative_test-1", MockTokenData, MockToken, MockFakePublicKey, jwt.ErrTokenSignatureInvalid},
// 			{"negative_test-2", MockTokenData, MockInvalidToken, MockPublicKey, jwt.ErrTokenSignatureInvalid},
// 			{"negative_test-3", MockFakeTokenData, MockFakeToken, MockPublicKey, jwt.ErrTokenExpired},
// 		}
// 	)

// 	for _, tt := range arrayTests {
// 		tt := tt
// 		t.Run(tt.description, func(t *testing.T) {
// 			sbj := &tt.tData
// 			e := sbj.ParseValidateRsaToken(tt.token, tt.inKey)
// 			assert.ErrorIs(t, e, tt.outE, "")
// 		})
// 	}
// }

// func Test_ParseValidateToken_P(t *testing.T) {
// 	var (
// 		arrayTests = []struct {
// 			description string
// 			tData       TokenData
// 			token       string
// 			inKey       interface{}
// 		}{
// 			{"positive_test", MockTokenData, MockToken, MockPublicKey},
// 		}
// 	)

// 	for _, tt := range arrayTests {
// 		tt := tt
// 		t.Run(tt.description, func(t *testing.T) {
// 			sbj := &tt.tData
// 			e := sbj.ParseValidateRsaToken(tt.token, tt.inKey)
// 			assert.Nil(t, e, "")
// 		})
// 	}
// }
