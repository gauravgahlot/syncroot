package salesforce

// Contact represents a contact object in Salesforce.
type Contact struct {
	// ContactID is the unique identifier for the contact.
	ContactID string `json:"contact_id"`

	// Name is the name of the contact.
	Name Name `json:"name"`

	// ContactEmail is the email address of the contact.
	ContactEmail string `json:"contact_email"`

	// PhoneNumber is the phone number of the contact.
	PhoneNumber string `json:"phone_number"`

	// CreatedAt is the timestamp when the contact was created.
	CreatedAt string `json:"createdAt"`

	// UpdatedAt is the timestamp when the contact was last updated.
	UpdatedAt string `json:"updatedAt"`
}

type Name struct {
	First string `json:"first_name"`
	Last  string `json:"last_name"`
}
