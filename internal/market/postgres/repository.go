package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/statistico/statistico-odds-warehouse/internal/market"
)

type MarketRepository struct {
	connection *sql.DB
}

func (r *MarketRepository) InsertMarket(m *market.Market) error {
	var exists bool

	err := r.connection.QueryRow(`SELECT exists (SELECT id FROM market where id = $1)`, m.ID).Scan(&exists)

	if err != nil {
		return err
	}

	if !exists {
		if err := r.insertMarket(m); err != nil {
			return err
		}
	}

	insertId := uuid.New()

	for _, run := range m.Runners {
		err := r.insertRunner(run, m.ID, m.Timestamp, insertId)

		if err != nil {
			return err
		}
	}

	return err
}

func (r *MarketRepository) insertMarket(m *market.Market) error {
	builder := r.queryBuilder()

	_, err := builder.
		Insert("market").
		Columns(
			"id",
			"event_id",
			"event_date",
			"competition_id",
			"season_id",
			"name",
			"exchange",
			"side",
			"timestamp",
		).
		Values(
			m.ID,
			m.EventID,
			m.EventDate.Unix(),
			m.CompetitionID,
			m.SeasonID,
			m.Name,
			m.Exchange,
			m.Side,
			m.Timestamp,
		).
		Exec()

	return err
}

func (r *MarketRepository) insertRunner(runner *market.Runner, marketID string, timestamp int64, insertId uuid.UUID) error {
	builder := r.queryBuilder()

	_, err := builder.
		Insert("market_runner").
		Columns(
			"market_id",
			"id",
			"name",
			"price",
			"size",
			"timestamp",
			"insert_id",
		).
		Values(
			marketID,
			runner.ID,
			runner.Name,
			runner.Price,
			runner.Size,
			timestamp,
			insertId.String(),
		).
		Exec()

	return err
}

func (r *MarketRepository) queryBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(r.connection)
}

func NewMarketRepository(connection *sql.DB) *MarketRepository {
	return &MarketRepository{connection: connection}
}
