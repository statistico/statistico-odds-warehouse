-- +goose Up
-- +goose StatementBegin
ALTER TABLE market_runner
ADD COLUMN insert_id VARCHAR NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE market_runner
DROP COLUMN insert_id;
-- +goose StatementEnd
