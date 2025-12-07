package models

import "time"

type Secret struct {
	ID           string    `json:"id"`
	EncryptedData string  `json:"encryptedData"`
	CreatedAt    time.Time `json:"createdAt"`
}

type CreateSecretRequest struct {
	EncryptedData string `json:"encryptedData"`
}

type CreateSecretResponse struct {
	ID string `json:"id"`
}

type GetSecretResponse struct {
	EncryptedData string `json:"encryptedData"`
}



