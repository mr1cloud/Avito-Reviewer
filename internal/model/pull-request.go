package model

import (
	"github.com/mr1cloud/Avito-Reviewer/internal/validation"

	"github.com/go-playground/validator/v10"
)

//goland:noinspection GoUnhandledErrorResult
func init() {
	validation.Validate.RegisterValidation("pr_status", ValidatePrStatus)
}

type PrStatus string

const (
	PrStatusOpen   PrStatus = "open"
	PrStatusMerged PrStatus = "merged"
)

func (pr PrStatus) IsValid() bool {
	switch pr {
	case PrStatusOpen, PrStatusMerged:
		return true
	}
	return false
}

func ValidatePrStatus(fl validator.FieldLevel) bool {
	return PrStatus(fl.Field().String()).IsValid()
}

type PullRequest struct {
	PullRequestID   string   `json:"pull_request_id" db:"pull_request_id"`
	PullRequestName string   `json:"pull_request_name" db:"pull_request_name"`
	AuthorID        string   `json:"author_id" db:"author_id"`
	Status          PrStatus `json:"status" db:"status"`
}

type PullRequestShort struct {
	PullRequestID   string   `json:"pull_request_id" db:"pull_request_id"`
	PullRequestName string   `json:"pull_request_name" db:"pull_request_name"`
	AuthorID        string   `json:"author_id" db:"author_id"`
	Status          PrStatus `json:"status" db:"status"`
}
