package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

func OrderUnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	m, err := handler(ctx, req)
	// логика после вызова
	return m, err
}
