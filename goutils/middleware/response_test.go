package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.uber.org/zap"
)

func Test_beforeWriteHeader(t *testing.T) {
	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Create a mock zap logger
	logger, _ := zap.NewDevelopment()

	// Call the function under test
	beforeWriteHeader(logger, rr, req, http.StatusOK, nil)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Add more assertions here if needed
}
func TestStandardRespond(t *testing.T) {
	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Create a mock zap logger
	logger, _ := zap.NewDevelopment()

	// Call the function under test
	StandardRespond(logger, rr, req, http.StatusOK, "Hello, World!")

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := `"Hello, World!"`
	if strings.TrimSuffix(rr.Body.String(), "\n") != expectedBody {
		t.Errorf("Expected response body %v, but got %v",
			expectedBody, rr.Body.String())
	}

	// Add more assertions here if needed
}
