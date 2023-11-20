package main

import (
	"fmt"
	cfg "grpc-boilerplate/internal/config"

	grpcAPI "grpc-boilerplate/pkg/infrastructure/grpc/api"
)

func main() {
	config := cfg.GetConfig()

	fmt.Println(config)

	//run server
	grpcAPI.RunServer()
}
