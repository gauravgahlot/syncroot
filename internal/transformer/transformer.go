package transformer

import (
	"github.com/gauravgahlot/syncroot/internal/transformer/hubspot"
	"github.com/gauravgahlot/syncroot/internal/transformer/salesforce"
	"github.com/gauravgahlot/syncroot/internal/types"
)

type Transformer interface {
	// ToProvider converts the internal data to the provider's format.
	ToProvider(input types.Object) (interface{}, error)

	// FromProvider converts the provider's data to the internal format.
	FromProvider(input interface{}) (types.Object, error)
}

// Ensure that transformers for different providers implement
// the Transformer interface.
var (
	_ Transformer = &salesforce.SFTransformer{}
	_ Transformer = &hubspot.HubSpotTransformer{}
)
