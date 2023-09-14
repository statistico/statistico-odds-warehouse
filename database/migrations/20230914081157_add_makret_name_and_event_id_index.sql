-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_market_event_id_market_name on market(event_id, name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_market_event_id_market_name;
-- +goose StatementEnd
