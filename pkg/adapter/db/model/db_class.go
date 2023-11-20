package model

import "grpc-boilerplate/pkg/shared/enum"

type ClassDb struct {
	ClassId          int32  `gorm:"column:class_id"`
	ClassName        string `gorm:"column:class_name"`
	ClassDescription string `gorm:"column:class_desc"`
	Level            string `gorm:"column:level"`
	Status           string `gorm:"column:status"`
}

func (ClassDb) TableName() string {
	return string(enum.ClassDbTable)
}
