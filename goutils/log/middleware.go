package log

import (
	"net/http"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r = r.WithContext(assignRequestID(ctx))
		next.ServeHTTP(w, r)
	})
}
