package model

import (
	dbCfg "grpc-boilerplate/internal/config/db"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"grpc-boilerplate/pkg/shared/enum"
)

var (
	errorInvalidDbInstance = status.Error(codes.Internal, "Invalid db instance")
)

type DbDriver interface {
	Db() interface{}
}

var instanceDb map[string]DbDriver = make(map[string]DbDriver)

func NewInstanceDb(config dbCfg.DatabaseConfig) (DbDriver, error) {
	var err error
	var dbName = config.Dbname

	switch config.Adapter {
	case enum.MySql:
		dbConn, sqlErr := NewMySQLDriver(config)
		if sqlErr != nil {
			err = sqlErr
		}
		instanceDb[dbName] = dbConn
	default:
		err = errorInvalidDbInstance
	}

	return instanceDb[dbName], err
}
