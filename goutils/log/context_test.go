package log

import (
	"context"
	"testing"
)

func TestGetRequestID(t *testing.T) {
	// Create a context with a request ID value
	ctx := context.WithValue(context.Background(), contextKeyRequestID, "12345")

	// Call the GetRequestID function
	requestID := GetRequestID(ctx)

	// Check if the returned request ID matches the expected value
	if requestID != "12345" {
		t.Errorf("Unexpected request ID, expected %s, got %s", "12345", requestID)
	}
}
