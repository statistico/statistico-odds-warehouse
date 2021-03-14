package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/statistico/statistico-odds-warehouse/internal/app/market"
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
		).
		Values(
			m.ID,
			m.EventID,
			m.EventDate.Unix(),
			m.CompetitionID,
			m.SeasonID,
			m.Name,
			m.Exchange,
		).
		Exec()

	return err
}

func (r *MarketRepository) InsertRunners(runners []*market.Runner) error {
	builder := r.queryBuilder()

	for _, runner := range runners {
		if runner.BackPrice != nil {
			_, err := builder.
				Insert("market_runner").
				Columns(
					"market_id",
					"runner_id",
					"name",
					"price",
					"size",
					"timestamp",
					"side",
				).
				Values(
					runner.MarketID,
					runner.ID,
					runner.Name,
					runner.BackPrice.Value,
					runner.BackPrice.Size,
					runner.BackPrice.Timestamp.Unix(),
					"BACK",
				).
				Exec()

			if err != nil {
				return err
			}
		}

		if runner.LayPrice != nil {
			_, err := builder.
				Insert("market_runner").
				Columns(
					"market_id",
					"runner_id",
					"name",
					"price",
					"size",
					"timestamp",
					"side",
				).
				Values(
					runner.MarketID,
					runner.ID,
					runner.Name,
					runner.LayPrice.Value,
					runner.LayPrice.Size,
					runner.LayPrice.Timestamp.Unix(),
					"LAY",
				).
				Exec()

			if err != nil {
				return err
			}
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
		var mr market.MarketRunner
		var value float32
		var size float32
		var side string
		var date int64
		var timestamp int64

		err := rows.Scan(
			&mr.MarketID,
			&mr.EventID,
			&date,
			&mr.CompetitionID,
			&mr.SeasonID,
			&mr.MarketName,
			&mr.Exchange,
			&mr.RunnerID,
			&mr.RunnerName,
			&value,
			&size,
			&side,
			&timestamp,
		)

		if err != nil {
			return markets, err
		}

		mr.EventDate = time.Unix(date, 0)

		mr.Price = market.Price{
			Value:     value,
			Size:      size,
			Side:      side,
			Timestamp: time.Unix(timestamp, 0),
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
