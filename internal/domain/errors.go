package domain

import "errors"

var (
	ErrInvalidRequest     = errors.New("invalid request")
	ErrNotFound           = errors.New("resource not found")
	ErrInternalError      = errors.New("internal server error")
	ErrFailedToDecodeJSON = errors.New("failed to decode JSON")
	ErrMethodNotAllowed   = errors.New("method not allowed")
)
