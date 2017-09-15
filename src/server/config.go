package server

import "fmt"

const (
	DefaultGrpcPort = 18870
	DefaultDBAddr   = "todo"
)

type Config struct {
	Debug      bool
	GrpcPort   int
	HealthPort int
	DBAddr     string
}

func (c *Config) GetGrpcPortString() string {
	return fmt.Sprintf(":%d", c.GrpcPort)
}

func NewConfig() *Config {
	return &Config{
		GrpcPort: DefaultGrpcPort,
		DBAddr:   DefaultDBAddr,
	}
}
