package test

import (
	"github.com/asxcandrew/secret-server/storage"
	"github.com/asxcandrew/secret-server/storage/model"
	"github.com/stretchr/testify/mock"
)

func NewStorageMock() storage.Storage {
	return storage.Storage{
		Secret: &secretRepositoryMock{},
	}
}

type secretRepositoryMock struct {
	mock.Mock
}

func (r *secretRepositoryMock) Get(h string) (*model.Secret, error) {
	return &model.Secret{Hash: h}, r.Called().Error(0)
}

func (r *secretRepositoryMock) Create(m *model.Secret) error {
	return r.Called().Error(0)
}
