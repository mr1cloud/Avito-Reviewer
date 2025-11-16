package pull_request

import "net/http"

// PullRequestAlreadyExistsError indicates that a pull request with the given ID already exists.
type PullRequestAlreadyExistsError struct {
	pullRequestId string
}

func (e *PullRequestAlreadyExistsError) ErrorStatusCode() int { return http.StatusConflict }

func (e *PullRequestAlreadyExistsError) Code() string { return "PR_EXISTS" }

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

func (e *NoCandidateReviewersError) ErrorStatusCode() int { return http.StatusConflict }

func (e *NoCandidateReviewersError) Code() string { return "NO_CANDIDATE" }

func (e *NoCandidateReviewersError) Error() string {
	return "no active replacement candidate in team"
}

func NewNoCandidateReviewersError() *NoCandidateReviewersError {
	return &NoCandidateReviewersError{}
}

// PullRequestsAlreadyMergedError indicates that cannot reassign on merged pull request.
type PullRequestsAlreadyMergedError struct{}

func (e *PullRequestsAlreadyMergedError) ErrorStatusCode() int { return http.StatusConflict }

func (e *PullRequestsAlreadyMergedError) Code() string { return "PR_MERGED" }

func (e *PullRequestsAlreadyMergedError) Error() string {
	return "cannot reassign on merged PR"
}

func NewPullRequestsAlreadyMergedError() *PullRequestsAlreadyMergedError {
	return &PullRequestsAlreadyMergedError{}
}

// UserNotAssignedForReviewError indicates that the user is not assigned to this pull request.
type UserNotAssignedForReviewError struct{}

func (e *UserNotAssignedForReviewError) ErrorStatusCode() int { return http.StatusConflict }

func (e *UserNotAssignedForReviewError) Code() string { return "NOT_ASSIGNED" }

func (e *UserNotAssignedForReviewError) Error() string {
	return "reviewer is not assigned to this PR"
}

func NewUserNotAssignedForReviewError() *UserNotAssignedForReviewError {
	return &UserNotAssignedForReviewError{}
}
