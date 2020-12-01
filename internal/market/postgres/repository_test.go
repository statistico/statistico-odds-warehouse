package postgres_test

import (
	"github.com/statistico/statistico-odds-warehouse/internal/market"
	"github.com/statistico/statistico-odds-warehouse/internal/market/postgres"
	"github.com/statistico/statistico-odds-warehouse/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMarketRepository_InsertMarket(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, []string{"market"})
	repo := postgres.NewMarketRepository(conn)

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		marketCounts := []struct {
			Market      *market.Market
			MarketCount int8
		}{
			{newMarket("1.2729821", "OVER_UNDER_25", "BACK"), 1},
			{newMarket("1.2729822", "OVER_UNDER_25", "BACK"), 2},
			{newMarket("1.2729823", "OVER_UNDER_25", "BACK"), 3},
		}

		for _, tc := range marketCounts {
			insertMarket(t, repo, tc.Market)

			var marketCount int8

			row := conn.QueryRow("select count(*) from market")

			if err := row.Scan(&marketCount); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			assert.Equal(t, tc.MarketCount, marketCount)
		}
	})
}

func TestMarketRepository_InsertRunners(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, []string{"market_runner"})
	repo := postgres.NewMarketRepository(conn)

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		runners := []*market.Runner{
			{
				ID:    423721,
				Name:  "Over 2.5 Goals",
				Price: market.Price{
					Value:     1.95,
					Size:  1591.01,
					Timestamp: time.Unix(1606824710, 0),
				},
			},
			{
				ID:    423721,
				Name:  "Under 2.5 Goals",
				Price: market.Price{
					Value:     2.05,
					Size:  11.55,
					Timestamp: time.Unix(1606824710, 0),
				},
			},
		}

		runnerCounts := []struct {
			Runners     []*market.Runner
			RunnerCount int8
		}{
			{runners, 2},
			{runners, 4},
			{runners, 6},
		}

		for _, tc := range runnerCounts {
			insertRunners(t, repo, tc.Runners)

			var runnerCount int8

			row := conn.QueryRow("select count(*) from market_runner")

			if err := row.Scan(&runnerCount); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			assert.Equal(t, tc.RunnerCount, runnerCount)
		}
	})
}

func newMarket(marketID, name, side string) *market.Market {
	return &market.Market{
		ID:            marketID,
		Name:          name,
		EventID:       1827711,
		CompetitionID: 8,
		SeasonID:      17420,
		EventDate:     time.Now(),
		Side:          side,
		Exchange:      "betfair",
	}
}

func insertMarket(t *testing.T, repo *postgres.MarketRepository, m *market.Market) {
	if err := repo.InsertMarket(m); err != nil {
		t.Errorf("Error when inserting market into the database: %s", err.Error())
	}
}

func insertRunners(t *testing.T, repo *postgres.MarketRepository, r []*market.Runner) {
	if err := repo.InsertRunners(r); err != nil {
		t.Errorf("Error when inserting market into the database: %s", err.Error())
	}
}
