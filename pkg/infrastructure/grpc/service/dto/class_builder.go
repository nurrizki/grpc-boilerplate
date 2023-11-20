package dto

import (
	"grpc-boilerplate/pkg/domain/entity"
	proto "grpc-boilerplate/pkg/infrastructure/grpc/proto/class"

	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClassBuilder struct{}

func (*ClassBuilder) GetClassResponse(in *entity.GetClassResponse) (*proto.GetClassResponse, error) {
	var out *proto.GetClassResponse
	if err := mapstructure.Decode(in, &out); err != nil {
		return &proto.GetClassResponse{}, status.Error(codes.Internal, "parseFailed")
	}

	return out, nil
}

func (*ClassBuilder) GetClassDistinctResponse(in *entity.GetClassDistinctResponse) (*proto.GetClassDistinctResponse, error) {
	var out *proto.GetClassDistinctResponse
	if err := mapstructure.Decode(in, &out); err != nil {
		return &proto.GetClassDistinctResponse{}, status.Error(codes.Internal, "parseFailed")
	}

	return out, nil
}
