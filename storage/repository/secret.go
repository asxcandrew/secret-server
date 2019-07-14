package repository

import (
	"github.com/asxcandrew/secret-server/storage/model"
)

type SecretRepository interface {
	Get(string) (*model.Secret, error)
	Create(*model.Secret) error
	CommitView(*model.Secret) error
	RemainingViews(*model.Secret) (int, error)
}
