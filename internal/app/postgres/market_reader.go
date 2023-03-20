package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/statistico/statistico-odds-warehouse/internal/app"
	"time"
)

type marketReader struct {
	connection *sql.DB
}

func (m *marketReader) ExchangeMarketRunnerOdds(eventID uint64, market, runner, exchange string, limit uint32) ([]*app.Odds, error) {
	b := m.queryBuilder()

	rows, err := b.
		Select(
			"mr.price",
			"mr.timestamp",
		).
		From("market m").
		Join("join market_runner mr on mr.market_id = m.id").
		Where(sq.Eq{"event_id": eventID}).
		Where(sq.Eq{"exchange": exchange}).
		Where(sq.Eq{"name": market}).
		Where(sq.Eq{"mr.name": runner}).
		OrderBy("mr.timestamp DESC").
		Limit(uint64(limit)).
		Query()

	if err != nil {
		return []*app.Odds{}, err
	}

	defer rows.Close()

	var odds []*app.Odds

	for rows.Next() {
		var o *app.Odds
		var timestamp int64

		if err := rows.Scan(&o.Price, &timestamp); err != nil {
			return []*app.Odds{}, err
		}

		o.Timestamp = time.Unix(timestamp, 0)

		odds = append(odds, o)
	}

	return odds, nil
}

func (m *marketReader) queryBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(m.connection)
}

func NewMarketReader(connection *sql.DB) app.MarketReader {
	return &marketReader{connection: connection}
}
