package usecase

import "net/http"

type ResnetPredictI interface {
	GetPredictApiUrl() string
	GetTransport() http.RoundTripper
}
