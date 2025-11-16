package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/mr1cloud/Avito-Reviewer/internal/config"
	"github.com/mr1cloud/Avito-Reviewer/internal/controller/rest"
	"github.com/mr1cloud/Avito-Reviewer/internal/controller/rest/chi"
	"github.com/mr1cloud/Avito-Reviewer/internal/logger"
	"github.com/mr1cloud/Avito-Reviewer/internal/service/pull-request"
	"github.com/mr1cloud/Avito-Reviewer/internal/service/team"
	"github.com/mr1cloud/Avito-Reviewer/internal/service/user"
	"github.com/mr1cloud/Avito-Reviewer/internal/store"
	"github.com/mr1cloud/Avito-Reviewer/internal/store/pg"
)

// App is interface to control and shutdown main app
type App interface {
	// Run starts main app processes and graceful shutdown on <CTRL+C> or context ends.
	Run(ctx context.Context) error
}

type app struct {
	logger *logger.Logger

	restServer rest.Server
	store      store.Store

	users        user.User
	teams        team.Team
	pullRequests pull_request.PullRequest
}

// NewApp creates new instance of App
func NewApp(logger *logger.Logger, cfg config.Config) App {
	var a app

	a.logger = logger.WithFields("layer", "app")

	// creating store
	store, err := pg.NewStore(logger, cfg.Store)
	if err != nil {
		a.logger.Fatalf("error creating store: %s", err)
	}
	a.store = store

	// creating services
	a.users = user.NewService(logger.WithFields("layer", "service-users"), a.store.UsersRepository())
	a.teams = team.NewService(logger.WithFields("layer", "service-teams"), a.store.TeamsRepository())
	a.pullRequests = pull_request.NewService(logger.WithFields("layer", "service-pull-requests"), a.store.PullRequestsRepository(), a.users, a.teams)

	// creating rest server
	a.restServer = chi.NewServer(logger.WithFields("layer", "rest"), cfg.Rest, a.users, a.teams, a.pullRequests)

	return &a
}

func (a *app) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// receive sys signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup

	// Run rest server
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := a.restServer.Run(ctx); err != nil {
			a.logger.Errorf("error running rest server: %s", err)
		}
		cancel()
	}()

	a.logger.Info("app is running")

	// wait until signal will come or context will end
	select {
	case <-quit:
		a.logger.Info("shutdown signal received")
	case <-ctx.Done():
		a.logger.Info("context canceled")
	}

	a.logger.Info("stopping all service")
	// context cancel => bot stopping
	cancel()
	// wait until bot will be stopped
	wg.Wait()

	a.logger.Info("service was gracefully shutdown")
	return nil
}
