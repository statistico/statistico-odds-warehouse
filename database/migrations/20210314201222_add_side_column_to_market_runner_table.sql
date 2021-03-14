-- +goose Up
-- +goose StatementBegin
ALTER TABLE market
DROP COLUMN side;

ALTER TABLE market_runner
ADD COLUMN side VARCHAR DEFAULT 'BACK';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE market
ADD COLUMN side VARCHAR DEFAULT 'BACK';

ALTER TABLE market_runner
DROP COLUMN side;
-- +goose StatementEnd
