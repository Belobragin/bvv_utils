package log

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestIDMiddleware(t *testing.T) {
	// Create a new request with a dummy URL
	req, err := http.NewRequest("GET", "/dummy", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a dummy handler

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Write a response
		w.WriteHeader(http.StatusOK)
	})

	// Create the middleware handler
	middleware := RequestIDMiddleware(handler)

	// Serve the request using the middleware
	middleware.ServeHTTP(rr, req)

	// Check if the response status code is correct
	if rr.Code != http.StatusOK {
		t.Errorf("Unexpected status code, expected %d, got %d", http.StatusOK, rr.Code)
	}
}
