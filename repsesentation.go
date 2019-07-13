package server

import (
	"encoding/xml"
	"time"

	"github.com/asxcandrew/secret-server/storage/model"
)

type SecretEntity struct {
	XMLName        xml.Name  `xml:"Secret"`
	Hash           string    `json:"hash" xml:"hash"`
	SecretText     string    `json:"secretText" xml:"secretText"`
	CreatedAt      time.Time `json:"createdAt" xml:"createdAt"`
	ExpiresAt      time.Time `json:"expiresAt" xml:"expiresAt"`
	RemainingViews int       `json:"remainingViews" xml:"remainingViews"`
}

func modelToSecretEntity(m *model.Secret) *SecretEntity {
	return &SecretEntity{
		Hash:       m.Hash,
		SecretText: m.Body,
		CreatedAt:  m.CreatedAt,
		ExpiresAt:  m.ExpiresAt,
	}
}
