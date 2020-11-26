package postgres

import (
	"database/sql"
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/statistico/statistico-odds-warehouse/internal/market"
)

type MarketRepository struct {
	connection *sql.DB
}

func (r *MarketRepository) Insert(m *market.Market) error {
	builder := r.queryBuilder()

	runners, err := json.Marshal(m.Runners)

	if err != nil {
		return err
	}

	_, err = builder.
		Insert("market").
		Columns(
			"id",
			"event_id",
			"name",
			"exchange",
			"side",
			"runners",
			"timestamp",
		).
		Values(
			m.ID,
			m.EventID,
			m.Name,
			m.Exchange,
			m.Side,
			string(runners),
			m.Timestamp,
		).
		Exec()

	return err
}

func (r *MarketRepository) Get(q *market.RepositoryQuery) ([]*market.Market, error) {
	builder := r.queryBuilder()

	query := builder.Select("market.*").From("market")

	rows, err := buildQuery(query, q).Query()

	if err != nil {
		return []*market.Market{}, err
	}

	return rowsToMarketSlice(rows)
}

func buildQuery(b sq.SelectBuilder, q *market.RepositoryQuery) sq.SelectBuilder {
	if q.EventID != nil {
		b = b.Where(sq.Eq{"event_id": *q.EventID})
	}

	if q.MarketName != nil {
		b = b.Where(sq.Eq{"name": *q.MarketName})
	}

	if q.Side != nil {
		b = b.Where(sq.Eq{"side": *q.Side})
	}

	if q.SortBy != nil && *q.SortBy == "timestamp_asc" {
		b = b.OrderBy("timestamp ASC")
	}

	if q.SortBy != nil && *q.SortBy == "timestamp_desc" {
		b = b.OrderBy("timestamp DESC")
	}

	if q.SortBy == nil {
		b = b.OrderBy("timestamp DESC")
	}

	return b
}

func rowsToMarketSlice(rows *sql.Rows) ([]*market.Market, error) {
	defer rows.Close()

	var markets []*market.Market

	for rows.Next() {
		var m market.Market
		var runners string

		err := rows.Scan(
			&m.ID,
			&m.EventID,
			&m.Name,
			&m.Exchange,
			&m.Side,
			&runners,
			&m.Timestamp,
		)

		if err != nil {
			return markets, err
		}

		err = json.Unmarshal([]byte(runners), &m.Runners)

		if err != nil {
			return markets, err
		}

		markets = append(markets, &m)
	}

	return markets, nil
}

func (r *MarketRepository) queryBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(r.connection)
}

func NewMarketRepository(connection *sql.DB) *MarketRepository {
	return &MarketRepository{connection: connection}
}
