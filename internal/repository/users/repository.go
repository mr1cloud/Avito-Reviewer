package users

import (
	"context"

	"github.com/mr1cloud/Avito-Reviewer/internal/model"
)

// UsersRepository stores users data
//
//goland:noinspection GoNameStartsWithPackageName
type UsersRepository interface {
	// GetUserById retrieves a user by userId
	GetUserById(ctx context.Context, userId string) (*model.User, error)
	// UpdateUserIsActiveById updates the isActive status of a user by userId
	UpdateUserIsActiveById(ctx context.Context, userId string, isActive bool) (*model.User, error)
}
