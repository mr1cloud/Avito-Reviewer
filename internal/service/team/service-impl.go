package team

import (
	"context"
	"errors"

	"github.com/mr1cloud/Avito-Reviewer/internal/logger"
	"github.com/mr1cloud/Avito-Reviewer/internal/model"
	"github.com/mr1cloud/Avito-Reviewer/internal/repository"
	"github.com/mr1cloud/Avito-Reviewer/internal/repository/teams"

	serviceerrors "github.com/mr1cloud/Avito-Reviewer/internal/error"
)

type service struct {
	teamsRepository teams.TeamsRepository
	logger          *logger.Logger
}

func (s *service) CreateTeam(ctx context.Context, teamName string, members model.TeamMembers) (*model.Team, error) {
	err := s.teamsRepository.CreateTeam(ctx, teamName, members)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrConflict):
			return nil, NewTeamAlreadyExistsError(teamName)
		default:
			return nil, err
		}
	}

	return s.GetTeam(ctx, teamName)
}

func (s *service) UpdateTeam(ctx context.Context, teamName string, members model.TeamMembers) (*model.Team, error) {
	team, err := s.GetTeam(ctx, teamName)
	if err != nil {
		return nil, err
	}

	err = s.teamsRepository.UpdateTeam(ctx, teamName, team.Members, members)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, serviceerrors.NewNotFoundError()
		default:
			return nil, err
		}
	}

	return s.GetTeam(ctx, teamName)
}

func (s *service) GetTeam(ctx context.Context, teamName string) (*model.Team, error) {
	team, err := s.teamsRepository.GetTeam(ctx, teamName)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, serviceerrors.NewNotFoundError()
		default:
			return nil, err
		}
	}
	return team, nil
}

func NewService(logger *logger.Logger, teamsRepository teams.TeamsRepository) Team {
	return &service{
		teamsRepository: teamsRepository,
		logger:          logger,
	}
}
