package error

import "net/http"

type ServiceError interface {
	error
	ErrorStatusCode() int
	Code() string
}

// NotFoundError indicates that a resource was not found.
type notFoundError struct{}

// ErrorStatusCode returns the HTTP status code for the error.
func (e *notFoundError) ErrorStatusCode() int { return http.StatusNotFound }

// Code returns the error code.
func (e *notFoundError) Code() string { return "NOT_FOUND" }

// Error returns the error message.
func (e *notFoundError) Error() string { return "resource not found" }

func NewNotFoundError() ServiceError {
	return &notFoundError{}
}
