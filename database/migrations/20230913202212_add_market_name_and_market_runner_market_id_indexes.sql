-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_market_name on market(name);
CREATE INDEX idx_market_runner_market_id on market_runner(market_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_market_name;
DROP INDEX idx_market_runner_market_id;
-- +goose StatementEnd
