package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config holds the configuration for Syncroot.
type Config struct {
	Environment    string `default:"development" split_words:"true"`
	Server         Server
	Forwarder      Worker
	Syncer         Worker
	DLQHandler     Worker `split_words:"true"`
	WebhookHandler Worker `split_words:"true"`
}

// Server holds the server configuration.
type Server struct {
	Port int `default:"3000"`
}

type Worker struct {
	Count int `default:"1"`
	Topic string
}

func NewFromEnv() (*Config, error) {
	c := &Config{}
	if err := envconfig.Process("", c); err != nil {
		return nil, err
	}

	return c, nil
}
