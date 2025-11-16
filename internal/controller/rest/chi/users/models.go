package users

import "github.com/mr1cloud/Avito-Reviewer/internal/model"

type SetUserIsActiveRequest struct {
	UserID   string `json:"user_id" validate:"required,max=32"`
	IsActive bool   `json:"is_active" validate:"required"`
}

type UserResponse struct {
	User model.User `json:"user"`
}

type PullRequestsAssignedToUserResponse struct {
	UserID       string                   `json:"user_id"`
	PullRequests []model.PullRequestShort `json:"pull_requests"`
}
