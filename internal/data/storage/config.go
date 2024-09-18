package storage

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
