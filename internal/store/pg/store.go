package pg

import (
	"fmt"

	"github.com/mr1cloud/Avito-Reviewer/internal/logger"
	"github.com/mr1cloud/Avito-Reviewer/internal/store"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type pgStore struct {
	db     *sqlx.DB
	logger *logger.Logger
}

func (s *pgStore) Close() error {
	return s.db.Close()
}

func NewStore(logger *logger.Logger, cfg Config) (store.Store, error) {
	// pgs connection string
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DB)

	// opening sql connection
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// try to ping db
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log := logger.WithFields("store", "pg")

	return &pgStore{
		db:     db,
		logger: log,
	}, nil
}
