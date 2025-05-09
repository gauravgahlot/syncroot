package hubspot

// Contact represents a contact object in HubSpot.
type Contact struct {
	// ID is the unique identifier for the contact.
	ID string `json:"id"`

	// FirstName of the contact.
	FirstName string `json:"first_name"`

	// LastName of the contact.
	LastName string `json:"last_name"`

	// Email is the email address of the contact.
	Email string `json:"email"`

	// PhoneNumber is the phone number of the contact.
	PhoneNumber string `json:"phone_number"`

	// CreatedAt is the timestamp when the contact was created.
	CreatedAt string `json:"createdAt"`

	// UpdatedAt is the timestamp when the contact was last updated.
	UpdatedAt string `json:"updatedAt"`
}
