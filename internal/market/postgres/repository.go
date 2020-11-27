package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/statistico/statistico-odds-warehouse/internal/market"
)

type MarketRepository struct {
	connection *sql.DB
}

func (r *MarketRepository) InsertOverUnderMarket(m *market.OverUnderMarket) error {
	builder := r.queryBuilder()

	_, err := builder.
		Insert("market_over_under").
		Columns(
			"id",
			"event_id",
			"name",
			"exchange",
			"side",
			"over_price",
			"over_size",
			"under_price",
			"under_size",
			"timestamp",
		).
		Values(
			m.ID,
			m.EventID,
			m.Name,
			m.Exchange,
			m.Side,
			m.Over.Price,
			m.Over.Size,
			m.Under.Price,
			m.Under.Size,
			m.Timestamp,
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
