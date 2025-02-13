-- +goose Up
-- +goose StatementBegin
CREATE TABLE market (
    id VARCHAR NOT NULL PRIMARY KEY,
    event_id INTEGER NOT NULL,
    event_date INTEGER NOT NULL,
    competition_id INTEGER NOT NULL,
    season_id INTEGER NOT NULL,
    name VARCHAR NOT NULL,
    exchange VARCHAR NOT NULL
);

CREATE TABLE market_runner (
    id INTEGER NOT NULL,
    market_id VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    label VARCHAR,
    side VARCHAR NOT NULL,
    price FLOAT NOT NULL,
    size FLOAT NOT NULL,
    timestamp INTEGER NOT NULL
);

CREATE INDEX idx_market_name on market(name);
CREATE INDEX idx_market_runner_market_id on market_runner(market_id);
CREATE INDEX idx_market_event_id_market_name on market(event_id, name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE market;
DROP TABLE market_runner;

DROP INDEX idx_market_name;
DROP INDEX idx_market_runner_market_id;
DROP INDEX idx_market_event_id_market_name;
-- +goose StatementEnd
