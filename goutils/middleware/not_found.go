package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

func NotFound(z *zap.Logger, f func(a any) (http.Handler, error), a any) http.Handler {
	if w, err := f(a); err != nil {
		z.Error(err.Error())
		return http.NotFoundHandler()
	} else {
		return w
	}
}
