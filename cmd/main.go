package main

import (
	"log"
	"os"

	"go.uber.org/zap"

	"github.com/gauravgahlot/syncroot/cmd/root"
	"github.com/gauravgahlot/syncroot/internal/config"
)

func main() {
	cfg, err := config.NewFromEnv()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	logger, err := initializeLogger(cfg.Environment)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	logger.Info("configguration: %+v", zap.Any("config", cfg))

	defer logger.Sync()
	if err = root.Command(logger, cfg).Execute(); err != nil {
		logger.Error("syncer service exited with error", zap.Error(err))

		os.Exit(1)
	}
}

func initializeLogger(env string) (*zap.Logger, error) {
	if env == "dev" || env == "development" || env == "local" {
		return zap.NewDevelopment()
	}

	return zap.NewProduction()
}
