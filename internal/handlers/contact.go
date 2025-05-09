package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type contactHandler struct {
	logger *zap.Logger
}

func NewContactHandler(logger *zap.Logger) *contactHandler {
	return &contactHandler{
		logger: logger,
	}
}

// CreateContact handles POST /contacts requests.
func (h *contactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("creating contact")

	var contact map[string]string
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		h.logger.Error("failed to decode request body", zap.Error(err))
		http.Error(w, "invalid request", http.StatusBadRequest)

		return
	}

	// simulate saving contact to a database
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Contact created successfully"))
}

// GetContact handles GET /contacts/{id} requests.
func (h *contactHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	contactID := r.PathValue("id")
	h.logger.Info("retrieving contact", zap.String("id", contactID))

	// simulate fetching contact from a database
	contact := map[string]string{
		"id":    contactID,
		"name":  "John Doe",
		"email": "john@doe.io",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
	w.WriteHeader(http.StatusOK)
}

// UpdateContact handles PUT /contacts/{id} requests.
func (h *contactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	contactID := r.PathValue("id")
	h.logger.Info("updating contact", zap.String("id", contactID))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Contact updated successfully"))
}

// DeleteContact handles DELETE /contacts/{id} requests.
func (h *contactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	contactID := r.PathValue("id")
	h.logger.Info("deleting contact", zap.String("id", contactID))

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Contact deleted successfully"))
}
