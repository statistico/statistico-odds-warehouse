package postgres

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse"
	"time"
)

type marketReader struct {
	connection *sql.DB
}

func (m *marketReader) ExchangeMarketRunnerOdds(eventID uint64, market, runner, exchange string, limit uint32) ([]*warehouse.Odds, error) {
	b := m.queryBuilder()

	rows, err := b.
		Select(
			"mr.price",
			"mr.size",
			"mr.side",
			"mr.timestamp",
		).
		From("market m").
		Join("market_runner mr on mr.market_id = m.id").
		Where(sq.Eq{"m.event_id": eventID}).
		Where(sq.Eq{"m.exchange": exchange}).
		Where(sq.Eq{"m.name": market}).
		Where(sq.Eq{"mr.side": "BACK"}).
		Where(sq.Eq{"mr.name": runner}).
		OrderBy("mr.timestamp DESC").
		Limit(uint64(limit)).
		Query()

	if err != nil {
		return []*warehouse.Odds{}, err
	}

	defer rows.Close()

	var odds []*warehouse.Odds

	for rows.Next() {
		var o warehouse.Odds
		var timestamp int64

		if err := rows.Scan(&o.Value, &o.Size, &o.Side, &timestamp); err != nil {
			return []*warehouse.Odds{}, err
		}

		o.Timestamp = time.Unix(timestamp, 0)

		odds = append(odds, &o)
	}

	return odds, nil
}

func (m *marketReader) MarketsByEventID(eventID uint64, q *warehouse.MarketReaderQuery) ([]*warehouse.Market, error) {
	b := m.queryBuilder()

	fmt.Printf("Open connections %d for Event %d\n", m.connection.Stats().OpenConnections, eventID)

	query := b.
		Select("*").
		From("market m").
		Where(sq.Eq{"m.event_id": eventID}).
		OrderBy("name ASC", "exchange ASC")

	if len(q.Market) > 0 {
		query = query.Where(sq.Eq{"name": q.Market})
	}

	if len(q.Exchange) > 0 {
		query = query.Where(sq.Eq{"exchange": q.Exchange})
	}

	rows, err := query.Query()

	if err != nil {
		return []*warehouse.Market{}, err
	}

	defer rows.Close()

	var markets []*warehouse.Market

	for rows.Next() {
		var mk warehouse.Market
		var date int64

		if err := rows.Scan(
			&mk.ID,
			&mk.EventID,
			&date,
			&mk.CompetitionID,
			&mk.SeasonID,
			&mk.Name,
			&mk.Exchange,
		); err != nil {
			return []*warehouse.Market{}, err
		}

		mk.EventDate = time.Unix(date, 0)

		runners, err := m.marketRunners(b, mk.ID, eventID)

		if err != nil {
			return []*warehouse.Market{}, err
		}

		mk.Runners = runners

		markets = append(markets, &mk)
	}

	return markets, nil
}

func (m *marketReader) marketRunners(b sq.StatementBuilderType, marketID string, eventID uint64) ([]*warehouse.Runner, error) {
	fmt.Printf("Open connections %d for Event %d and runners\n", m.connection.Stats().OpenConnections, eventID)

	rows, err := b.
		Select("DISTINCT ON(name) *").
		From("market_runner").
		Where(sq.Eq{"market_id": marketID}).
		Where(sq.Eq{"side": "BACK"}).
		OrderBy("name ASC", "timestamp DESC").
		Query()

	if err != nil {
		return []*warehouse.Runner{}, err
	}

	defer rows.Close()

	var runners []*warehouse.Runner

	for rows.Next() {
		var r warehouse.Runner
		var o warehouse.Odds
		var timestamp int64

		if err := rows.Scan(&r.MarketID, &r.ID, &r.Name, &o.Value, &o.Size, &timestamp, &o.Side); err != nil {
			return []*warehouse.Runner{}, err
		}

		o.Timestamp = time.Unix(timestamp, 0)
		r.BackPrice = &o

		runners = append(runners, &r)
	}

	return runners, nil
}

func (m *marketReader) queryBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(m.connection)
}

func NewMarketReader(connection *sql.DB) warehouse.MarketReader {
	return &marketReader{connection: connection}
}
