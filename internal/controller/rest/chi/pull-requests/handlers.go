package pull_requests

import (
	"errors"
	"net/http"

	"github.com/mr1cloud/Avito-Reviewer/internal/controller/rest/tools"
	"github.com/mr1cloud/Avito-Reviewer/internal/service/pull-request"

	serviceerrors "github.com/mr1cloud/Avito-Reviewer/internal/error"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	Router       *chi.Mux
	PullRequests pull_request.PullRequest
}

func NewPullRequestsHandler(pullRequests pull_request.PullRequest) *Handlers {
	handlers := &Handlers{
		Router:       chi.NewRouter(),
		PullRequests: pullRequests,
	}

	return handlers
}

// PostCreatePullRequest godoc
//
//	@Summary		Create a new pull request
//	@Description	Creates a new pull request with the given ID, name, and author ID.
//	@Tags			PullRequests
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreatePrRequest	true	"Pull request creation data"
//	@Success		201		{object}	PullRequestResponse
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		500		{object}	model.ErrorResponse
//	@Router			/pullRequest/create [post]
func (h *Handlers) PostCreatePullRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreatePrRequest
		if err := tools.DecodeJSON(r, &req); err != nil {
			tools.RespondWithError(w, err.(serviceerrors.ServiceError).ErrorStatusCode(), "USER_FATAL", err.Error())
			return
		}

		pullRequest, err := h.PullRequests.CreatePullRequest(r.Context(), req.PullRequestID, req.PullRequestName, req.AuthorID)
		if err != nil {
			var srvErr serviceerrors.ServiceError
			if errors.As(err, &srvErr) {
				tools.RespondWithError(w, srvErr.ErrorStatusCode(), srvErr.Code(), srvErr.Error())
			} else {
				tools.RespondWithError(w, http.StatusInternalServerError, "FATAL", "Internal server error")
			}
			return
		}

		tools.RespondJSON(w, http.StatusCreated, PullRequestResponse{
			PullRequest: *pullRequest,
		})
	}
}

// PostMergePullRequest godoc
//
//	@Summary		Merge a pull request
//	@Description	Merges the pull request with the given ID.
//	@Tags			PullRequests
//	@Accept			json
//	@Produce		json
//	@Param			request	body		MergePrRequest	true	"Pull request merge data"
//	@Success		200		{object}	PullRequestResponse
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		500		{object}	model.ErrorResponse
//	@Router			/pullRequest/merge [post]
func (h *Handlers) PostMergePullRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req MergePrRequest
		if err := tools.DecodeJSON(r, &req); err != nil {
			tools.RespondWithError(w, err.(serviceerrors.ServiceError).ErrorStatusCode(), "USER_FATAL", err.Error())
			return
		}

		pullRequest, err := h.PullRequests.MergePullRequest(r.Context(), req.PullRequestID)
		if err != nil {
			var srvErr serviceerrors.ServiceError
			if errors.As(err, &srvErr) {
				tools.RespondWithError(w, srvErr.ErrorStatusCode(), srvErr.Code(), srvErr.Error())
			} else {
				tools.RespondWithError(w, http.StatusInternalServerError, "FATAL", "Internal server error")
			}
			return
		}

		tools.RespondJSON(w, http.StatusOK, PullRequestResponse{
			PullRequest: *pullRequest,
		})
	}
}

// PostReassignReviewerPullRequest godoc
//
//	@Summary		Reassign a reviewer for a pull request
//	@Description	Reassigns the specified reviewer for the given pull request to a new reviewer.
//	@Tags			PullRequests
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ReassignReviewerPrRequest	true	"Pull request reviewer reassignment data"
//	@Success		200		{object}	PullRequestResponse
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		500		{object}	model.ErrorResponse
//	@Router			/pullRequest/reassign [post]
func (h *Handlers) PostReassignReviewerPullRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ReassignReviewerPrRequest
		if err := tools.DecodeJSON(r, &req); err != nil {
			tools.RespondWithError(w, err.(serviceerrors.ServiceError).ErrorStatusCode(), "USER_FATAL", err.Error())
			return
		}

		pullRequest, replacedBy, err := h.PullRequests.ReassignPullRequestReviewers(r.Context(), req.PullRequestID, req.OldReviewerID)
		if err != nil {
			var srvErr serviceerrors.ServiceError
			if errors.As(err, &srvErr) {
				tools.RespondWithError(w, srvErr.ErrorStatusCode(), srvErr.Code(), srvErr.Error())
			} else {
				tools.RespondWithError(w, http.StatusInternalServerError, "FATAL", "Internal server error")
			}
			return
		}

		tools.RespondJSON(w, http.StatusOK, PullRequestResponse{
			PullRequest: *pullRequest,
			ReplacedBy:  replacedBy,
		})
	}
}

// GetPullRequestsStats godoc
//
//	@Summary		Get pull requests statistics
//	@Description	Retrieves statistics about pull requests.
//	@Tags			PullRequests
//	@Produce		json
//	@Success		200	{object}	PullRequestsStatsResponse
//	@Failure		500	{object}	model.ErrorResponse
//	@Router			/pullRequest/stats [get]
func (h *Handlers) GetPullRequestsStats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats, err := h.PullRequests.GetPullRequestsStats(r.Context())
		if err != nil {
			var srvErr serviceerrors.ServiceError
			if errors.As(err, &srvErr) {
				tools.RespondWithError(w, srvErr.ErrorStatusCode(), srvErr.Code(), srvErr.Error())
			} else {
				tools.RespondWithError(w, http.StatusInternalServerError, "FATAL", "Internal server error")
			}
			return
		}

		tools.RespondJSON(w, http.StatusOK, stats)
	}
}
