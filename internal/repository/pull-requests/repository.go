package pull_requests

import (
	"context"

	"github.com/mr1cloud/Avito-Reviewer/internal/model"
)

// PullRequestsRepository stores pull requests data
//
//goland:noinspection GoNameStartsWithPackageName
type PullRequestsRepository interface {
	// InsertPullRequest creates a new pull request with the given ID, name, and author ID.
	InsertPullRequest(ctx context.Context, pullRequestId, pullRequestName, authorId string, assignedReviewers model.TeamMembers) error
	// SelectPullRequestById retrieves a pull request by its unique identifier.
	SelectPullRequestById(ctx context.Context, pullRequestId string) (*model.PullRequest, error)
	// UpdatePullRequestStatus updates the status of a pull request identified by its ID.
	UpdatePullRequestStatus(ctx context.Context, pullRequestId, status string) error
	// UpdatePullRequestAssignedReviewers updates the list of assigned reviewers for a pull request.
	UpdatePullRequestAssignedReviewers(ctx context.Context, pullRequestId string, oldAssignedReviewer string, newAssignedReviewers model.TeamMember) error
	// SelectPullRequestsAssignedToUser retrieves all pull requests where the given user is assigned as a reviewer.
	SelectPullRequestsAssignedToUser(ctx context.Context, userId string) ([]model.PullRequestShort, error)
	// GetPullRequestsStats retrieves statistics about pull requests.
	GetPullRequestsStats(ctx context.Context) (map[string]int, error)
}
