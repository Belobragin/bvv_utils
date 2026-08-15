package token

/*
this type is an output sent to the token consumer, with fingerprint
*/
type TokenFingerOutput struct {
	// nullable:false
	// example: eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiYjNiMDU4NzEtY2MxYy00OTY5LWI5ZWEtN2ZkZDk1MGNmNGY4Iiwic2l0ZV9pZCI6IktaIiwicm9sZV9pZCI6IjQiLCJpc3MiOiJGaW5uRmxhcmVKd3QiLCJzdWIiOiIzIiwiYXVkIjpbInN1YnNjcmlwdGlvbiIsImZhdm9yaXRlIiwidXNlciIsImRlbGl2ZXJ5Il0sImV4cCI6MTY5NzAzODIzMH0.0wSK4AtmE_mPfehAW-Zq3h1BjPsRVUxGHgmuUrMTu8KgajJmXEPDjwt5_U8oW9Lp1wmb__nIAXzTDgOEk-gAqPRWd6QV236oCeFrpQIEG_swNFOU6Zpi9odpsb4ll1QoNuQsjuEYyGSu5KP0LcWP5bBz7nJvqBxrNCkkOaTzVDQMHYVaP6GXrGunReT5W43uAZvJAz0Xa7dJkQfHPNb0HL5w5N9lJ6IjlhpdtILv2ueHqp9jq-uaktG7Vn1HHmiWbOMyA4xy3Vu-mf8_Yyx639FbVFxc5R9YEk6dJJ1QplNoYCaOuvL7cF-Hd1rYrA5vAklkDnDaIJusMlLj7tqEKg
	// Extensions:
	// x-order: "0"
	Token string `json:"token"`
	// nullable:false
	// example: LC+OfcFr5rKZ+D13ko9yOUkSlujki6JhJsXnc7rtwz6nA1wbPAkcpx8zm5xsXOcKAf9AWOo6csxe0QfAcRdkHQ==
	// Extensions:
	// x-order: "1"
	Fingerprint string `json:"isst_auth_cookie"`
}
