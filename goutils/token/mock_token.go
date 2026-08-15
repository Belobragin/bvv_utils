package token

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	MockToken        = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDU2NzgiLCJmX3VzZXJfaWQiOiIiLCJwYXJ0bmVyX2lkIjoiODc2NTQzMjEiLCJzaXRlX2lkIjoiS1oiLCJyb2xlX2lkIjoiIiwiaXNzIjoiTW9ja1NlcnZpY2UiLCJzdWIiOiIxMjM0NTY3OCIsImF1ZCI6WyJGaW5uRmxhcmVGYXZvcml0ZSJdLCJleHAiOjE4NDk4NjY5MjZ9.yf5Xgpg3w9p-MPQxXgZSLgSZVXlnkyST1-K966X6DW5v-PkJk--U_oZuhWhT0f1eUgCsef31IpB_gTo5Nixu_zhGLbOzUDAv0pB4hYeF7NAu-Fgp1SpGDUGi48dpcE_Dn7rlAX_cHPkl0MzsMs2FWikoifVaoKBgV54EuAFHBW_nsAVo3A41i3KTH5sOBLlbNz1dfBo73aGjimdGC1wkO41LpQbojIDiTaAirCHf5R6nctedpdLYHWnFs1B7uCb7pCJY64ZcamkfmOAQEcLr-AOBJ68a86s4-lEOWMQZcprm0rMV_w9h-oOASlRkRJxjzHItXTnVompIFzQtSGSYFQ`
	MockInvalidToken = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDU2NzgiLCJwYXJ0bmVyX2lkIjoiODc2NTQzMjEiLCJzaXRlX2lkIjoiS1oiLCJpc3MiOiJGaW5uRmxhcmVKd3QiLCJzdWIiOiIxMjM0NTY3OCIsImF1ZCI6WyJGaW5uRmxhcmVGYXZvcml0ZSJdLCJleHAiOjE4NDk4NjY5MjZ9.D8ty2mIHCWTVb-KUS9-k3zZph0-qmkejygXkO_BAoZY0oY9TMcI2YSz6kGAj09M20p1thxbo3eyDzpXaf81-tIHLApjtImsiGEfJk0L7y-11nWtnXWNcy1uK5YNGfYR6D6KolYJainKWf-VChsIpEzPvHDty5CapZ0_xZksvo49ySCyRMcK5PZGQyao4neGofvRC1eE5QlA-YpJhhaS_lq0C8NiUcpQjGlugXUKaknfAZfY9P-oOv8Yz-rbKKM2mdFdJG2wmqneUKNGwYG4t-SRhKTkSBpKAmgOwGno0-GBf0LCM5CXaMT3ZBBIwf3qd8xhMudBYINQsW0QnvnR9fQ`
	MockFakeToken    = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDU2NzgiLCJmX3VzZXJfaWQiOiIiLCJwYXJ0bmVyX2lkIjoiODc2NTQzMjEiLCJzaXRlX2lkIjoiS1oiLCJyb2xlX2lkIjoiIiwiaXNzIjoiTW9ja1NlcnZpY2UiLCJzdWIiOiIxMjM0NTY3OCIsImF1ZCI6WyJGaW5uRmxhcmVGYXZvcml0ZSJdLCJleHAiOjE2NjA0NzgxMjZ9.mAV4_bY7_E33xyZq209HNLxuJ1wbzb2rQknSkhTU8sgOL7cLv8VJdCI2EaTTWwisVdDI8K63XFm0z3CD7mk5gMKPf1TBLXEGkpwF8HZoQKgbidZ1gzeGpQEUKuCZ8adiZcyXfynNIMKkkKVj95_z4wt7vhdEkkQEvkN-Dlgts9sKD_7X0tKoopF9S8PNYUOMXofauHw7cnpik2cI27Cn_h4WSeGjdaSg4Gy5dlMY-uu9OqD5kdpu0ksOrXvlRfPuW05O3PQDsnl7DbpUTj5jxiEMV9JgMbHejeSGpBja39GmhwZkhULjo4NLCayDAleMFxQ5uXwH72Pk5K1EkTrPXg`
	MockSub          = int64(12345678)
	MockSubS         = "12345678"
	MockIss          = "MockService"
	MockAud          = "FinnFlareFavorite"
	MockTime         = "2028-08-14T11:55:26.371"
	MockFakeTime     = "2022-08-14T11:55:26.371"
	MockSiteID       = "KZ"
	MockFakeSiteID   = "BR"
	MockUserID       = int64(12345678)
	MockUserUUIDS    = "12345678"
	MockPartnerID    = int64(87654321)
	MockPartnerUUIDS = "87654321"
	MockPassword     = "mypassword"
	MockPhone        = "9122222345"
	MockLogin        = "testpartner"
	MockPswHash      = "$2a$15$wssLh8t0.xP6SxZbfGvj2uKeBCmPBYhAGF7s3h05ErLb9nAXEPAOi"
	MockUserUUID     = "b3b05871-cc1c-4969-b9ea-7fdd950cf4f8"
	MockPartnerUUID  = "1ec491a7-9242-63e8-91db-73221cb4dc63"

	TimeFormatContext = "2006-01-02T15:04:05.000"
	// additional output models (all entities) time format:
	// TimeFormat        = "2006-01-02T15:04:05-07:00"
	// TimeFormatContext1 = "2006-01-02"
)

var (
	MockExp, _     = time.Parse(TimeFormatContext, MockTime)
	MockExpAt      = jwt.NumericDate{MockExp}
	MockFakeExp, _ = time.Parse(TimeFormatContext, MockFakeTime)
	MockFakeExpAt  = jwt.NumericDate{MockExp}

	SqlMockSub         = sql.NullInt64{Valid: true, Int64: MockSub}
	SqlMockIss         = sql.NullString{Valid: true, String: MockIss}
	SqlMockAud         = []string{MockAud}
	SqlMockExp         = sql.NullTime{Valid: true, Time: MockExp}
	SqlMockUserUUID    = sql.NullInt64{Valid: true, Int64: MockUserID}
	SqlMockPartnerUUID = sql.NullInt64{Valid: true, Int64: MockPartnerID}
	SqlMockSiteID      = sql.NullString{Valid: true, String: MockSiteID}

	registeredClaims = jwt.RegisteredClaims{
		Issuer:    MockIss,
		Subject:   MockSubS,
		Audience:  jwt.ClaimStrings{MockAud},
		ExpiresAt: &MockExpAt,
	}
)

var (
	TestPubKey = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3g2VNW7zLxsPdBOFQYz7
4xfDlDH0z5TNGRqTJ2N2yP7vaI3PGdh9fj4HrhYLlpmVL+Q8AFvTl1Hs/pOGw0Wr
S8FhfPCJqP3XsIsdqtyPS+/eBXDZjz9SdR5UfIt4bm+Khu0BzykV04CzFH96jFK4
R0vSrQWvsf6vdvJVZTyGMV5/vrkqQOPyBmytTlXiqn6V98zDvjAVJoyv8AMbY7C9
LkyEpYL6xJwsmPnnkbUKeB8zkcvzABI8OtqTiIIBHdsdlNuqUIlu+szYyJTL4mlH
+X251AhiWtyl0s/UGKccomELXaKIm8t3AoLc0ydzT1FHCcO5eDaluvCaaW+89AnG
rwIDAQAB
-----END PUBLIC KEY-----`)

	TestPrivateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA3g2VNW7zLxsPdBOFQYz74xfDlDH0z5TNGRqTJ2N2yP7vaI3P
