package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SSL      string
	}
	HttpServer struct {
		Port string
	}
}

var C config

type ReadConfigOption struct {
	AppEnv string
}

func ReadConfig(options ReadConfigOption) {
	Config := &C

	viper.AddConfigPath(filepath.Join(rootDir(), "config"))
	viper.SetConfigName("config")

	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("could not read configuration: %v", err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("could not unmarshal config: %v", err)
		os.Exit(1)
	}
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
