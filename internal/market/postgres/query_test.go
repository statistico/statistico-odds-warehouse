package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/statistico/statistico-odds-warehouse/internal/market"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_buildMarketRunnerQuery(t *testing.T) {
	t.Run("builds query using RunnerQuery struct", func(t *testing.T) {
		t.Helper()

		gt := float32(1.50)
		lt := float32(3.95)

		from := time.Unix(1606729441, 0)
		to := time.Unix(1606729541, 0)

		q := market.RunnerQuery{
			Name:        "Over 2.5 Goals",
			Line:        "MAX",
			GreaterThan: &gt,
			LessThan:    &lt,
			CompetitionIDs: []uint64{123, 456},
			SeasonIDs: []uint64{999, 000},
			DateFrom:    &from,
			DateTo:      &to,
		}

		b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

		query := buildMarketRunnerQuery(&q, &b)

		expectedSql := "SELECT " +
			"m.id, " +
			"m.name, " +
			"m.event_id, " +
			"m.event_date, " +
			"m.competition_id, " +
			"m.season_id, " +
			"m.exchange, " +
			"m.side, " +
			"m.timestamp, " +
			"mr.id, " +
			"mr.name, " +
			"mr.price, " +
			"mr.size " +
			"FROM " +
			"market m " +
			"JOIN ( " +
			"SELECT " +
			"DISTINCT on (market_id) * " +
			"FROM " +
			"market_runner mr " +
			"WHERE mr.name = $1 AND price > $2 AND price < $3 " +
			"ORDER BY " +
			"market_id, mr.price DESC ) as mr ON " +
			"m.id = mr.market_id " +
			"WHERE m.event_date > $4 AND m.event_date < $5 " +
			"AND m.competition_id IN ($6,$7) " +
			"AND m.season_id IN ($8,$9)"

		sql, _, _ := query.ToSql()

		assert.Equal(t, expectedSql, sql)
	})
}