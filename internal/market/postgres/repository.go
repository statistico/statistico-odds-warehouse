package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/statistico/statistico-odds-warehouse/internal/market"
	"time"
)

type MarketRepository struct {
	connection *sql.DB
}

func (r *MarketRepository) InsertMarket(m *market.Market) error {
	var exists bool

	err := r.connection.QueryRow(`SELECT exists (SELECT id FROM market where id = $1)`, m.ID).Scan(&exists)

	if err != nil || exists {
		return err
	}

	builder := r.queryBuilder()

	_, err = builder.
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
		).
		Exec()

	return err
}

func (r *MarketRepository) InsertRunners(runners []*market.Runner) error {
	builder := r.queryBuilder()

	for _, runner := range runners {
		_, err := builder.
			Insert("market_runner").
			Columns(
				"market_id",
				"runner_id",
				"name",
				"price",
				"size",
				"timestamp",
			).
			Values(
				runner.MarketID,
				runner.ID,
				runner.Name,
				runner.Price.Value,
				runner.Price.Size,
				runner.Price.Timestamp.Unix(),
			).
			Exec()

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *MarketRepository) MarketRunners(q *market.RunnerQuery) ([]*market.MarketRunner, error) {
	builder := r.queryBuilder()

	rows, err := buildMarketRunnerQuery(q, &builder).Query()

	if err != nil {
		return []*market.MarketRunner{}, err
	}

	var markets []*market.MarketRunner

	for rows.Next() {
		var mkt market.Market
		var run market.Runner
		var date int64
		var timestamp int64

		err := rows.Scan(
			&mkt.ID,
			&mkt.EventID,
			&date,
			&mkt.CompetitionID,
			&mkt.SeasonID,
			&mkt.Name,
			&mkt.Exchange,
			&mkt.Side,
			&run.MarketID,
			&run.ID,
			&run.Name,
			&run.Price.Value,
			&run.Price.Size,
			&timestamp,
		)

		if err != nil {
			return markets, err
		}

		mkt.EventDate = time.Unix(date, 0)
		run.Price.Timestamp = time.Unix(timestamp, 0)

		mr := market.MarketRunner{
			Market: mkt,
			Runner: run,
		}

		markets = append(markets, &mr)
	}

	return markets, nil
}

func (r *MarketRepository) queryBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(r.connection)
}

func NewMarketRepository(connection *sql.DB) *MarketRepository {
	return &MarketRepository{connection: connection}
}
