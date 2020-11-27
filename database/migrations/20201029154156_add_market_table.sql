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

CREATE TABLE market_btts (
    id VARCHAR NOT NULL,
    event_id INTEGER NOT NULL,
    name VARCHAR NOT NULL,
    exchange VARCHAR NOT NULL,
    side VARCHAR NOT NULL,
    yes_price FLOAT NOT NULL,
    yes_size FLOAT NOT NULL,
    no_price FLOAT NOT NULL,
    no_size FLOAT NOT NULL,
    timestamp INTEGER NOT NULL
);

CREATE TABLE market_match_odds (
    id VARCHAR NOT NULL,
    event_id INTEGER NOT NULL,
    name VARCHAR NOT NULL,
    exchange VARCHAR NOT NULL,
    side VARCHAR NOT NULL,
    home_price FLOAT NOT NULL,
    home_size FLOAT NOT NULL,
    away_price FLOAT NOT NULL,
    away_size FLOAT NOT NULL,
    draw_price FLOAT NOT NULL,
    draw_size FLOAT NOT NULL,
    timestamp INTEGER NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE market_over_under;
DROP TABLE market_btts;
DROP TABLE market_match_odds;
-- +goose StatementEnd
