package config

import (
	"log"

	"github.com/mr1cloud/Avito-Reviewer/internal/controller/rest/chi"
	"github.com/mr1cloud/Avito-Reviewer/internal/logger"
	"github.com/mr1cloud/Avito-Reviewer/internal/store/pg"
	"github.com/mr1cloud/Avito-Reviewer/internal/validation"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config is struct that contain app options
type Config struct {
	Logger struct {
		Level            string `env:"LEVEL" env-default:"info" env-description:"Log level: debug, info, warn, error, fatal, panic" validate:"oneof=debug info warn error fatal panic"`
		RotateFileConfig logger.RotateFileConfig
	} `env-prefix:"LOGGER_"`
	Rest  chi.Config `env-prefix:"REST_"`
	Store pg.Config  `env-prefix:"STORE_"`
}

// Load reads config from env into struct
func Load() Config {
	var cfg Config

	// reading config from env
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("error loading config: %s", err)
	}

	// validate config
	if err := validation.Validate.Struct(&cfg); err != nil {
		log.Fatalf("error config validation: %s", err)
	}
	return cfg
}
