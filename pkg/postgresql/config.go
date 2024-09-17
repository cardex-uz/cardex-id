package postgresql

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

const defaultPort = 5432

type Config struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Name     string `json:"name" yaml:"name"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	SSLMode  string `json:"ssl_mode" yaml:"ssl_mode"`
}

func (c *Config) withDefaults() (conf Config) {
	if c != nil {
		conf = *c
	}

	if conf.Port == 0 {
		conf.Port = defaultPort
	}

	return
}

func MustLoadPath(configPath string) *Config {
	// check config file
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("Config file does not exist: " + configPath)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		panic("Cannot read config: " + err.Error())
	}

	return &cfg
}
