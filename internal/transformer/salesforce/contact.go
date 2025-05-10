package salesforce

import (
	"errors"
	"strings"

	"github.com/gauravgahlot/syncroot/internal/types"
)

// Contact represents a contact object in Salesforce.
type Contact struct {
	ContactID    string `json:"contact_id"`
	Name         Name   `json:"name"`
	ContactEmail string `json:"contact_email"`
	PhoneNumber  string `json:"phone_number"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type Name struct {
	First string `json:"first_name"`
	Last  string `json:"last_name"`
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
		ContactID:    contact.ID,
		ContactEmail: contact.Email,
		PhoneNumber:  contact.Phone,
		CreatedAt:    contact.CreatedAt,
		UpdatedAt:    contact.UpdatedAt,
		Name: Name{
			First: first,
			Last:  last,
		},
	}, nil
}

func (t contactTf) fromProvider(input interface{}) (types.Object, error) {
	sfContact, ok := input.(*Contact)
	if !ok {
		return nil, errors.New("invalid type provided to ContactTransformer")
	}

	fullName := strings.TrimSpace(sfContact.Name.First + " " + sfContact.Name.Last)

	return &types.Contact{
		ID:        sfContact.ContactID,
		FullName:  fullName,
		Email:     sfContact.ContactEmail,
		Phone:     sfContact.PhoneNumber,
		CreatedAt: sfContact.CreatedAt,
		UpdatedAt: sfContact.UpdatedAt,
	}, nil
}
