package server

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/asxcandrew/secret-server/middleware"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func MakeSecretHandler(s SecretService, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
		kithttp.ServerBefore(middleware.HTTPToContext()),
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
		return nil, InvalidInputError
	}

	decoder := schema.NewDecoder()

	err := decoder.Decode(body, r.PostForm)

	if err != nil {
		return nil, InvalidInputError
	}

	err = body.Validate()

	if err != nil {
		return nil, InvalidInputError
	}

	return body, nil
}

func decodeGetSecretRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	val, ok := vars["hash"]

	if !ok {
		return nil, NotFoundError
	}

	return &GetSecretRequest{Hash: val}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if val, ok := ctx.Value(middleware.AcceptHeaderJSONContextKey).(bool); ok && val {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		err = json.NewEncoder(w).Encode(response)
	}
	if val, ok := ctx.Value(middleware.AcceptHeaderXMLContextKey).(bool); ok && val {
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
		err = xml.NewEncoder(w).Encode(response)
	}
	return err
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch err {
	case NotFoundError:
		w.WriteHeader(http.StatusNotFound)
	case InvalidInputError:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
