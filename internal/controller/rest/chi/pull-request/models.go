package pull_request

import "github.com/mr1cloud/Avito-Reviewer/internal/model"

type CreatePrRequest struct {
	PullRequestID   string `json:"pull_request_id" validate:"required"`
	PullRequestName string `json:"pull_request_name" validate:"required"`
	AuthorID        string `json:"author_id" validate:"required"`
}

type PullRequestResponse struct {
	PullRequest model.PullRequest `json:"pr"`
}
