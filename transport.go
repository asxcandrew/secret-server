package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func MakeSecretHandler(s SecretService, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	createSecretHandler := kithttp.NewServer(
		MakeCreateSecretEndpoint(s),
		decodeCreateSecretRequest,
		encodeResponse,
		opts...,
	)

	getSecretHandler := kithttp.NewServer(
		MakeGetSecretEndpoint(s),
		decodeGetSecretRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/secret/{hash}", getSecretHandler).Methods("GET")
	r.Handle("/secret", createSecretHandler).Methods("POST")

	return r
}

func decodeCreateSecretRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body = &CreateSecretRequest{}

	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	decoder := schema.NewDecoder()

	err := decoder.Decode(body, r.PostForm)

	if err != nil {
		return nil, err
	}

	err = body.Validate()

	if err != nil {
		return nil, err
	}

	return body, nil
}

func decodeGetSecretRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	val, ok := vars["hash"]

	if !ok {
		return nil, errors.New("Bad request")
	}

	return GetSecretRequest{Hash: val}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
}
