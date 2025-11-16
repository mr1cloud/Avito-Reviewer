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

type PullRequestsRepository struct {
	DB     *sqlx.DB
	Logger *logger.Logger
}

func (p *PullRequestsRepository) InsertPullRequest(ctx context.Context, pullRequestId, pullRequestName, authorId string, assignedReviewers model.TeamMembers) error {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		p.Logger.Errorf("error starting create pull request transaction: %v", err)
		return err
	}
	defer func() {
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				p.Logger.Errorf("error rolling back create pull request transaction: %v", rollErr)
			}
			return
		}
	}()

	query := `INSERT INTO pull_requests (pull_request_id, pull_request_name, author_id, status) VALUES ($1, $2, $3, 'OPEN');`
	_, err = tx.ExecContext(ctx, query, pullRequestId, pullRequestName, authorId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return repository.ErrConflict
			}
		}
		p.Logger.Errorf("error inserting pull request: %v", err)
		return err
	}

	query = `INSERT INTO assigned_reviewers (pull_request_id, reviewer_id) VALUES ($1, $2);`
	for _, assignedReviewer := range assignedReviewers {
		_, err = tx.ExecContext(ctx, query, pullRequestId, assignedReviewer.UserID)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == "23505" {
					return repository.ErrConflict
				}
			}
			p.Logger.Errorf("error inserting assigned reviewer: %v", err)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		p.Logger.Errorf("error committing create pull request transaction: %v", err)
		return err
	}

	return nil
}

func (p *PullRequestsRepository) SelectPullRequestById(ctx context.Context, pullRequestId string) (*model.PullRequest, error) {
	var pr model.PullRequest

	query := `
	SELECT pr.*,
		   COALESCE(json_agg(ar.reviewer_id::text) FILTER (WHERE ar.reviewer_id IS NOT NULL),
					'[]'::json) AS assigned_reviewers
	FROM pull_requests pr
			 LEFT JOIN assigned_reviewers ar ON pr.pull_request_id = ar.pull_request_id
	WHERE pr.pull_request_id = $1
	GROUP BY pr.pull_request_id;
	`
	err := p.DB.GetContext(ctx, &pr, query, pullRequestId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		p.Logger.Errorf("error getting pull request by id: %v", err)
		return nil, err
	}

	return &pr, nil
}

func (p *PullRequestsRepository) UpdatePullRequestStatus(ctx context.Context, pullRequestId, status string) error {
	query := `UPDATE pull_requests SET status = $1 WHERE pull_request_id = $2;`
	_, err := p.DB.ExecContext(ctx, query, status, pullRequestId)
	if err != nil {
		p.Logger.Errorf("error updating pull request status: %v", err)
		return err
	}

	return nil
}

func (p *PullRequestsRepository) UpdatePullRequestAssignedReviewers(ctx context.Context, pullRequestId string, oldAssignedReviewer string, newAssignedReviewers model.TeamMember) error {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		p.Logger.Errorf("error starting update pull request assigned reviewers transaction: %v", err)
		return err
	}
	defer func() {
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				p.Logger.Errorf("error rolling back update pull request assigned reviewers transaction: %v", rollErr)
			}
			return
		}
	}()

	query := `DELETE FROM assigned_reviewers WHERE pull_request_id = $1 AND reviewer_id = $2;`
	_, err = tx.ExecContext(ctx, query, pullRequestId, oldAssignedReviewer)
	if err != nil {
		p.Logger.Errorf("error deleting old assigned reviewer: %v", err)
		return err
	}

	query = `INSERT INTO assigned_reviewers (pull_request_id, reviewer_id) VALUES ($1, $2);`
	_, err = tx.ExecContext(ctx, query, pullRequestId, newAssignedReviewers.UserID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return repository.ErrConflict
			}
		}
		p.Logger.Errorf("error inserting new assigned reviewer: %v", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		p.Logger.Errorf("error committing update pull request assigned reviewers transaction: %v", err)
		return err
	}

	return nil
}

func (p *PullRequestsRepository) SelectPullRequestsAssignedToUser(ctx context.Context, userId string) ([]model.PullRequestShort, error) {
	var prs []model.PullRequestShort

	query := `
	SELECT pr.pull_request_id,
		   pr.pull_request_name,
		   pr.author_id,
		   pr.status
	FROM pull_requests pr
			 JOIN assigned_reviewers ar ON pr.pull_request_id = ar.pull_request_id
	WHERE ar.reviewer_id = $1;
	`
	err := p.DB.SelectContext(ctx, &prs, query, userId)
	if err != nil {
		p.Logger.Errorf("error getting pull requests assigned to user: %v", err)
		return nil, err
	}

	if prs == nil {
		prs = make([]model.PullRequestShort, 0)
	}

	return prs, nil
}
