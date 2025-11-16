package pull_request

import "net/http"

// PullRequestAlreadyExistsError indicates that a pull request with the given ID already exists.
type PullRequestAlreadyExistsError struct {
	pullRequestId string
}

// ErrorStatusCode returns the HTTP status code for the error.
func (e *PullRequestAlreadyExistsError) ErrorStatusCode() int { return http.StatusConflict }

// Code returns the error code.
func (e *PullRequestAlreadyExistsError) Code() string { return "PR_EXISTS" }

// Error returns the error message.
func (e *PullRequestAlreadyExistsError) Error() string {
	return "PR " + e.pullRequestId + " already exists"
}

func NewPullRequestAlreadyExistsError(pullRequestId string) *PullRequestAlreadyExistsError {
	return &PullRequestAlreadyExistsError{
		pullRequestId: pullRequestId,
	}
}

// NoCandidateReviewersError indicates that there are no active replacement candidate in team.
type NoCandidateReviewersError struct{}

// ErrorStatusCode returns the HTTP status code for the error.
func (e *NoCandidateReviewersError) ErrorStatusCode() int { return http.StatusConflict }

// Code returns the error code.
func (e *NoCandidateReviewersError) Code() string { return "NO_CANDIDATE" }

// Error returns the error message.
func (e *NoCandidateReviewersError) Error() string {
	return "no active replacement candidate in team"
}

func NewNoCandidateReviewersError() *NoCandidateReviewersError {
	return &NoCandidateReviewersError{}
}

// PullRequestsAlreadyMergedError indicates that cannot reassign on merged pull request.
type PullRequestsAlreadyMergedError struct{}

// ErrorStatusCode returns the HTTP status code for the error.
func (e *PullRequestsAlreadyMergedError) ErrorStatusCode() int { return http.StatusConflict }

// Code returns the error code.
func (e *PullRequestsAlreadyMergedError) Code() string { return "PR_MERGED" }

// Error returns the error message.
func (e *PullRequestsAlreadyMergedError) Error() string {
	return "cannot reassign on merged PR"
}

func NewPullRequestsAlreadyMergedError() *PullRequestsAlreadyMergedError {
	return &PullRequestsAlreadyMergedError{}
}

// UserNotAssignedForReviewError indicates that the user is not assigned to this pull request.
type UserNotAssignedForReviewError struct{}

// ErrorStatusCode returns the HTTP status code for the error.
func (e *UserNotAssignedForReviewError) ErrorStatusCode() int { return http.StatusConflict }

// Code returns the error code.
func (e *UserNotAssignedForReviewError) Code() string { return "NOT_ASSIGNED" }

// Error returns the error message.
func (e *UserNotAssignedForReviewError) Error() string {
	return "reviewer is not assigned to this PR"
}

func NewUserNotAssignedForReviewError() *UserNotAssignedForReviewError {
	return &UserNotAssignedForReviewError{}
}
