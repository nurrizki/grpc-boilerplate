package connection

import (
	"fmt"
	cfg "grpc-boilerplate/internal/config"
	dbCfg "grpc-boilerplate/internal/config/db"
	dr "grpc-boilerplate/pkg/adapter/db"
	"log"
	"time"
)

var (
	SDB         dr.DbDriver
	MiniProject dr.DbDriver
)

func NewDbConn(database dbCfg.DatabaseConfig) (dr.DbDriver, error) {
	var dbConn dr.DbDriver
	var err error

	currentWaitTime := 2
	trialCount := 0

	for dbConn == nil && trialCount < 5 {
		trialCount++
		dbConn, err = dr.NewInstanceDb(database)
		if err != nil {
			fmt.Println("unable connecting to DB.")
			if trialCount == 5 {
				return nil, err
			}
			fmt.Println("retrying in", currentWaitTime, "second...")
			time.Sleep(time.Duration(currentWaitTime) * time.Second)
			currentWaitTime = currentWaitTime * 2
		}
	}
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

func init() {
	fmt.Println("Connect to All Databases")
	conf := cfg.GetConfig()
	var err error

	SDB, err = NewDbConn(conf.Database.MySQL.Sdb)
	if err != nil {
		log.Fatalf("unable to connect to NDS SDB")
	}

	MiniProject, err = NewDbConn(conf.Database.MySQL.MiniProject)
	if err != nil {
		log.Fatalf("unable to connect to NDS SDB")
	}
}
