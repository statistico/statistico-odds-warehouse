-- +goose Up
-- +goose StatementBegin
CREATE TABLE market_over_under (
    id VARCHAR NOT NULL,
    event_id INTEGER NOT NULL,
    name VARCHAR NOT NULL,
    exchange VARCHAR NOT NULL,
    side VARCHAR NOT NULL,
    over_price FLOAT NOT NULL,
    over_size FLOAT NOT NULL,
    under_price FLOAT NOT NULL,
    under_size FLOAT NOT NULL,
    timestamp INTEGER NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE market_over_under;
-- +goose StatementEnd
