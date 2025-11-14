package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mr1cloud/Avito-Reviewer/internal/logger"
	"github.com/mr1cloud/Avito-Reviewer/internal/model"
	"github.com/mr1cloud/Avito-Reviewer/internal/repository"

	"github.com/jmoiron/sqlx"
)

type UsersRepository struct {
	DB     *sqlx.DB
	Logger *logger.Logger
}

func (u *UsersRepository) GetUserById(ctx context.Context, userId string) (*model.User, error) {
	var user model.User

	query := `SELECT * FROM users WHERE user_id = $1`
	err := u.DB.GetContext(ctx, &user, query, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		u.Logger.Errorf("error getting user by id: %v", err)
		return nil, err
	}

	return &user, nil
}

func (u *UsersRepository) UpdateUserIsActiveById(ctx context.Context, userId string, isActive bool) (*model.User, error) {
	var user model.User

	query := `UPDATE users SET is_active = $1 WHERE user_id = $2 RETURNING *`
	err := u.DB.GetContext(ctx, &user, query, isActive, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		u.Logger.Errorf("error updating user is_active by id: %v", err)
		return nil, err
	}

	return &user, nil
}
