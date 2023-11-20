package class

import (
	"context"
	"grpc-boilerplate/pkg/domain/entity"
	proto "grpc-boilerplate/pkg/infrastructure/grpc/proto/class"
)

type ClassInputPort interface {
	GetClass(ctx context.Context, req *proto.GetClassRequest) (*proto.GetClassResponse, error)
	GetClassDistinct(ctx context.Context, req *proto.GetClassDistinctRequest) (*proto.GetClassDistinctResponse, error)
}

type ClassOutputPort interface {
	GetClassResponse(*entity.GetClassResponse) (*proto.GetClassResponse, error)
	GetClassDistinctResponse(*entity.GetClassDistinctResponse) (*proto.GetClassDistinctResponse, error)
}
