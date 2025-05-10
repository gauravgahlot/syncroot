package worker

import (
	"fmt"

	"github.com/gauravgahlot/syncroot/internal/db"
	"github.com/gauravgahlot/syncroot/internal/providers"
	"github.com/gauravgahlot/syncroot/internal/types"
	"go.uber.org/zap"
)

type Factory struct{}

func (f Factory) NewWorker(logger *zap.Logger, topic string, workerType types.WorkerType) (Worker, error) {
	db := db.NewInMemoryStore()
	providers := providers.Initialize(logger)

	switch workerType {
	case types.WorkerTypeForwarder:
		{
			fwd, err := NewForwarder(logger, providers, db, topic)
			if err != nil {
				return nil, fmt.Errorf("failed creating a forwarder: %w", err)
			}

			return fwd, nil
		}
	case types.WorkerTypeSyncer:
		// to be created like forwarder
		return &Syncer{}, nil
	case types.WorkerTypeDLQHandler:
		// to be created like forwarder
		return &DLQHandler{}, nil
	case types.WorkerTypeWebhookHandler:
		// to be created like forwarder
		return &WebhookHandler{}, nil
	default:
		return nil, fmt.Errorf("invalid worker type: %s", workerType)
	}
}
