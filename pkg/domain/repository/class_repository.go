package repository

import (
	"context"
	"grpc-boilerplate/pkg/adapter/db/model"
	"grpc-boilerplate/pkg/domain/entity"

	"github.com/opentracing/opentracing-go"
)

type ClassRepository interface {
	GetClassDb(ctx context.Context, span opentracing.Span, req entity.GetClassRequest) ([]model.ClassDb, error)
	GetClassDistinct(ctx context.Context, span opentracing.Span, req entity.GetClassDistinctRequest) ([]entity.ClassDistinct, error)
}
