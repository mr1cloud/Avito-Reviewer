package chi

import (
	"fmt"
	"net/http"

	"github.com/mr1cloud/Avito-Reviewer/internal/controller/rest/middleware"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func (s *server) initRoutes() {
	r := chi.NewRouter()

	r.Use(middleware.LoggerMiddleware(s.logger))
	r.Use(chimiddleware.Recoverer)

	r.Route(s.cfg.BasePath, func(r chi.Router) {

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
