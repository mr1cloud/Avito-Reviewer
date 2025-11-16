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

	// ReassignPullRequestReviewers reassigns reviewers for the pull request with the given ID, replacing oldReviewerId with new reviewers.
	ReassignPullRequestReviewers(ctx context.Context, pullRequestId, oldReviewerId string) (*model.PullRequest, *string, error)

	// GetPullRequestsAssignedToUser retrieves all pull requests assigned to the user with the given ID.
	GetPullRequestsAssignedToUser(ctx context.Context, userId string) ([]model.PullRequestShort, error)

	// GetPullRequestsStats retrieves statistics about pull requests.
	GetPullRequestsStats(ctx context.Context) (map[string]int, error)
}
