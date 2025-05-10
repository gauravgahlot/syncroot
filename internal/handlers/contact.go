package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gauravgahlot/syncroot/internal/enqueuer"
	"github.com/gauravgahlot/syncroot/internal/types"
	"go.uber.org/zap"
)

type contactHandler struct {
	logger   *zap.Logger
	enqueuer enqueuer.Enqueuer
	topic    string
}

func NewContactHandler(logger *zap.Logger, eq enqueuer.Enqueuer, topic string) *contactHandler {
	return &contactHandler{
		logger:   logger,
		enqueuer: eq,
		topic:    topic,
	}
}

// CreateContact handles POST /contacts requests.
func (h *contactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("creating contact")

	var contact types.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		h.logger.Error("failed to decode request body", zap.Error(err))
		http.Error(w, "invalid request", http.StatusBadRequest)

		return
	}

	// enqueue contact for processing
	err := h.enqueuer.Enqueue(enqueuer.EnqueueRequest{
		Operation: types.OperationCreate,
		Object:    &contact,
		Topic:     h.topic,
	})
	if err != nil {
		// handle error
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
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

	// enqueue contact for processing
	err := h.enqueuer.Enqueue(enqueuer.EnqueueRequest{
		Operation: types.OperationUpdate,
		Object: &types.Contact{
			ID: contactID,
		},
		Topic: h.topic,
	})
	if err != nil {
		// handle error
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteContact handles DELETE /contacts/{id} requests.
func (h *contactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	contactID := r.PathValue("id")

	// enqueue contact for processing
	err := h.enqueuer.Enqueue(enqueuer.EnqueueRequest{
		Operation: types.OperationDelete,
		Object: &types.Contact{
			ID: contactID,
		},
		Topic: h.topic,
	})
	if err != nil {
		// handle error
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
