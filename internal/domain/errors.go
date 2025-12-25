package domain

import "errors"

var (
	ErrInvalidRequest         = errors.New("invalid request")
	ErrTaskExists             = errors.New("task already exists")
	ErrNotFound               = errors.New("resource not found")
	ErrInternalError          = errors.New("internal server error")
	ErrFailedToDecodeJSON     = errors.New("failed to decode JSON")
	ErrQueryParameterRequired = errors.New("query parameter is required")
)
