package enqueuer

import (
	"github.com/gauravgahlot/syncroot/internal/types"
	"go.uber.org/zap"
)

// EnqueueRequest represents a request to enqueue an object for processing.
type EnqueueRequest struct {
	Topic     string
	Operation types.Operation
	Object    types.Object
}

type Enqueuer interface {
	Enqueue(request EnqueueRequest) error
}

func NewEnqueuer(logger *zap.Logger) (Enqueuer, error) {
	// in real world, we would be creating a Kafka connection here.

	return &enqueuer{
		logger: logger,
	}, nil
}

type enqueuer struct {
	logger *zap.Logger
}

func (e *enqueuer) Enqueue(request EnqueueRequest) error {
	e.logger.Info("Enqueuing request",
		zap.String("topic", request.Topic),
		zap.String("operation", string(request.Operation)),
		zap.String("objectID", request.Object.GetID()),
	)

	return nil
}
