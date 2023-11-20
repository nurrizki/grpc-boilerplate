package class

import (
	"context"
	proto "grpc-boilerplate/pkg/infrastructure/grpc/proto/class"
	"grpc-boilerplate/pkg/shared/tracing"
	"grpc-boilerplate/pkg/usecase/class"
	// "github.com/golang/protobuf/ptypes/empty"
)

type ClassService struct {
	uc class.ClassInputPort
}

func NewClassService(uc class.ClassInputPort) *ClassService {
	return &ClassService{uc: uc}
}

func (s *ClassService) GetClass(ctx context.Context, req *proto.GetClassRequest) (*proto.GetClassResponse, error) {
	sp, _ := tracing.CreateRootSpan(ctx, "Get Class")
	defer sp.Finish()

	res, err := s.uc.GetClass(ctx, req)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (s *ClassService) GetClassDistinct(ctx context.Context, in *proto.GetClassDistinctRequest) (*proto.GetClassDistinctResponse, error) {
	sp, _ := tracing.CreateRootSpan(ctx, "Get Class Distinct")
	defer sp.Finish()

	res, err := s.uc.GetClassDistinct(ctx, in)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
