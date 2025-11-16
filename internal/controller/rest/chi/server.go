package chi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/mr1cloud/Avito-Reviewer/internal/controller/rest"
	"github.com/mr1cloud/Avito-Reviewer/internal/logger"

	"github.com/mr1cloud/Avito-Reviewer/internal/controller/rest/chi/pull-requests"
	"github.com/mr1cloud/Avito-Reviewer/internal/controller/rest/chi/teams"
	"github.com/mr1cloud/Avito-Reviewer/internal/controller/rest/chi/users"

	"github.com/mr1cloud/Avito-Reviewer/internal/service/pull-request"
	"github.com/mr1cloud/Avito-Reviewer/internal/service/team"
	"github.com/mr1cloud/Avito-Reviewer/internal/service/user"

	_ "github.com/mr1cloud/Avito-Reviewer/docs/swagger"

	"github.com/swaggo/swag"
)

type server struct {
	app         *http.Server
	openApiSpec interface{}
	cfg         Config

	logger *logger.Logger

	users        user.User
	teams        team.Team
	pullRequests pull_request.PullRequest

	usersHandlers        *users.Handlers
	teamsHandlers        *teams.Handlers
	pullRequestsHandlers *pull_requests.Handlers
}

func NewServer(logger *logger.Logger, cfg Config, userService user.User, teamService team.Team, pullRequestService pull_request.PullRequest) rest.Server {
	var s server

	// set config
	s.cfg = cfg

	// set logger
	s.logger = logger.WithFields("layer", "rest")

	// set services
	s.users = userService
	s.teams = teamService
	s.pullRequests = pullRequestService

	// init handlers
	s.usersHandlers = users.NewUsersHandler(s.users, s.pullRequests)
	s.teamsHandlers = teams.NewTeamsHandler(s.teams)
	s.pullRequestsHandlers = pull_requests.NewPullRequestsHandler(s.pullRequests)

	// create http server
	s.app = &http.Server{
		Addr: cfg.Host + ":" + strconv.Itoa(cfg.Port),
	}

	// read openapi spec
	if s.cfg.DocsEnabled {
		if err := s.ReadOpenAPISpec(); err != nil {
			s.logger.Errorf("failed to read openapi spec: %v", err)
		}
	}

	// init routes
	s.initRoutes()

	s.logger.Info("chi rest server was initialized")

	return &s
}

func (s *server) Run(ctx context.Context) error {
	s.logger.Info("starting chi rest server on " + s.cfg.Host + ":" + strconv.Itoa(s.cfg.Port))

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		if err := s.app.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				s.logger.Errorf("error running chi rest server: %s", err)
			}
			cancel()
		}
	}()

	<-ctx.Done()
	s.logger.Info("stopping chi rest server")
	if err := s.app.Shutdown(context.Background()); err != nil {
		s.logger.Errorf("error shutting down chi rest server: %s", err)
		return err
	}
	s.logger.Info("chi rest server stopped")
	return nil
}

func (s *server) ReadOpenAPISpec() error {
	result, err := swag.ReadDoc()
	if err != nil {
		return err
	}

	var spec interface{}
	err = json.Unmarshal([]byte(result), &spec)
	if err != nil {
		return err
	}

	s.openApiSpec = spec
	return nil
}
