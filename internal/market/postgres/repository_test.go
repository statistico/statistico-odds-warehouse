package postgres_test

import (
	"github.com/statistico/statistico-odds-warehouse/internal/market"
	"github.com/statistico/statistico-odds-warehouse/internal/market/postgres"
	"github.com/statistico/statistico-odds-warehouse/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMarketRepository_Insert(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "market_over_under")
	repo := postgres.NewMarketRepository(conn)

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		tradeCounts := []struct {
			Market *market.OverUnderMarket
			Count int8
		}{
			{newOverUnderMarket(182981, "OVER_UNDER_25", "BACK", time.Now()), 1},
			{newOverUnderMarket(182981, "OVER_UNDER_25", "BACK", time.Now()), 2},
			{newOverUnderMarket(182981, "OVER_UNDER_25", "BACK", time.Now()), 3},
		}

		for _, tc := range tradeCounts {
			insertOverUnderMarket(t, repo, tc.Market)

			var count int8

			row := conn.QueryRow("select count(*) from market_over_under")

			if err := row.Scan(&count); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			assert.Equal(t, tc.Count, count)
		}
	})
}

func TestMarketRepository_InsertBTTSMarket(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "market_btts")
	repo := postgres.NewMarketRepository(conn)

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		tradeCounts := []struct {
			Market *market.BTTSMarket
			Count int8
		}{
			{newBTTSMarket(182981, "BACK", time.Now()), 1},
			{newBTTSMarket(182981, "BACK", time.Now()), 2},
			{newBTTSMarket(182981, "BACK", time.Now()), 3},
		}

		for _, tc := range tradeCounts {
			insertBTTSMarket(t, repo, tc.Market)

			var count int8

			row := conn.QueryRow("select count(*) from market_btts")

			if err := row.Scan(&count); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			assert.Equal(t, tc.Count, count)
		}
	})
}

func TestMarketRepository_MatchOddsMarket(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "market_match_odds")
	repo := postgres.NewMarketRepository(conn)

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		tradeCounts := []struct {
			Market *market.MatchOddsMarket
			Count int8
		}{
			{newMatchOddsMarket(182981, "BACK", time.Now()), 1},
			{newMatchOddsMarket(182981, "BACK", time.Now()), 2},
			{newMatchOddsMarket(182981, "BACK", time.Now()), 3},
		}

		for _, tc := range tradeCounts {
			insertMatchOddsMarket(t, repo, tc.Market)

			var count int8

			row := conn.QueryRow("select count(*) from market_match_odds")

			if err := row.Scan(&count); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			assert.Equal(t, tc.Count, count)
		}
	})
}

func newOverUnderMarket(eventID uint64, name, side string, t time.Time) *market.OverUnderMarket {
	over := market.PriceSize{
		Price: 1.95,
		Size:  1591.01,
	}

	under := market.PriceSize{
		Price: 2.05,
		Size:  1591.01,
	}

	return &market.OverUnderMarket{
		ID:             "1.2729821",
		EventID:        eventID,
		Name:           name,
		Side:           side,
		Exchange:   "betfair",
		Over: over,
		Under: under,
		Timestamp:      t.Unix(),
	}
}

func newBTTSMarket(eventID uint64, side string, t time.Time) *market.BTTSMarket {
	yes := market.PriceSize{
		Price: 1.95,
		Size:  1591.01,
	}

	no := market.PriceSize{
		Price: 2.05,
		Size:  1591.01,
	}

	return &market.BTTSMarket{
		ID:             "1.2729821",
		EventID:        eventID,
		Name:           "BOTH_TEAMS_TO_SCORE",
		Side:           side,
		Exchange:       "betfair",
		Yes:            yes,
		No:             no,
		Timestamp:      t.Unix(),
	}
}

func newMatchOddsMarket(eventID uint64, side string, t time.Time) *market.MatchOddsMarket {
	home := market.PriceSize{
		Price: 1.95,
		Size:  1591.01,
	}

	away := market.PriceSize{
		Price: 2.05,
		Size:  1591.01,
	}

	draw := market.PriceSize{
		Price: 2.05,
		Size:  1591.01,
	}

	return &market.MatchOddsMarket{
		ID:             "1.2729821",
		EventID:        eventID,
		Name:           "MATCH_ODDS",
		Side:           side,
		Exchange:       "betfair",
		Home:            home,
		Away:             away,
		Draw:           draw,
		Timestamp:      t.Unix(),
	}
}

func insertOverUnderMarket(t *testing.T, repo *postgres.MarketRepository, m *market.OverUnderMarket) {
	if err := repo.InsertOverUnderMarket(m); err != nil {
		t.Errorf("Error when inserting market into the database: %s", err.Error())
	}
}

func insertBTTSMarket(t *testing.T, repo *postgres.MarketRepository, m *market.BTTSMarket) {
	if err := repo.InsertBTTSMarket(m); err != nil {
		t.Errorf("Error when inserting market into the database: %s", err.Error())
	}
}

func insertMatchOddsMarket(t *testing.T, repo *postgres.MarketRepository, m *market.MatchOddsMarket) {
	if err := repo.InsertMatchOddsMarket(m); err != nil {
		t.Errorf("Error when inserting market into the database: %s", err.Error())
	}
}
