package interceptors

import (
	"context"
	"fmt"
	"grpc/common"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func LoggerInterceptor(ctx context.Context,
	req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	reqStr := fmt.Sprintf("%v", req)
	logger := common.GetLogger()
	logger.Info(info.FullMethod, zap.String("request", reqStr))

	h, err := handler(ctx, req)

	return h, err
}
