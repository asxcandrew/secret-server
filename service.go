package server

import (
	"github.com/asxcandrew/secret-server/storage"
	"github.com/asxcandrew/secret-server/storage/model"
	uuid "github.com/satori/go.uuid"
)

type SecretService interface {
	Get(string) (*model.Secret, error)
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

func (s *secretService) Get(hash string) (*model.Secret, error) {
	return s.storage.Secret.Get(hash)
}

func (s *secretService) Create(m *model.Secret) error {
	u := uuid.NewV4()

	m.Hash = u.String()
	return s.storage.Secret.Create(m)
}
