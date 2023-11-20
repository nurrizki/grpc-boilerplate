package model

// import (
// 	dbCfg "grpc-boilerplate/internal/config/db"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type DriverPostgreSQL struct {
// 	config dbCfg.DatabaseConfig
// 	db     *gorm.DB
// }

// func NewPostgreSQLDriver(config dbCfg.DatabaseConfig) (DbDriver, error) {
// 	dbConn, err := ConnectPostgreSQL(config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &DriverPostgreSQL{
// 		config: config,
// 		db:     dbConn,
// 	}, nil
// }

// func ConnectPostgreSQL(config dbCfg.DatabaseConfig) (*gorm.DB, error) {
// 	user := config.Username
// 	password := config.Password
// 	host := config.Host
// 	port := config.Port
// 	dbname := config.Dbname

// 	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// 	dsn := user + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"

// 	dbConn, err := gorm.Open(postgres.Open(dsn))
// 	return dbConn, err
// }

// // Db get db instance of gorm
// func (m *DriverPostgreSQL) Db() interface{} {
// 	return m.db
// }