Gdh9fj4HrhYLlpmVL+Q8AFvTl1Hs/pOGw0WrS8FhfPCJqP3XsIsdqtyPS+/eBXDZ
jz9SdR5UfIt4bm+Khu0BzykV04CzFH96jFK4R0vSrQWvsf6vdvJVZTyGMV5/vrkq
QOPyBmytTlXiqn6V98zDvjAVJoyv8AMbY7C9LkyEpYL6xJwsmPnnkbUKeB8zkcvz
ABI8OtqTiIIBHdsdlNuqUIlu+szYyJTL4mlH+X251AhiWtyl0s/UGKccomELXaKI
m8t3AoLc0ydzT1FHCcO5eDaluvCaaW+89AnGrwIDAQABAoIBAQCcTE9EZy5Bb/i2
qa7YbxY1yNpfi7JYEtMy4mPGC9rmq5t0qqyl8p6Sp7LcNKQ/gydhqHMbhYcwMfhx
UQJAnN9v2VR3jxA2pUaW+7UMOWeb+eD7T22zyMyPp6Osbrd8dIK7akQlvzOfxi1H
NlvaPbd8Z8CJDUGYNj7nqLe3JJ0ITb2jtAtC8qS5xU/aqXR5oAPikY/sWWyd/4AB
Izh7h6wzQ2i6d8a0qKLICiL/dRVfvfisreZIx/3xGP4st53plz1k5W7LDaWPGXdc
xwU3VRyOcTzpP7z+8Vcb6LddXWKqgxSi2hv0Vd4AzU8bmgwi0XDrQ6+ZEtNIGwsN
uZhWi0oBAoGBAP4vvEYUXn+YnNuW5kjCUunGfqyisA7E+vG9AlpKr3Q0GLklyfiy
aymZ/6ZNhuRLw6Hyxs71O4wOub7m55bFs7ourGbibzgRbNKcrxOBr8cNEJM4pSeD
cI54vZdAFf2m7iizFku84mYQEJLUZ5uvDG8M/7j5nfqJTwRdTk2B5owvAoGBAN+j
KBg5GNgr+6bJFljnn7/YhtNyr9hvJ2ccxpu+wfSNel+6R1BW2pvLCjiC7D/UaY0l
Buwh7yswOs5zYHEE8BCnpHAfkvU2TAsQl21sgpBXtgq+4PseMvmh+NMpAj4X7i23
7dxW9XvzD4qH+MsGGBf95PYc2DKkSGM4chHYvU2BAoGALcgzRPwOOTUnq3V2bMxz
Wp/h2P5Eb6SS/k9oyRKCn49ylIZ15lFcc9XE8hMspJnw1o5/uG77FrLgSb4VIFbj
lDkr3CwUlCivTQ/jiMPnARcdUSb2uDM4ZrDglQl57IHQi4wWC/YtYOyrz4ZqUuQo
fSf+NzbeeW2ZydN0cwlemQkCgYA5xYo8B7P1VBdkVJojSWKpSqQ5x2zJup0xM+vS
nAq6xDmStSprBKTgjikLxDaHnrXNsn0BTxrby0/FuWR2jdH+W3BNTrrGAcrPKPAV
YmywIRhz2i+Ab6K3fjIrSnK72665vMQDoas04+tl3A5XVDGxMPOO6JCpNGtGxnQh
TB5LAQKBgQDpHIGv896oZiF1YFTUfApI+ZV2Tp2btKT5a/sh8QvUFTdkIKHIwxi6
2Oh8PbLwTgB58mjW76MTa/KAdKb5tUVmGB/BG7zlOgEQw/UX5F+tSlOfdibSzPSz
3KKOdsy65QWvpm8cQMlvgmGjDu4R1U0FWGhfOXY5kdrjeMpF1J0NBw==
-----END RSA PRIVATE KEY-----`)

	FakePubKey = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3g2VNW7zLxsPdBOFQYz7
4xfDlDH0z5TNGRqTJ2N2yP7vaI3PGdh9fj4HrhYLlpmVL+Q8AFvTl1Hs/pOGw0Wr
S8FhfPCJqP3XsIsdqtyPS+/eBXDZjz9SdR5UfIt4bm+Khu0BzykV04CzFH96jFK4
R0vSrQWvsf6vdvJVZTyGMV5/vrkqQOPyBmytTlXiqn6V98zDvjAVJoyv8AMbY7C9
LkyEpYL6xJwsmPnnkbUKeB8zkcvzABI8OtqTiIIBHdsdlNuqUIlu+szYyJTL4mlH
+X251AhiWtyl0s/UGKccomELXaKIm8t3AoLc0ydzT1FHCcO5eDaluvCaaW+89AnG
rwIDAQAC
-----END PUBLIC KEY-----`)

	FakePrivateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA3g2VNW7zLxsPdBOFQYz74xfDlDH0z5TNGRqTJ2N2yP7vaI3P
