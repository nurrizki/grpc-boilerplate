package grpc_tracing

import (
	"context"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecoveryHandlerFunc func(p interface{}) (err error)

// RecoveryHandlerFuncContext is a function that recovers from the panic `p` by returning an `error`.
// The context can be used to extract request scoped metadata and context values.
type RecoveryHandlerFuncContext func(ctx context.Context, p interface{}) (err error)

// UnaryServerInterceptor returns a new unary server interceptor for panic recovery.
func UnaryServerInterceptorPanic(opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateOptions(opts)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				err = recoverFrom(ctx, r, o.recoveryHandlerFunc, info.FullMethod, debug.Stack())
			}
		}()

		resp, err := handler(ctx, req)
		panicked = false
		return resp, err
	}
}

func recoverFrom(ctx context.Context, p interface{}, r RecoveryHandlerFuncContext, method string, trace interface{}) error {
	if r == nil {
		// logger.WriteLogPanic(logger.Error, "PANIC", fmt.Sprintf("%v", p), method, codes.Internal, trace)
		return status.Errorf(codes.Internal, "%v", p)
	}
	return r(ctx, p)
}
