package main

import (
	"context"

	"github.com/mr1cloud/Avito-Reviewer/internal/app"
	"github.com/mr1cloud/Avito-Reviewer/internal/config"
	"github.com/mr1cloud/Avito-Reviewer/internal/logger"
)

func main() {
	cfg := config.Load()
	log := logger.NewLogger("reviewer", cfg.Logger.Level, &cfg.Logger.RotateFileConfig)
	ctx := context.Background()

	// creating app
	application := app.NewApp(log, cfg)

	// running rest server
	err := application.Run(ctx)
	if err != nil {
		log.Fatalf("error running application: %s", err)
	}
}
