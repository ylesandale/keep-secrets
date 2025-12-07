package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"keep-secrets/server/models"
	"keep-secrets/server/storage"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type SecretHandler struct {
	db *storage.DB
}

func NewSecretHandler(db *storage.DB) *SecretHandler {
	return &SecretHandler{db: db}
}

func (h *SecretHandler) CreateSecret(w http.ResponseWriter, r *http.Request) {
	var req models.CreateSecretRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.EncryptedData) == "" {
		http.Error(w, "Encrypted data is required", http.StatusBadRequest)
		return
	}

	if len(req.EncryptedData) > 10*1024*1024 {
		http.Error(w, "Secret too large (max 10MB)", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()

	if err := h.db.SaveSecret(id, req.EncryptedData); err != nil {
		http.Error(w, "Failed to save secret", http.StatusInternalServerError)
		return
	}

	response := models.CreateSecretResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *SecretHandler) GetSecret(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "Secret ID is required", http.StatusBadRequest)
		return
	}

	encryptedData, err := h.db.GetAndDeleteSecret(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Secret not found or already viewed", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve secret", http.StatusInternalServerError)
		return
	}

	response := models.GetSecretResponse{EncryptedData: encryptedData}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

