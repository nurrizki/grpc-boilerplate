package logging

// import (
// 	"context"
// 	"grpc-boilerplate/pkg/shared/logger"
// 	"grpc-boilerplate/pkg/shared/util/metadata"

// 	"github.com/google/uuid"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// type validator interface {
// 	Validate() error
// }

// // LoggingServerInteceptor returns a new unary server interceptor for OpenTracing.
// func LoggingServerInteceptor() grpc.UnaryServerInterceptor {
// 	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 		m := metadata.MappingData(ctx)
// 		logUUID, _ := uuid.NewRandom()
// 		logger.WriteLog(logger.Info, req, m, "request", "", info.FullMethod, logUUID.String(), codes.OK)

// 		if v, ok := req.(validator); ok {
// 			if err := v.Validate(); err != nil {
// 				logger.WriteLog(logger.Error, req, m, "request", err.Error(), info.FullMethod, logUUID.String(), codes.InvalidArgument)
// 				return nil, status.Errorf(codes.InvalidArgument, err.Error())
// 			}
// 		}
// 		// Calls the handler
// 		res, err := handler(ctx, req)
// 		if err != nil {
// 			if e, ok := status.FromError(err); ok {
// 				logger.WriteLog(logger.Error, nil, m, "response", e.Message(), info.FullMethod, logUUID.String(), e.Code())
// 			}
// 		} else {
// 			logger.WriteLog(logger.Info, res, m, "response", codes.OK.String(), info.FullMethod, logUUID.String(), codes.OK)
// 		}
// 		return res, err
// 	}
// }
