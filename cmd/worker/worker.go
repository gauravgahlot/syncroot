package worker

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/gauravgahlot/syncroot/internal/config"
	"github.com/gauravgahlot/syncroot/internal/types"
	"github.com/gauravgahlot/syncroot/internal/worker"
)

// Command returns the command to start a worker.
func Command(logger *zap.Logger, cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "worker",
		Short: "Start a worker",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runE(cmd.Context(), logger, cfg)
		},
	}
}

func runE(ctx context.Context, logger *zap.Logger, cfg *config.Config) error {
	workerType := os.Getenv("WORKER_TYPE")

	var topic string
	switch types.WorkerType(workerType) {
	case types.WorkerTypeForwarder:
		topic = cfg.Forwarder.Topic
	case types.WorkerTypeSyncer:
		topic = cfg.Syncer.Topic
	case types.WorkerTypeDLQHandler:
		topic = cfg.DLQHandler.Topic
	case types.WorkerTypeWebhookHandler:
		topic = cfg.WebhookHandler.Topic
	}

	factory := worker.Factory{}
	w, err := factory.NewWorker(logger, topic, types.WorkerType(workerType))
	if w == nil {
		return fmt.Errorf("failed creating a worker: %w", err)
	}

	if err != nil {
		return fmt.Errorf("failed creating a worker: %w", err)
	}

	logger.Info("starting worker",
		zap.String("id", os.Getenv("WORKER_ID")),
		zap.String("type", workerType),
	)

	for {
		err := w.Work(ctx)
		if err != nil {
			logger.Error("worker error",
				zap.String("id", os.Getenv("WORKER_ID")),
				zap.String("type", workerType),
				zap.Error(err),
			)

			os.Exit(1)
		}
	}
}
