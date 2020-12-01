package postgres

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/statistico/statistico-odds-warehouse/internal/market"
)

func buildMarketRunnerQuery(q *market.RunnerQuery, b *sq.StatementBuilderType) sq.SelectBuilder {
	query := b.
		Select(
			"m.id",
			"m.name",
			"m.event_id",
			"m.event_date",
			"m.competition_id",
			"m.season_id",
			"m.exchange",
			"m.side",
			"m.timestamp",
			"mr.id",
			"mr.name",
			"mr.price",
			"mr.size",
		).
		From("market m")

	join := sq.Select("DISTINCT on (market_id) *").
		From("market_runner mr").
		Where(sq.Eq{"mr.name": q.Name})

	if q.GreaterThan != nil {
		join = join.Where(sq.Gt{"price": *q.GreaterThan})
	}

	if q.LessThan != nil {
		join = join.Where(sq.Lt{"price": *q.GreaterThan})
	}

	if q.Line == "CLOSING" {
		join = join.OrderBy("mr.timestamp DESC")
	}

	if q.Line == "MAX" {
		join = join.OrderBy("market_id", "mr.price DESC")
	}

	query = query.JoinClause(join.Prefix("JOIN (").Suffix(") as mr ON m.id = mr.market_id"))

	if q.DateFrom != nil {
		query = query.Where(sq.Gt{"m.event_date": q.DateFrom.Unix()})
	}

	if q.DateTo != nil {
		query = query.Where(sq.Lt{"m.event_date": q.DateTo.Unix()})
	}

	if len(q.CompetitionIDs) > 0 {
		query = query.Where(sq.Eq{"m.competition_id": q.CompetitionIDs})
	}

	if len(q.SeasonIDs) > 0 {
		query = query.Where(sq.Eq{"m.season_id": q.SeasonIDs})
	}

	return query
}
