package server

import (
	"fmt"
	"time"

	"github.com/asxcandrew/secret-server/storage/model"
	"github.com/go-kit/kit/log"
)

type lggmdwr struct {
	logger log.Logger
	SecretService
}

func LogginggMiddleware(
	logger log.Logger,
) ServiceMiddleware {
	return func(next SecretService) SecretService {
		logger = log.With(logger, "service", "secret")
		return &lggmdwr{logger, next}
	}
}

func (s *lggmdwr) Get(h string) (item *model.Secret, rv int, err error) {
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

func (s *lggmdwr) Create(m *model.Secret) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "create",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.SecretService.Create(m)
}
