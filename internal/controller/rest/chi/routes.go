package chi

import (
	"fmt"
	"net/http"

	"github.com/mr1cloud/Avito-Reviewer/internal/controller/rest/middleware"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *server) initRoutes() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: s.cfg.AllowOrigins,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders: []string{"Link"},
	}))
	r.Use(middleware.LoggerMiddleware(s.logger))
	r.Use(chimiddleware.Recoverer)

	r.Route(s.cfg.BasePath, func(r chi.Router) {
		// Teams routes
		r.Route("/team", func(r chi.Router) {
			r.Post("/add", s.teamsHandlers.PostAddTeam())
			r.Put("/update", s.teamsHandlers.PutUpdateTeam())
			r.Get("/get", s.teamsHandlers.GetTeam())
		})

		// Users routes
		r.Route("/users", func(r chi.Router) {
			r.Post("/setIsActive", s.usersHandlers.PostSetUserIsActive())
			r.Get("/getReview", s.usersHandlers.GetPullRequestsAssignedToUser())
		})

		// Pull requests routes
		r.Route("/pullRequest", func(r chi.Router) {
			r.Post("/create", s.pullRequestsHandlers.PostCreatePullRequest())
		})
	})

	if s.cfg.DocsEnabled {
		r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
			htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
				SpecContent: s.openApiSpec,
				DarkMode:    true,
			})
			if err != nil {
				s.logger.Errorf("failed to load swagger docs: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "text/html")
			_, err = w.Write([]byte(fmt.Sprintf("%s", htmlContent)))
			if err != nil {
				s.logger.Errorf("failed to write swagger docs response: %v", err)
			}
		})
	}

	s.app.Handler = r
}
