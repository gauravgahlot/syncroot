package types

// CRMProvider represents a CRM provider configuration.
type CRMProvider struct {
	// Name of the CRM provider.
	Name string `json:"name"`

	// WebhookSecret is the secret used for webhook verification.
	WebhookSecret string `json:"webhookSecret"`

	// Headers expected in the webhook request.
	Hearders []string `json:"hearders"`

	// Objects that the provider can handle.
	Objects []Object `json:"objects"`
}

// Object represents a CRM object that can be handled by the provider.
type Object struct {
	// Type of the object (e.g., "contact", "deal").
	Type string `json:"type"`

	// Event that triggers the webhook (e.g., "created", "updated").
	Event string `json:"event"`

	// Properties that are expected in the webhook payload.
	Properties []string `json:"properties"`
}
