package server

import (
	"context"
	"errors"
	"time"

	"github.com/asxcandrew/secret-server/storage/model"

	"github.com/go-kit/kit/endpoint"
	validation "github.com/go-ozzo/ozzo-validation"
)

type GetSecretRequest struct {
	Hash string
}

type CreateSecretRequest struct {
	Secret           string `schema:"secret,required"`
	ExpireAfterViews int    `schema:"expireAfterViews,required"`
	ExpireAfter      int    `schema:"expireAfter,required"`
}

func (r *CreateSecretRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.ExpireAfter, validation.Min(0)),
		validation.Field(&r.ExpireAfterViews, validation.Min(1)),
	)
}

func MakeGetSecretEndpoint(s SecretService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetSecretRequest)

		secret, err := s.Get(req.Hash)

		if err != nil {
			return nil, err
		}
		if time.Now().After(secret.ExpiresAt) {
			return nil, errors.New("Expired")
		}

		res := modelToSecretEntity(secret)

		return res, nil
	}
}

func MakeCreateSecretEndpoint(s SecretService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*CreateSecretRequest)

		secret := &model.Secret{
			Body:       req.Secret,
			ViewsLimit: req.ExpireAfterViews,
			ExpiresAt:  time.Now().Local().Add(time.Minute * time.Duration(req.ExpireAfter)),
		}

		err := s.Create(secret)

		if err != nil {
			return nil, err
		}

		res := modelToSecretEntity(secret)
		res.RemainingViews = req.ExpireAfterViews

		return res, nil
	}
}