Gdh9fj4HrhYLlpmVL+Q8AFvTl1Hs/pOGw0WrS8FhfPCJqP3XsIsdqtyPS+/eBXDZ
jz9SdR5UfIt4bm+Khu0BzykV04CzFH96jFK4R0vSrQWvsf6vdvJVZTyGMV5/vrkq
QOPyBmytTlXiqn6V98zDvjAVJoyv8AMbY7C9LkyEpYL6xJwsmPnnkbUKeB8zkcvz
ABI8OtqTiIIBHdsdlNuqUIlu+szYyJTL4mlH+X251AhiWtyl0s/UGKccomELXaKI
m8t3AoLc0ydzT1FHCcO5eDaluvCaaW+89AnGrwIDAQABAoIBAQCcTE9EZy5Bb/i2
qa7YbxY1yNpfi7JYEtMy4mPGC9rmq5t0qqyl8p6Sp7LcNKQ/gydhqHMbhYcwMfhx
UQJAnN9v2VR3jxA2pUaW+7UMOWeb+eD7T22zyMyPp6Osbrd8dIK7akQlvzOfxi1H
NlvaPbd8Z8CJDUGYNj7nqLe3JJ0ITb2jtAtC8qS5xU/aqXR5oAPikY/sWWyd/4AB
Izh7h6wzQ2i6d8a0qKLICiL/dRVfvfisreZIx/3xGP4st53plz1k5W7LDaWPGXdc
xwU3VRyOcTzpP7z+8Vcb6LddXWKqgxSi2hv0Vd4AzU8bmgwi0XDrQ6+ZEtNIGwsN
uZhWi0oBAoGBAP4vvEYUXn+YnNuW5kjCUunGfqyisA7E+vG9AlpKr3Q0GLklyfiy
aymZ/6ZNhuRLw6Hyxs71O4wOub7m55bFs7ourGbibzgRbNKcrxOBr8cNEJM4pSeD
cI54vZdAFf2m7iizFku84mYQEJLUZ5uvDG8M/7j5nfqJTwRdTk2B5owvAoGBAN+j
KBg5GNgr+6bJFljnn7/htNyr9hvJ2ccxpu+wfSNel+6R1BW2pvLCjiC7D/UaY0l
Buwh7yswOs5zYHEE8BCnpHAfkvU2TAsQl21sgpBXtgq+4PseMvmh+NMpAj4X7i23
7dxW9XvzD4qH+MsGGBf95PYc2DKkSGM4chHYvU2BAoGALcgzRPwOOTUnq3V2bMxz
Wp/h2P5Eb6SS/k9oyRKCn49ylIZ15lFcc9XE8hMspJnw1o5/uG77FrLgSb4VIFbj
lDkr3CwUlCivTQ/jiMPnARcdUSb2uDM4ZrDglQl57IHQi4wWC/YtYOyrz4ZqUuQo
fSf+NzbeeW2ZydN0cwlemQkCgYA5xYo8B7P1VBdkVJojSWKpSqQ5x2zJup0xM+vS
nAq6xDmStSprBKTgjikLxDaHnrXNsn0BTxrby0/FuWR2jdH+W3BNTrrGAcrPKPAV
YmywIRhz2i+Ab6K3fjIrSnK72665vMQDoas04+tl3A5XVDGxMPOO6JCpNGtGxnQh
TB5LAQKBgQDpHIGv896oZiF1YFTUfApI+ZV2Tp2btKT5a/sh8QvUFTdkIKHIwxi6
2Oh8PbLwTgB58mjW76MTa/KAdKb5tUVmGB/BG7zlOgEQw/UX5F+tSlOfdibSzPSz
3KKOdsy65QWvpm8cQMlvgmGjDu4R1U0FWGhfOXY5kdrjeMpF1J0NBx==
-----END RSA PRIVATE KEY-----`)

	MockPrivateKey, _     = jwt.ParseRSAPrivateKeyFromPEM(TestPrivateKey)
	MockPublicKey, _      = jwt.ParseRSAPublicKeyFromPEM(TestPubKey)
	MockFakePrivateKey, _ = jwt.ParseRSAPrivateKeyFromPEM(FakePrivateKey)
	MockFakePublicKey, _  = jwt.ParseRSAPublicKeyFromPEM(FakePubKey)
)

// var (
// 	MockTokenData = TokenData{
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			Issuer:    MockIss,
// 			Subject:   MockSubS,
// 			Audience:  jwt.ClaimStrings{MockAud},
// 			ExpiresAt: &MockExpAt,
// 		},
// 		UserUUID:    MockUserUUIDS,
// 		PartnerUUID: MockPartnerUUIDS,
// 		SiteID:      MockSiteID,
// 	}
// 	MockFakeTokenData = TokenData{
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			Issuer:    MockIss,
// 			Subject:   MockSubS,
// 			Audience:  jwt.ClaimStrings{MockAud},
// 			ExpiresAt: &MockFakeExpAt,
// 		},
// 		UserUUID:    MockUserUUIDS,
// 		PartnerUUID: MockPartnerUUIDS,
// 		SiteID:      MockSiteID,
// 	}
// )
