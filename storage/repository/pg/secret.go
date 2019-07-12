package pg

import (
	"github.com/asxcandrew/secret-server/storage/model"
	"github.com/go-pg/pg"
)

type SecretRepository struct {
	db *pg.DB
}

func NewPGSecretRepository(db *pg.DB) *SecretRepository {
	return &SecretRepository{
		db: db,
	}
}

func (r *SecretRepository) Get(h string) (*model.Secret, error) {
	m := &model.Secret{Hash: h}
	err := r.db.Select(m)

	return m, err
}
