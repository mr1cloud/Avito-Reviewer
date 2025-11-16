package pull_request

import (
	"context"
	"errors"

	"github.com/mr1cloud/Avito-Reviewer/internal/logger"
	"github.com/mr1cloud/Avito-Reviewer/internal/model"
	"github.com/mr1cloud/Avito-Reviewer/internal/repository"
	"github.com/mr1cloud/Avito-Reviewer/internal/repository/pull-requests"
	"github.com/mr1cloud/Avito-Reviewer/internal/service/team"
	"github.com/mr1cloud/Avito-Reviewer/internal/service/user"

	serviceerrors "github.com/mr1cloud/Avito-Reviewer/internal/error"
)

type service struct {
	pullRequestsRepository pull_requests.PullRequestsRepository
	teams                  team.Team
	users                  user.User
	logger                 *logger.Logger
}

func (s *service) CreatePullRequest(ctx context.Context, pullRequestId, pullRequestName, authorId string) (*model.PullRequest, error) {
	user, err := s.users.GetUserById(ctx, authorId)
	if err != nil {
		return nil, err
	}

	team, err := s.teams.GetTeam(ctx, user.TeamName)
	if err != nil {
		return nil, err
	}

	assignedReviewers := team.Members.GetActiveMembers(2, authorId)
	if len(assignedReviewers) == 0 {
		return nil, NewNoCandidateReviewersError()
	}

	err = s.pullRequestsRepository.InsertPullRequest(ctx, pullRequestId, pullRequestName, authorId, assignedReviewers)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrAlreadyExists):
			return nil, NewPullRequestAlreadyExistsError(pullRequestId)
		default:
			return nil, err
		}
	}

	return s.GetPullRequest(ctx, pullRequestId)
}

func (s *service) GetPullRequest(ctx context.Context, pullRequestId string) (*model.PullRequest, error) {
	pullRequest, err := s.pullRequestsRepository.SelectPullRequestById(ctx, pullRequestId)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, serviceerrors.NewNotFoundError()
		default:
			return nil, err
		}
	}

	return pullRequest, nil
}

func (s *service) MergePullRequest(ctx context.Context, pullRequestId string) (*model.PullRequest, error) {
	return nil, nil
}

func (s *service) GetPullRequestsAssignedToUser(ctx context.Context, userId string) ([]model.PullRequestShort, error) {
	_, err := s.users.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	pullRequests, err := s.pullRequestsRepository.SelectPullRequestsAssignedToUser(ctx, userId)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, serviceerrors.NewNotFoundError()
		default:
			return nil, err
		}
	}

	return pullRequests, nil
}

func NewService(logger *logger.Logger, pullRequestsRepository pull_requests.PullRequestsRepository, users user.User, teams team.Team) PullRequest {
	return &service{
		pullRequestsRepository: pullRequestsRepository,
		teams:                  teams,
		users:                  users,
		logger:                 logger.WithFields("layer", "service_pull_requests"),
	}
}
