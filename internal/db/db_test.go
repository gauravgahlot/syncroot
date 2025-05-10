package db_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gauravgahlot/syncroot/internal/db"
	"github.com/gauravgahlot/syncroot/internal/types"
)

func TestInMemoryStore(t *testing.T) {
	ctx := context.Background()
	store := db.NewInMemoryStore()

	t.Run("CreateContact", func(t *testing.T) {
		contact := &types.Contact{
			ID:       "3fa85f64-5717-4562-b3fc-2c963f66afa6",
			FullName: "John Doe",
			Email:    "john@example.com",
			Phone:    "123-456-7890",
		}

		created, err := store.CreateContact(ctx, contact)
		assert.NoError(t, err)
		assert.Equal(t, contact, created)

		_, err = store.CreateContact(ctx, contact)
		assert.Error(t, err)

		invalidContact := &types.Contact{
			FullName: "Invalid Contact",
		}
		_, err = store.CreateContact(ctx, invalidContact)
		assert.Error(t, err)
	})

	t.Run("GetContact", func(t *testing.T) {
		contact, err := store.GetContact(ctx, "contact1")
		assert.NoError(t, err)
		assert.Equal(t, "John Doe", contact.FullName)

		_, err = store.GetContact(ctx, "nonexistent")
		assert.Error(t, err)
		assert.Equal(t, db.ErrContactNotFound, err)
	})

	t.Run("ListContacts", func(t *testing.T) {
		contact2 := &types.Contact{
			ID:       "3fa85f64-5717-4562-b3fc-2c963f66afa6",
			FullName: "Jane Smith",
			Email:    "jane@example.com",
			Phone:    "987-654-3210",
		}
		_, err := store.CreateContact(ctx, contact2)
		assert.NoError(t, err)

		contacts, err := store.ListContacts(ctx)
		assert.NoError(t, err)
		assert.Len(t, contacts, 2)
	})

	t.Run("UpdateContact", func(t *testing.T) {
		contact, err := store.GetContact(ctx, "contact1")
		assert.NoError(t, err)

		contact.FullName = "John Updated"
		updated, err := store.UpdateContact(ctx, contact)
		assert.NoError(t, err)
		assert.Equal(t, "John Updated", updated.FullName)

		contact, err = store.GetContact(ctx, "3fa85f64-5717-4562-b3fc-2c963f66afa6")
		assert.NoError(t, err)
		assert.Equal(t, "John Updated", contact.FullName)

		nonexistent := &types.Contact{
			ID:       "nonexistent",
			FullName: "Nonexistent Contact",
		}
		_, err = store.UpdateContact(ctx, nonexistent)
		assert.Error(t, err)
		assert.Equal(t, db.ErrContactNotFound, err)

		invalidContact := &types.Contact{
			FullName: "Invalid Contact",
		}
		_, err = store.UpdateContact(ctx, invalidContact)
		assert.Error(t, err)
	})

	t.Run("DeleteContact", func(t *testing.T) {
		err := store.DeleteContact(ctx, "3fa85f64-5717-4562-b3fc-2c963f66afa6")
		assert.NoError(t, err)

		_, err = store.GetContact(ctx, "3fa85f64-5717-4562-b3fc-2c963f66afa6")
		assert.Error(t, err)
		assert.Equal(t, db.ErrContactNotFound, err)

		err = store.DeleteContact(ctx, "nonexistent")
		assert.Error(t, err)
		assert.Equal(t, db.ErrContactNotFound, err)
	})
}
