package container

import (
	"grpc-boilerplate/pkg/adapter/connection"
	"grpc-boilerplate/pkg/adapter/repository/db"
	"grpc-boilerplate/pkg/shared/enum"
	"grpc-boilerplate/pkg/usecase/class"

	"grpc-boilerplate/pkg/infrastructure/grpc/service/dto"

	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, _ := di.NewBuilder()

	_ = builder.Add([]di.Def{
		{Name: string(enum.ClassCTN), Build: classUsecase},
	}...)
	return &Container{
		ctn: builder.Build(),
	}
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func classUsecase(_ di.Container) (interface{}, error) {
	repo := db.NewClassDb(connection.MiniProject)
	// return repo, nil
	out := &dto.ClassBuilder{}
	return class.NewClassInteractor(repo, out), nil
}
