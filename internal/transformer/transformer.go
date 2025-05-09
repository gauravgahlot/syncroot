package transformer

import (
	"github.com/gauravgahlot/syncroot/internal/transformer/hubspot"
	"github.com/gauravgahlot/syncroot/internal/transformer/salesforce"
)

type Transformer interface {
	// ToProvider converts the internal data to the provider's format.
	ToProvider(input interface{}) (interface{}, error)

	// FromProvider converts the provider's data to the internal format.
	FromProvider(input interface{}) (interface{}, error)
}

// Ensure that transformers for different providers implement
// the Transformer interface.
var (
	_ Transformer = &salesforce.SFTransformer{}
	_ Transformer = &hubspot.HubspotTransformer{}
)
