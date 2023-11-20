package class

import (
	"context"

	"grpc-boilerplate/pkg/domain/entity"
	"grpc-boilerplate/pkg/domain/repository"
	proto "grpc-boilerplate/pkg/infrastructure/grpc/proto/class"
	"grpc-boilerplate/pkg/shared/enum"
	"grpc-boilerplate/pkg/shared/tracing"
)

type ClassInteractor struct {
	repo repository.ClassRepository
	out  ClassOutputPort
}

func NewClassInteractor(r repository.ClassRepository, out ClassOutputPort) ClassInputPort {
	return &ClassInteractor{
		repo: r,
		out:  out,
	}
}

func (u *ClassInteractor) GetClass(ctx context.Context, req *proto.GetClassRequest) (*proto.GetClassResponse, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, req)

	res, err := u.repo.GetClassDb(ctx, sp, entity.GetClassRequest{
		ClassName: req.ClassName,
		Status:    req.Status,
		Level:     req.Level,
	})
	if err != nil {
		tracing.LogError(sp, err)
		return &proto.GetClassResponse{}, err
	}

	tracing.LogResponse(sp, res)
	return u.out.GetClassResponse(&entity.GetClassResponse{Data: res})
}

func (u *ClassInteractor) GetClassDistinct(ctx context.Context, req *proto.GetClassDistinctRequest) (*proto.GetClassDistinctResponse, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, req)

	res, err := u.repo.GetClassDistinct(ctx, sp, entity.GetClassDistinctRequest{
		Status: req.Status,
	})
	if err != nil {
		tracing.LogError(sp, err)
		return &proto.GetClassDistinctResponse{}, err
	}

	tracing.LogResponse(sp, res)
	return u.out.GetClassDistinctResponse(&entity.GetClassDistinctResponse{Data: res})
}
