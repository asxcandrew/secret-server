package server

import (
	"fmt"
	"time"

	"github.com/asxcandrew/secret-server/storage/model"
	"github.com/go-kit/kit/metrics"
)

func MonitoringMiddleware(
	requestCount metrics.Counter,
	requestLatency metrics.Histogram,
) ServiceMiddleware {
	return func(next SecretService) SecretService {
		return &instrmw{requestCount, requestLatency, next}
	}
}

type instrmw struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	SecretService
}

func (mw *instrmw) Create(m *model.Secret) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.SecretService.Create(m)
	return
}

func (mw *instrmw) Get(h string) (m *model.Secret, rv int, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	m, rv, err = mw.SecretService.Get(h)
	return
}
