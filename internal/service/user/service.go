package user

import (
	"context"

	"github.com/mr1cloud/Avito-Reviewer/internal/model"
)

type User interface {
	// GetUserById retrieves a user by their unique identifier
	GetUserById(ctx context.Context, userId string) (*model.User, error)
	// UpdateUserIsActiveById updates the isActive status of a user by their unique identifier
	UpdateUserIsActiveById(ctx context.Context, userId string, isActive bool) (*model.User, error)
}
