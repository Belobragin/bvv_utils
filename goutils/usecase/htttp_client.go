package usecase

import "net/http"

type HttpClientI interface {
	GetHttpClient() *http.Client
}
