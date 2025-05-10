package hubspot

import (
	"errors"
	"strings"

	"github.com/gauravgahlot/syncroot/internal/types"
)

// Contact represents a contact object in HubSpot.
type Contact struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// contactTf handles transformation logic for Contact.
type contactTf struct{}

func (t contactTf) toProvider(input types.Object) (interface{}, error) {
	contact, ok := input.(*types.Contact)
	if !ok {
		return nil, errors.New("invalid type provided to ContactTransformer")
	}

	parts := strings.Fields(contact.FullName)
	first, last := "", ""
	if len(parts) > 0 {
		first = parts[0]
	}
	if len(parts) > 1 {
		last = strings.Join(parts[1:], " ")
	}

	return &Contact{
		ID:          contact.ID,
		Email:       contact.Email,
		PhoneNumber: contact.Phone,
		CreatedAt:   contact.CreatedAt,
		UpdatedAt:   contact.UpdatedAt,
		FirstName:   first,
		LastName:    last,
	}, nil
}

func (t contactTf) fromProvider(input interface{}) (types.Object, error) {
	sfContact, ok := input.(*Contact)
	if !ok {
		return nil, errors.New("invalid type provided to ContactTransformer")
	}

	fullName := strings.TrimSpace(sfContact.FirstName + " " + sfContact.LastName)

	return &types.Contact{
		ID:        sfContact.ID,
		FullName:  fullName,
		Email:     sfContact.Email,
		Phone:     sfContact.PhoneNumber,
		CreatedAt: sfContact.CreatedAt,
		UpdatedAt: sfContact.UpdatedAt,
	}, nil
}
