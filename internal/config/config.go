package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config holds the configuration for Syncroot.
type Config struct {
	Environment string `default:"development" split_words:"true"`
	Server      Server
	Worker      Worker
}

// Server holds the server configuration.
type Server struct {
	Port int `default:"3000"`
}

type WorkerType string

const (
	Forwarder WorkerType = "forwarder"
	Syncer    WorkerType = "syncer"
)

type Worker struct {
	Count int        `default:"1"`
	Type  WorkerType `default:"forwarder"`
}

func NewFromEnv() (*Config, error) {
	c := &Config{}
	if err := envconfig.Process("", c); err != nil {
		return nil, err
	}

	return c, nil
}
