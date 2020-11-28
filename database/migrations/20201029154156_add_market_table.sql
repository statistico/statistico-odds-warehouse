-- +goose Up
-- +goose StatementBegin
CREATE TABLE market (
    id VARCHAR NOT NULL PRIMARY,
    event_id INTEGER NOT NULL,
    competition_id INTEGER NOT NULL,
    season_id INTEGER NOT NULL,
    name VARCHAR NOT NULL,
    exchange VARCHAR NOT NULL,
    timestamp INTEGER NOT NULL
);

CREATE TABLE market_runner (
    market_id VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    price FLOAT NOT NULL,
    size FLOAT NOT NULL,
    timestamp INTEGER NOT NULL
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE market;
DROP TABLE market_runner;
-- +goose StatementEnd
