package db

import (
	"context"

	dr "grpc-boilerplate/pkg/adapter/db"
	"grpc-boilerplate/pkg/adapter/db/model"
	"grpc-boilerplate/pkg/domain/entity"
	"grpc-boilerplate/pkg/shared/tracing"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ClassDb struct {
	repo dr.DbDriver
	db   *gorm.DB
}

func NewClassDb(driver dr.DbDriver) *ClassDb {
	return &ClassDb{
		repo: driver,
		db:   driver.Db().(*gorm.DB),
	}
}

func (r *ClassDb) GetClassDb(ctx context.Context, span opentracing.Span, req entity.GetClassRequest) ([]model.ClassDb, error) {
	sp := tracing.CreateSubChildSpan(span, "GetClassDb")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	var class []model.ClassDb

	db := r.db.Debug()

	if req.ClassName != "" {
		db = db.Where("class_name = ?", req.ClassName)
	}

	if req.Level != "" {
		db = db.Where("level = ?", req.Level)
	}

	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}

	if err := db.Find(&class).Error; err != nil {
		tracing.LogError(sp, status.Errorf(codes.Internal, err.Error()))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	tracing.LogResponse(sp, class)
	return class, nil
}

func (r *ClassDb) GetClassDistinct(ctx context.Context, span opentracing.Span, req entity.GetClassDistinctRequest) ([]entity.ClassDistinct, error) {
	sp := tracing.CreateSubChildSpan(span, "GetClassDistinct")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	var class []model.ClassDb
	var resp []entity.ClassDistinct

	db := r.db.Debug()

	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}

	if err := db.Distinct("class_name").Find(&class).Error; err != nil {
		tracing.LogError(sp, status.Errorf(codes.Internal, err.Error()))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, v := range class {
		resp = append(resp, entity.ClassDistinct{
			ClassName: v.ClassName,
		})
	}

	tracing.LogResponse(sp, class)
	return resp, nil
}
