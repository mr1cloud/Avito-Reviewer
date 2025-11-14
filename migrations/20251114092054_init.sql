-- +goose Up
-- +goose StatementBegin
CREATE TYPE pr_status AS ENUM ('OPEN', 'MERGED');

CREATE TABLE teams (
    team_name VARCHAR(64) PRIMARY KEY UNIQUE
);

CREATE TABLE users (
    user_id VARCHAR(32) PRIMARY KEY UNIQUE,
    username VARCHAR(64) NOT NULL UNIQUE,
    is_active BOOLEAN NOT NULL,
    team_name VARCHAR(64) NOT NULL,

    FOREIGN KEY (team_name) REFERENCES teams (team_name) ON DELETE CASCADE
);

CREATE TABLE pull_requests (
    pull_request_id VARCHAR(64) PRIMARY KEY UNIQUE,
    pull_request_name VARCHAR(128) NOT NULL,
    author_id VARCHAR(32) NOT NULL,
    status pr_status NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    merged_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE assigned_reviewers (
    pull_request_id VARCHAR(64) NOT NULL,
    reviewer_id VARCHAR(32) NOT NULL,

    PRIMARY KEY (pull_request_id, reviewer_id),

    FOREIGN KEY (pull_request_id) REFERENCES pull_requests (pull_request_id) ON DELETE CASCADE,
    FOREIGN KEY (reviewer_id) REFERENCES users (user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS assigned_reviewers;
DROP TABLE IF EXISTS pull_requests;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS teams;

DROP TYPE IF EXISTS pr_status;
-- +goose StatementEnd
