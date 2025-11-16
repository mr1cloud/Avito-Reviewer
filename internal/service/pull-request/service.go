package pull_request

import (
	"context"

	"github.com/mr1cloud/Avito-Reviewer/internal/model"
)

type PullRequest interface {
	// CreatePullRequest creates a new pull request with the given ID, name, and author ID.
	CreatePullRequest(ctx context.Context, pullRequestId, pullRequestName, authorId string) (*model.PullRequest, error)

	// GetPullRequest retrieves the pull request with the given ID.
	GetPullRequest(ctx context.Context, pullRequestId string) (*model.PullRequest, error)

	// MergePullRequest merges the pull request with the given ID.
	MergePullRequest(ctx context.Context, pullRequestId string) (*model.PullRequest, error)

	// GetPullRequestsAssignedToUser retrieves all pull requests assigned to the user with the given ID.
	GetPullRequestsAssignedToUser(ctx context.Context, userId string) ([]model.PullRequestShort, error)
}
