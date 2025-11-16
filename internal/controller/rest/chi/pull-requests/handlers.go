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
