package worker

import (
	"context"

	"github.com/gauravgahlot/syncroot/internal/db"
	"github.com/gauravgahlot/syncroot/internal/enqueuer"
	"github.com/gauravgahlot/syncroot/internal/providers"
	"github.com/gauravgahlot/syncroot/internal/types"
	"go.uber.org/zap"
)

type Worker interface {
	Work(ctx context.Context) error
}

func NewForwarder(logger *zap.Logger, providers []providers.Provider, db db.Store, topic string) (Worker, error) {
	// start listening on the topic

	return &forwarder{
		logger:    logger,
		db:        db,
		providers: providers,
		topic:     topic,
	}, nil
}

type forwarder struct {
	Worker

	logger    *zap.Logger
	db        db.Store
	topic     string
	providers []providers.Provider
}

func (w forwarder) Work(ctx context.Context) error {
	// receive enqueued requests (enqueuer.EnqueueRequest) from the topic
	// here is a sample request
	req := enqueuer.EnqueueRequest{
		Operation: types.OperationCreate,
		Object: &types.Contact{
			ID:       "3fa85f64-5717-4562-b3fc-2c963f66afa6",
			FullName: "John Doe",
			Email:    "john@doe.io",
		},
	}

	// first store the object in the database
	switch req.Operation {
	case types.OperationCreate:
		_, err := w.db.CreateContact(ctx, req.Object.(*types.Contact))
		if err != nil {
			w.logger.Error("failed creating contact", zap.Error(err))

			return err
		}

		// other cases like Update, Delete, etc.
	}

	// sync the object with the providers
	for _, provider := range w.providers {
		go provider.SyncProvider(ctx, req.Object)
	}

	return nil
}

type Syncer struct {
	Worker
}

func (w Syncer) Work(ctx context.Context) error {
	return nil
}

type DLQHandler struct {
	Worker
}

func (w DLQHandler) Work(ctx context.Context) error {
	return nil
}

type WebhookHandler struct {
	Worker
}

func (w WebhookHandler) Work(ctx context.Context) error {
	return nil
}
