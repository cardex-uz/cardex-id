package config

import (
	grpc "cardex-id/pkg/grpc"
	pg "cardex-id/pkg/postgresql"
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env    string      `yaml:"env" env-default:"local"`
	Server grpc.Config `yaml:"server"`
	DB     pg.Config   `yaml:"db"`
	//MigrationPath string        `json:"migrationPath" yaml:"migrationPath"`
	//TokenTTL      time.Duration `json:"token-ttl" yaml:"token-ttl"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("Config path is empty")
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	// check config file
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("Config file does not exist: " + configPath)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("Cannot read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var value string

	flag.StringVar(&value, "config", "", "Path to config file")
	flag.Parse()

	if value == "" {
		value = os.Getenv("CONFIG_PATH")
	}

	return value
}
