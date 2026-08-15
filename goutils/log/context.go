package log

import (
	"context"

	"github.com/google/uuid"
)

type contextKey string

const contextKeyRequestID contextKey = "requestID"

func assignRequestID(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKeyRequestID, uuid.NewString())
}

func GetRequestID(ctx context.Context) string {
	requestID := ctx.Value(contextKeyRequestID)
	if ret, ok := requestID.(string); ok {
		return ret
	}

	return ""
}
