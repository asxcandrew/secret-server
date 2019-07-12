package server

import (
	"fmt"
	"time"

	"github.com/asxcandrew/galas/api/representation"
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

func (s *itemLoggingService) Get(id int) (item *model.Secret, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "get",
			"params", fmt.Sprintf("[id=%d]", id),
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.SecretService.Get(id)
}

func (s *itemLoggingService) Create(item *representation.ItemEntity, authorID int) (res *model.Secret, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "create",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.SecretService.Create(item, authorID)
}
