package pull_requests

import "github.com/mr1cloud/Avito-Reviewer/internal/model"

type CreatePrRequest struct {
	PullRequestID   string `json:"pull_request_id" validate:"required"`
	PullRequestName string `json:"pull_request_name" validate:"required"`
	AuthorID        string `json:"author_id" validate:"required"`
}

type MergePrRequest struct {
	PullRequestID string `json:"pull_request_id" validate:"required"`
}

type ReassignReviewerPrRequest struct {
	PullRequestID string `json:"pull_request_id" validate:"required"`
	OldReviewerID string `json:"old_reviewer_id" validate:"required"`
}

type PullRequestResponse struct {
	PullRequest model.PullRequest `json:"pr"`
	ReplacedBy  *string           `json:"replaced_by,omitempty"`
}
