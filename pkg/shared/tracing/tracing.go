package tracing

import (
	"context"
	"fmt"
	"grpc-boilerplate/pkg/shared/enum"
	"grpc-boilerplate/pkg/shared/util"
	"io"
	"net/http"
	"reflect"
	"runtime"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go/config"
)

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Init(service string) (opentracing.Tracer, io.Closer) {
	// Add Opentracing instrumentation
	defcfg := config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	cfg, err := defcfg.FromEnv()
	if err != nil {
		panic("Could not parse Jaeger env vars: " + err.Error())
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic("Could not initialize jaeger tracer: " + err.Error())
	}

	return tracer, closer
}

// TraceFunction wraps funtion with opentracing span adding tags for the function name and caller details
func TraceFunction(ctx context.Context, fn interface{}, params ...interface{}) (result []reflect.Value) {
	// Get function name
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	// Create child span
	parentSpan := opentracing.SpanFromContext(ctx)
	sp := opentracing.StartSpan(
		"Function - "+name,
		opentracing.ChildOf(parentSpan.Context()))
	defer sp.Finish()

	sp.SetTag("function", name)

	// Get caller function name, file and line
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	callerDetails := fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)
	sp.SetTag("caller", callerDetails)

	// Check params and call function
	f := reflect.ValueOf(fn)
	if f.Type().NumIn() != len(params) {
		e := fmt.Sprintf("Incorrect number of parameters calling wrapped function %s", name)
		panic(e)
	}
	inputs := make([]reflect.Value, len(params))
	for k, in := range params {
		inputs[k] = reflect.ValueOf(in)
	}
	return f.Call(inputs)
}

func CreateRootSpan(ctx context.Context, name string) (opentracing.Span, opentracing.Tracer) {
	parentSpan := opentracing.SpanFromContext(ctx)
	tracer := parentSpan.Tracer()

	// Get caller function name, file and line
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	callerDetails := fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)
	parentSpan.SetTag("caller", callerDetails)

	return parentSpan, tracer
}

func StartRootSpan(ctx context.Context, name string) (context.Context, io.Closer, opentracing.Span) {
	tracer, closer := Init(name)
	sp := tracer.StartSpan(string(enum.StartService))
	ctx = opentracing.ContextWithSpan(ctx, sp)
	return ctx, closer, sp
}

// CreateChildSpan creates a new opentracing span adding tags for the span name and caller details. Returns a Span.
// User must call `defer sp.Finish()`
func CreateChildSpan(ctx context.Context, name string) opentracing.Span {
	parentSpan := opentracing.SpanFromContext(ctx)
	sp := opentracing.StartSpan(
		name,
		opentracing.ChildOf(parentSpan.Context()))
	sp.SetTag("name", name)

	// Get caller function name, file and line
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	callerDetails := fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)
	sp.SetTag("caller", callerDetails)

	return sp
}

func CreateSubChildSpan(parentSpan opentracing.Span, name string) opentracing.Span {
	sp := opentracing.StartSpan(
		name,
		opentracing.ChildOf(parentSpan.Context()))
	sp.SetTag("name", name)

	// Get caller function name, file and line
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	callerDetails := fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)
	sp.SetTag("caller", callerDetails)

	return sp
}

func LogRequest(sp opentracing.Span, req interface{}) {
	sp.LogFields(log.Object(string(enum.Request), util.Stringify(req)))
}

func LogObject(sp opentracing.Span, name string, obj interface{}) {
	sp.LogFields(log.Object(name, util.Stringify(obj)))
}

func LogResponse(sp opentracing.Span, resp interface{}) {
	sp.LogFields(log.Object(string(enum.Response), util.Stringify(resp)))
}

func LogError(sp opentracing.Span, err error) {
	sp.LogFields(log.Object(string(enum.Error), err))
}

// NewTracedRequest generates a new traced HTTP request with opentracing headers injected into it
func NewTracedRequest(method string, url string, body io.Reader, span opentracing.Span) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err.Error())
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, method)
	span.Tracer().Inject(span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header))

	return req, err
}
