package server

import "errors"

var (
	NotFoundError     = errors.New("Secret not found")
	InvalidInputError = errors.New("Invalid input")
)
