package grpc_tracing

import (
	"context"

	"grpc-boilerplate/pkg/shared/util/metadata"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	grpcTag = opentracing.Tag{Key: string(ext.Component), Value: "sdb-svc"}
)

// UnaryServerInterceptor returns a new unary server interceptor for OpenTracing.
func UnaryServerInterceptor(tracer opentracing.Tracer, opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateOptions(opts)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if o.filterOutFunc != nil && !o.filterOutFunc(ctx, info.FullMethod) {
			return handler(ctx, req)
		}

		// TODO: need enhance > change param "service" using config service name
		// o.tracer = opentracing.GlobalTracer()

		opName := info.FullMethod
		if o.opNameFunc != nil {
			opName = o.opNameFunc(info.FullMethod)
		}

		// Create new server span
		newCtx, serverSpan := newServerSpanFromInbound(ctx, tracer, o.traceHeaderName, opName)
		if o.unaryRequestHandlerFunc != nil {
			o.unaryRequestHandlerFunc(serverSpan, req)
		}
		resp, err := handler(newCtx, req)
		// finishServerSpan(ctx, serverSpan, o.closer, err)
		return resp, err
	}
}

func newServerSpanFromInbound(ctx context.Context, tracer opentracing.Tracer, traceHeaderName, opName string) (context.Context, opentracing.Span) {
	// extract incoming metadata
	md := metadata.ExtractIncoming(ctx)

	// Get tracer for incoming metadata
	parentSpanContext, err := tracer.Extract(opentracing.HTTPHeaders, metadata.MetadataTextMap(md))
	if err != nil && err != opentracing.ErrSpanContextNotFound {
		grpclog.Infof("grpc_opentracing: failed parsing trace information: %v", err)
	}

	// Start Root Span using opName. opName is fullMethod (endpoint of current grpc)
	// Currently: grpcTag is unique each service (ServiceName)
	serverSpan := tracer.StartSpan(
		opName,
		// this is magical, it attaches the new span to the parent parentSpanContext, and creates an unparented one if empty.
		ext.RPCServerOption(parentSpanContext),
		grpcTag,
	)

	// Set Tag metadata
	metaData := metadata.MappingData(ctx)
	serverSpan.SetTag("md.client-name", metaData.ClientName)
	serverSpan.SetTag("md.client-ip", metaData.ClientIP)
	serverSpan.SetTag("md.activity-id", metaData.ActivityId)
	serverSpan.SetTag("md.frontend-ip", metaData.FrontEndIp)
	serverSpan.SetTag("md.uuid", metaData.Uuid)
	serverSpan.SetTag("md.x-trace-id", metaData.XTraceId)

	// Comment this if want delivery x-trace-id in each span
	//serverSpan.SetBaggageItem("md.x-trace-id-id", metaData.ClientName)

	// inject id of opentracing to tags
	injectOpentracingIdsToTags(traceHeaderName, serverSpan, Extract(ctx))
	// opentracing.SetGlobalTracer(tracer)
	return opentracing.ContextWithSpan(ctx, serverSpan), serverSpan
}
