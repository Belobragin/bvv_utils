package log

import (
	"bufio"
	"bytes"
	"encoding/json"
	"testing"

	"go.uber.org/zap"
)

const (
	testInfoMessage = "Http server started"
	testPort        = "9990"
	testLevelI      = "info"
	testCaller      = "log/logger_test.go:29"
)

type testInfo struct {
	Level   string
	Caller  string
	Message string
}

func Test_StructuredInit1(t *testing.T) {
	var b bytes.Buffer
	bWriter := bufio.NewWriter(&b)
	w := StructuredInit("INFO", TestCustomIO, bWriter)
	w.Info(testInfoMessage, zap.String("port", testPort))
	bWriter.Flush()
	var y testInfo
	err := json.Unmarshal(b.Bytes(), &y)
	if err != nil {
		t.Fatalf("Error on console logger unmarshalling was not expected: %v", err)
	}
	if m := y.Message; m != testInfoMessage {
		t.Fatalf("expected message %s, get %s", testInfoMessage, m)
	}
	if m := y.Level; m != testLevelI {
		t.Fatalf("expected level %s, get %s", testLevelI, m)
	}
	if m := y.Caller; m != testCaller {
		t.Fatalf("expected caller %s, get %s", testCaller, m)
	}
}

func Test_StructuredInit3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("mistake on StructuredInit() was not expected: get: %v", r)
		}
	}()
	w := StructuredInit("DEBUG", "", nil)
	w.Info(testInfoMessage, zap.String("port", testPort))
	w.Error(testInfoMessage, zap.String("port", testPort))
}
