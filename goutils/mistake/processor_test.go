package mistake

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/belobragin/bvv_utils/goutils/log"
)

const testErrMessage = "test error"

type testS struct {
	Error string
}

func Test_ErrorProcessor(t *testing.T) {
	var b bytes.Buffer
	bWriter := bufio.NewWriter(&b)
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("mistake on ErrorsProcessor() was not expected: get: %v", r)
		}
	}()
	c := make(chan struct{})
	errC := make(chan error, 1)
	w := log.StructuredInit("DEBUG", log.TestCustomIO, bWriter)
	go ErrorsProcessor(w, c, errC)
	errC <- errors.New(testErrMessage)
	time.Sleep(100 * time.Millisecond)
	close(c)
	bWriter.Flush()
	var y testS
	err := json.Unmarshal(b.Bytes(), &y)
	if err != nil {
		t.Fatalf("Error on console logger unmarshalling was not expected: %v", err)
	}
	if e := y.Error; e != testErrMessage {
		t.Fatalf("expected error %s, get %s", testErrMessage, e)
	}
}
