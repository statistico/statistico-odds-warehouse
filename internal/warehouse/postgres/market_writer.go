package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse"
)

type marketWriter struct {
	connection *sql.DB
}

func (w *marketWriter) InsertMarket(m *warehouse.Market) error {
	var exists bool

	err := w.connection.QueryRow(`SELECT exists (SELECT id FROM market where id = $1)`, m.ID).Scan(&exists)

	if err != nil || exists {
		return err
	}

	builder := w.queryBuilder()

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

func (w *marketWriter) InsertRunners(runners []*warehouse.Runner) error {
	builder := w.queryBuilder()

	for _, runner := range runners {
		if runner.BackPrice != nil {
			_, err := builder.
				Insert("market_runner").
				Columns(
					"id",
					"market_id",
					"name",
					"label",
					"side",
					"price",
					"size",
					"timestamp",
				).
				Values(
					runner.ID,
					runner.MarketID,
					runner.Name,
					runner.Label,
					"BACK",
					runner.BackPrice.Value,
					runner.BackPrice.Size,
					runner.BackPrice.Timestamp.Unix(),
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
					"id",
					"market_id",
					"name",
					"label",
					"side",
					"price",
					"size",
					"timestamp",
				).
				Values(
					runner.ID,
					runner.MarketID,
					runner.Name,
					runner.Label,
					"LAY",
					runner.LayPrice.Value,
					runner.LayPrice.Size,
					runner.LayPrice.Timestamp.Unix(),
				).
				Exec()

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (w *marketWriter) queryBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(w.connection)
}

func NewMarketWriter(connection *sql.DB) warehouse.MarketWriter {
	return &marketWriter{connection: connection}
}
