-- +goose Up
-- +goose StatementBegin
CREATE TABLE market (
    id VARCHAR NOT NULL PRIMARY KEY ,
    event_id INTEGER NOT NULL,
    event_date INTEGER NOT NULL,
    competition_id INTEGER NOT NULL,
    season_id INTEGER NOT NULL,
    name VARCHAR NOT NULL,
    exchange VARCHAR NOT NULL,
    side VARCHAR NOT NULL,
    timestamp INTEGER NOT NULL
);

CREATE TABLE market_runner (
    market_id VARCHAR NOT NULL,
    runner_id INTEGER NOT NULL,
    name VARCHAR NOT NULL,
    timestamp INTEGER NOT NULL
);

CREATE TABLE market_runner_price (
    market_id VARCHAR NOT NULL,
    runner_id INTEGER NOT NULL,
    price FLOAT NOT NULL,
    size FLOAT NOT NULL,
    timestamp INTEGER NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE market;
DROP TABLE market_runner;
DROP TABLE market_runner_price;
-- +goose StatementEnd
