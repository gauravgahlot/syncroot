package salesforce

import (
	"errors"
	"strings"

	"github.com/gauravgahlot/syncroot/internal/types"
)

type SFTransformer struct{}

func (t SFTransformer) ToProvider(input interface{}) (interface{}, error) {
	contact, ok := input.(*types.Contact)
	if !ok {
		return nil, errors.New("invalid type provided")
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

func (t SFTransformer) FromProvider(input interface{}) (interface{}, error) {
	sfContact, ok := input.(*Contact)
	if !ok {
		return nil, errors.New("invalid type provided")
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
