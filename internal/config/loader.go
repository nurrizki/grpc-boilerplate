package config

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"grpc-boilerplate/internal/config/db"
	"grpc-boilerplate/internal/config/server"

	"grpc-boilerplate/pkg/shared/util"

	"github.com/spf13/viper"
)

type config struct {
	Database db.DatabaseList
	Server   server.ServerList
}

var cfg config

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {
	viper.AddConfigPath(basepath + "/db")
	viper.SetConfigType("yaml")
	viper.SetConfigName("database.yml")
	err := viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("cannot load database config: %v", err))
	}

	viper.AddConfigPath(basepath + "/server")
	viper.SetConfigType("yaml")
	viper.SetConfigName("server.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("cannot load server config: %v", err))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.Unmarshal(&cfg)

	fmt.Println("==============CONFIG===============")
	fmt.Println(util.Stringify(cfg))
	fmt.Println("===================================")
}

func GetConfig() *config {
	return &cfg
}
