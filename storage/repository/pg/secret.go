package pg

import (
	"time"

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
	m := &model.Secret{}
	err := r.db.Model(m).Where("hash = ?", h).Select()

	return m, err
}

func (r *SecretRepository) Create(m *model.Secret) error {
	m.CreatedAt = time.Now()

	return r.db.Insert(m)
}
