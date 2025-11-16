-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION set_merged_at()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status = 'MERGED' AND OLD.status <> 'MERGED' THEN
        NEW.merged_at = NOW();
    ELSIF NEW.status <> 'MERGED' THEN
        NEW.merged_at = NULL;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_set_merged_at
BEFORE UPDATE ON pull_requests
FOR EACH ROW
EXECUTE FUNCTION set_merged_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS trg_set_merged_at ON pull_requests;
DROP FUNCTION IF EXISTS set_merged_at();
-- +goose StatementEnd
