package storage

import (
	"github.com/asxcandrew/secret-server/storage/repository"
	pgrepository "github.com/asxcandrew/secret-server/storage/repository/pg"
	"github.com/go-pg/pg"
)

type Storage struct {
	Secret repository.SecretRepository
}

func NewPGStorage(db *pg.DB) Storage {
	return Storage{
		Secret: pgrepository.NewPGSecretRepository(db),
	}
}
