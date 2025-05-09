package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gauravgahlot/syncroot/internal/config"
)

func TestNewFromEnv(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	os.Setenv("SERVER_PORT", "5000")
	os.Setenv("WORKER_COUNT", "2")
	os.Setenv("WORKER_TYPE", "syncer")

	cfg, err := config.NewFromEnv()
	assert.NoError(t, err)

	assert.Equal(t, "test", cfg.Environment)
	assert.Equal(t, 5000, cfg.Server.Port)
	assert.Equal(t, 2, cfg.Worker.Count)
	assert.Equal(t, config.Syncer, cfg.Worker.Type)
}
