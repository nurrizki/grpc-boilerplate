package entity

import "grpc-boilerplate/pkg/adapter/db/model"

type GetClassResponse struct {
	Data []model.ClassDb
}

type GetClassRequest struct {
	ClassName string
	Status    string
	Level     string
}

type GetClassDistinctRequest struct {
	Status string
}

type ClassDistinct struct {
	ClassName string
}

type GetClassDistinctResponse struct {
	Data []ClassDistinct
}
