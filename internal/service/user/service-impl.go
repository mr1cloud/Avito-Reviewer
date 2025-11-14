package user

import (
	"context"
	"errors"

	"github.com/mr1cloud/Avito-Reviewer/internal/logger"
	"github.com/mr1cloud/Avito-Reviewer/internal/model"
	"github.com/mr1cloud/Avito-Reviewer/internal/repository"
	"github.com/mr1cloud/Avito-Reviewer/internal/repository/users"

	serviceerrors "github.com/mr1cloud/Avito-Reviewer/internal/error"
)

type service struct {
	usersRepository users.UsersRepository
	logger          *logger.Logger
}

func (s *service) GetUserById(ctx context.Context, userId string) (*model.User, error) {
	user, err := s.usersRepository.GetUserById(ctx, userId)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, serviceerrors.NewNotFoundError()
		default:
			return nil, err
		}
	}
	return user, nil
}

func (s *service) UpdateUserIsActiveById(ctx context.Context, userId string, isActive bool) (*model.User, error) {
	user, err := s.usersRepository.UpdateUserIsActiveById(ctx, userId, isActive)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, serviceerrors.NewNotFoundError()
		default:
			return nil, err
		}
	}
	return user, nil
}

func NewService(logger *logger.Logger, usersRepo users.UsersRepository) User {
	return &service{
		usersRepository: usersRepo,
		logger:          logger,
	}
}
