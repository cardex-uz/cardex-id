package grpc

import "time"

type Config struct {
	Port    int           `json:"port" yaml:"port"`
	Timeout time.Duration `json:"timeout" yaml:"timeout"`
}
