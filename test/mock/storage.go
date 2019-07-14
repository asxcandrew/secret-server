package mock

import (
	"time"

	"github.com/asxcandrew/secret-server/storage/model"
	"github.com/stretchr/testify/mock"
)

type SecretRepositoryMock struct {
	mock.Mock
}

func (r *SecretRepositoryMock) Get(h string) (*model.Secret, error) {
	return &model.Secret{
		Hash:      h,
		ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(1)),
	}, r.Called().Error(0)
}

func (r *SecretRepositoryMock) Create(m *model.Secret) error {
	return r.Called().Error(0)
}

func (r *SecretRepositoryMock) CommitView(m *model.Secret) error {
	return r.Called().Error(0)
}

func (r *SecretRepositoryMock) RemainingViews(m *model.Secret) (int, error) {
	return r.Called().Int(0), r.Called().Error(1)
}
