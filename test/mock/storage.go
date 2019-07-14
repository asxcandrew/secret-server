package mock

import (
	"github.com/asxcandrew/secret-server/storage/model"
	"github.com/stretchr/testify/mock"
)

type SecretRepositoryMock struct {
	mock.Mock
}

func (r *SecretRepositoryMock) Get(h string) (*model.Secret, error) {
	return &model.Secret{Hash: h}, r.Called().Error(0)
}

func (r *SecretRepositoryMock) Create(m *model.Secret) error {
	return r.Called().Error(0)
}
