package providers

import (
	"context"

	hs "github.com/gauravgahlot/syncroot/internal/transformer/hubspot"
	sf "github.com/gauravgahlot/syncroot/internal/transformer/salesforce"
	"github.com/gauravgahlot/syncroot/internal/types"
	"go.uber.org/zap"
)

func Initialize(logger *zap.Logger) []Provider {
	return []Provider{
		salesforce{
			logger:      logger,
			transformer: sf.NewSFTransformer(),
		},
		hubSpot{
			logger:      logger,
			transformer: hs.NewHubSpotTransformer(),
		},
	}
}

type Provider interface {
	// Sync the object with the provider.
	SyncProvider(ctx context.Context, obj types.Object)
}

type salesforce struct {
	Provider

	logger      *zap.Logger
	transformer *sf.SFTransformer
}

type hubSpot struct {
	Provider

	logger      *zap.Logger
	transformer *hs.HubSpotTransformer
}

func (s salesforce) SyncProvider(ctx context.Context, obj types.Object) {
	_, err := s.transformer.ToProvider(obj)
	if err != nil {
		s.logger.Error("failed transforming object", zap.Error(err))

		// push to DLQ for retry
		return
	}

	// implement Salesforce-specific sync logic here
}

func (h hubSpot) SyncProvider(ctx context.Context, obj types.Object) {
	_, err := h.transformer.ToProvider(obj)
	if err != nil {
		h.logger.Error("failed transforming object", zap.Error(err))

		// push to DLQ for retry
		return
	}

	// Implement HubSpot-specific sync logic here
}
