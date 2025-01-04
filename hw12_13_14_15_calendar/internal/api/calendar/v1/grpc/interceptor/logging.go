package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
)

// LoggingInterceptor is a unary interceptor for logging requests.
func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	logger.Debugf("Received request: %v", info.FullMethod)

	resp, err := handler(ctx, req)
	if err != nil {
		logger.Errorf("Error handling request: %v", err)
		return nil, status.Errorf(status.Code(err), "LoggingInterceptor: %v", err)
	}

	logger.Debugf("Request handled successfully: %v", info.FullMethod)
	return resp, nil
}
