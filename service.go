package server

import (
	"errors"
	"time"

	"github.com/asxcandrew/secret-server/storage"
	"github.com/asxcandrew/secret-server/storage/model"
	uuid "github.com/satori/go.uuid"
)

type SecretService interface {
	Get(string) (*model.Secret, int, error)
	Create(*model.Secret) error
}

type secretService struct {
	storage storage.Storage
}

// NewSecretService creates the secret service with necessary dependencies.
func NewSecretService(storage storage.Storage) SecretService {
	return &secretService{
		storage: storage,
	}
}

func (s *secretService) Get(hash string) (*model.Secret, int, error) {
	m, err := s.storage.Secret.Get(hash)

	if err != nil {
		return nil, 0, err
	}

	if time.Now().After(m.ExpiresAt) {
		return nil, 0, errors.New("Secret is expired")
	}

	rv, err := s.storage.Secret.RemainingViews(m)

	if err != nil {
		return nil, 0, err
	}

	if rv < 1 {
		return nil, 0, errors.New("Secret has no views")
	}

	err = s.storage.Secret.CommitView(m)

	if err != nil {
		return nil, 0, err
	}
	return m, rv - 1, nil
}

func (s *secretService) Create(m *model.Secret) error {
	u := uuid.NewV4()

	m.Hash = u.String()
	return s.storage.Secret.Create(m)
}
