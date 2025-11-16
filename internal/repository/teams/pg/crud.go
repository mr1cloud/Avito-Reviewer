package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mr1cloud/Avito-Reviewer/internal/logger"
	"github.com/mr1cloud/Avito-Reviewer/internal/model"
	"github.com/mr1cloud/Avito-Reviewer/internal/repository"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
)

type TeamsRepository struct {
	DB     *sqlx.DB
	Logger *logger.Logger
}

func (t *TeamsRepository) InsertTeam(ctx context.Context, teamName string, members model.TeamMembers) error {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		t.Logger.Errorf("error starting create team transaction: %v", err)
		return err
	}
	defer func() {
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				t.Logger.Errorf("error rolling back create team transaction: %v", rollErr)
			}
			return
		}
	}()

	query := `INSERT INTO teams (team_name) VALUES ($1);`
	_, err = tx.ExecContext(ctx, query, teamName)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return repository.ErrConflict
			}
		}
		t.Logger.Errorf("error inserting team: %v", err)
		return err
	}

	query = `INSERT INTO users (user_id, username, is_active, team_name) VALUES ($1, $2, $3, $4);`
	for _, member := range members {
		_, err = tx.ExecContext(ctx, query, member.UserID, member.Username, member.IsActive, teamName)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == "23505" {
					return repository.ErrConflict
				}
			}
			t.Logger.Errorf("error inserting team member: %v", err)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		t.Logger.Errorf("error committing create team transaction: %v", err)
		return err
	}

	return nil
}

func (t *TeamsRepository) UpdateTeam(ctx context.Context, teamName string, oldMembers model.TeamMembers, newMembers model.TeamMembers) error {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		t.Logger.Errorf("error starting update team transaction: %v", err)
		return err
	}
	defer func() {
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				t.Logger.Errorf("error rolling back update team transaction: %v", rollErr)
			}
			return
		}
	}()

	query := `DELETE FROM users WHERE user_id = $1 AND team_name = $2;`
	for _, oldMember := range oldMembers {
		if !newMembers.Contains(oldMember.UserID) {
			_, err = tx.ExecContext(ctx, query, oldMember.UserID, teamName)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					continue
				}
				t.Logger.Errorf("error deleting old team member: %v", err)
				return err
			}
		}
	}

	query = `INSERT INTO users (user_id, username, is_active, team_name) VALUES ($1, $2, $3, $4);`
	for _, newMember := range newMembers {
		if !oldMembers.Contains(newMember.UserID) {
			_, err = tx.ExecContext(ctx, query, newMember.UserID, newMember.Username, newMember.IsActive, teamName)
			if err != nil {
				var pgErr *pgconn.PgError
				if errors.As(err, &pgErr) {
					if pgErr.Code == "23505" {
						return repository.ErrConflict
					}
				}
				t.Logger.Errorf("error inserting team member: %v", err)
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		t.Logger.Errorf("error committing update team transaction: %v", err)
		return err
	}

	return nil
}

func (t *TeamsRepository) SelectTeam(ctx context.Context, teamName string) (*model.Team, error) {
	var team model.Team

	query := `
	SELECT *,
       COALESCE(
               (SELECT json_agg(m)
                FROM (SELECT *
                      FROM users
                      WHERE team_name = $1) AS m),
               '[]'::json
       ) AS members
	FROM teams
	WHERE team_name = $1;
	`
	err := t.DB.GetContext(ctx, &team, query, teamName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		t.Logger.Errorf("error getting team by name: %v", err)
		return nil, err
	}

	return &team, nil
}
