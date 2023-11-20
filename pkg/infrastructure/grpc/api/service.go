package api

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	classSVC "grpc-boilerplate/pkg/adapter/api/grpc/class"
	classProto "grpc-boilerplate/pkg/infrastructure/grpc/proto/class"
	"grpc-boilerplate/pkg/infrastructure/interceptor"
	"grpc-boilerplate/pkg/infrastructure/interceptor/grpc_tracing"
	container "grpc-boilerplate/pkg/shared/di"
	"grpc-boilerplate/pkg/shared/enum"

	cfg "grpc-boilerplate/internal/config"
	"grpc-boilerplate/pkg/shared/tracing"

	classUC "grpc-boilerplate/pkg/usecase/class"

	"github.com/opentracing/opentracing-go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func RunServer() {
	log.Println("starting grpc server ...")
	config := cfg.GetConfig()

	opts := []grpc.ServerOption{}

	tracer, closer := tracing.Init("grpc-boilerplate")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	opts = append(opts, grpc.UnaryInterceptor(interceptor.ChainUnaryServer(
		grpc_tracing.UnaryServerInterceptor(tracer),
		grpc_tracing.UnaryServerInterceptorPanic(),
	)))

	opts = append(opts, grpc.KeepaliveParams(keepalive.ServerParameters{
		// MaxConnectionAge:      5 * time.Minute,
		// MaxConnectionIdle:     1 * time.Hour,
		// MaxConnectionAgeGrace: 5 * time.Second,
	}))

	grpcServer := grpc.NewServer(opts...)
	ctn := container.NewContainer()

	Apply(grpcServer, ctn)
	reflection.Register(grpcServer)

	svcHost := config.Server.Grpc.Host
	svcPort := config.Server.Grpc.Port
	// svcHost := "0.0.0.0"
	// svcPort := 2202

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", svcHost, svcPort))
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start rebuild MS server: %v", err)
		}
	}()

	fmt.Printf("Test gRPC server is running at %s:%d\n", svcHost, svcPort)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v\n", signal.String())
}

func Apply(server *grpc.Server, ctn *container.Container) {
	classProto.RegisterClassServiceServer(server, classSVC.NewClassService(ctn.Resolve(string(enum.ClassCTN)).(*classUC.ClassInteractor)))
}
