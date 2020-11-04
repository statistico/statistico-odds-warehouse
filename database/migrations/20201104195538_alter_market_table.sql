-- +goose Up
-- +goose StatementBegin
ALTER TABLE market
ADD COLUMN id VARCHAR;

ALTER TABLE market
RENAME COLUMN exchange_market TO exchange_runners;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE market
DROP COLUMN id;

ALTER TABLE market
RENAME COLUMN exchange_runners TO exchange_market;
-- +goose StatementEnd
