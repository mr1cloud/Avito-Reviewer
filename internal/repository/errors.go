package repository

import "errors"

var (
	ErrNotFound      = errors.New("record not found")
	ErrEmptyResult   = errors.New("no records found")
	ErrConflict      = errors.New("unique constraint violated")
	ErrAlreadyExists = errors.New("record already exists")
	ErrInvalidInput  = errors.New("invalid input")
)
