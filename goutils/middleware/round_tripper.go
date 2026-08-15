package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

type proxied interface {
	RoundTrip(*http.Request) (*http.Response, error)
}

type AuthRoundTripper struct {
	Proxied proxied
	// BaseURL   string
	Zapstruct *zap.Logger
}

func (m *AuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	zapstruct := m.Zapstruct
	// TODO: this is apiv3 dependable, if any:
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t))
	res, err := m.Proxied.RoundTrip(req)
	if err != nil {
		zapstruct.Error("Error in client request transport middleware", zap.Error(err))
		return res, err
	}
	return res, nil
}
