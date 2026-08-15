package database

import (
	"os"
	"testing"
)

const (
	testDSN     = "port=5432 host=localhost user=brl password=brl12345 dbname=avitohr sslmode=disable"
	testDSNFail = "port=5432 host=localhost user=brl password=brl12345 dbname=avitohr sslmode=enable"
)

// NEGATIVE:
func Test_checkDsnSSL1(t *testing.T) {
	if e := checkDsnSSL(testDSNFail); e == nil {
		t.Fatalf("error was expected:  want %v, get nil", e)
	}
}

// POSITIVE:
func Test_prepareDSN1(t *testing.T) {
	if _, e := prepareDSN(testDSN, "", SslEnvModeType, SslModeDisableDefault); e != nil {
		t.Fatalf("error was expected:  want %v, get nil", e)
	}
	if _, e := prepareDSN(testDSN, SslEnvModeType, "", SslModeDisableDefault); e != nil {
		t.Fatalf("error was expected:  want %v, get nil", e)
	}
}

// NEGATIVE:
func Test_prepareDSN2(t *testing.T) {
	os.Setenv(SslEnvModeType, "disable")
	os.Setenv(SslEnvRequired, "true")
	if _, e := prepareDSN(testDSN, SslEnvRequired, SslEnvModeType, SslModeDisableDefault); e == nil {
		t.Fatalf("error was expected:  want %v, get nil", e)
	}

	if _, e := prepareDSN(testDSN, SslEnvRequired, SslEnvModeType, SslModeDisableDefault); e == nil {
		t.Fatalf("error was expected:  want %v, get nil", e)
	}
}

// NEGATIVE:
func Test_prepareDSN3(t *testing.T) {
	os.Setenv(SslEnvModeType, sslModeVerifyCa)
	os.Setenv(SslEnvRequired, "true")
	if _, e := prepareDSN(testDSN, SslEnvRequired, SslEnvModeType, SslModeDisableDefault); e == nil {
		t.Fatalf("error was expected:  want %v, get nil", e)
	}
}

// POSITIVE:
func Test_prepareDSN4(t *testing.T) {
	os.Setenv(SslEnvModeType, "disable")
	os.Setenv(SslEnvRequired, "false")
	if _, e := prepareDSN(testDSN, SslEnvRequired, SslEnvModeType, SslModeDisableDefault); e != nil {
		t.Fatalf("error was not expected:  want nil, get %v", e)
	}
}
