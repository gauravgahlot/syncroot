package db

import (
	"context"
	"errors"

	"github.com/gauravgahlot/syncroot/internal/types"
)

// Store provides an API to perform CRUD operations on contacts.
type Store interface {
	CreateContact(ctx context.Context, contact *types.Contact) (*types.Contact, error)
	GetContact(ctx context.Context, id string) (*types.Contact, error)
	ListContacts(ctx context.Context) ([]*types.Contact, error)
	UpdateContact(ctx context.Context, contact *types.Contact) (*types.Contact, error)
	DeleteContact(ctx context.Context, id string) error
}

// ErrContactNotFound is returned when a contact is not found
var ErrContactNotFound = errors.New("contact not found")

// inMemoryStore for testing purposes.
type inMemoryStore struct {
	contacts map[string]*types.Contact
}

func NewInMemoryStore() *inMemoryStore {
	return &inMemoryStore{
		contacts: make(map[string]*types.Contact),
	}
}

func (s *inMemoryStore) CreateContact(ctx context.Context, contact *types.Contact) (*types.Contact, error) {
	if contact.ID == "" {
		return nil, errors.New("contact ID cannot be empty")
	}

	if _, exists := s.contacts[contact.ID]; exists {
		return nil, errors.New("contact already exists")
	}

	s.contacts[contact.ID] = contact

	return contact, nil
}

func (s *inMemoryStore) GetContact(ctx context.Context, id string) (*types.Contact, error) {
	contact, exists := s.contacts[id]
	if !exists {
		return nil, ErrContactNotFound
	}

	return contact, nil
}

func (s *inMemoryStore) ListContacts(ctx context.Context) ([]*types.Contact, error) {
	contacts := make([]*types.Contact, 0, len(s.contacts))

	for _, contact := range s.contacts {
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (s *inMemoryStore) UpdateContact(ctx context.Context, contact *types.Contact) (*types.Contact, error) {
	if contact.ID == "" {
		return nil, errors.New("contact ID cannot be empty")
	}

	if _, exists := s.contacts[contact.ID]; !exists {
		return nil, ErrContactNotFound
	}

	s.contacts[contact.ID] = contact

	return contact, nil
}

func (s *inMemoryStore) DeleteContact(ctx context.Context, id string) error {
	if _, exists := s.contacts[id]; !exists {
		return ErrContactNotFound
	}

	delete(s.contacts, id)

	return nil
}
