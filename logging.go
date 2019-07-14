package server

import (
	"fmt"
	"time"

	"github.com/asxcandrew/secret-server/storage/model"
	"github.com/go-kit/kit/log"
)

type secretLoggingService struct {
	logger log.Logger
	SecretService
}

// NewSecretLoggingService returns a new instance of a secretLoggingService.
func NewSecretLoggingService(logger log.Logger, s SecretService) SecretService {
	logger = log.With(logger, "service", "secret")

	return &secretLoggingService{logger, s}
}

func (s *secretLoggingService) Get(h string) (item *model.Secret, rv int, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "get",
			"params", fmt.Sprintf("[hash=%s]", h),
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.SecretService.Get(h)
}

func (s *secretLoggingService) Create(m *model.Secret) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "create",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.SecretService.Create(m)
}
