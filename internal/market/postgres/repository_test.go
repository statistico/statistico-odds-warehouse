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
	conn, cleanUp := test.GetConnection(t, []string{"market", "market_runner"})
	repo := postgres.NewMarketRepository(conn)

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		marketCounts := []struct {
			Market      *market.Market
			MarketCount int8
			RunnerCount int8
		}{
			{newMarket("1.2729821", "OVER_UNDER_25", "BACK", time.Now()), 1, 2},
			{newMarket("1.2729822", "OVER_UNDER_25", "BACK", time.Now()), 2, 4},
			{newMarket("1.2729823", "OVER_UNDER_25", "BACK", time.Now()), 3, 6},
		}

		for _, tc := range marketCounts {
			insertOverUnderMarket(t, repo, tc.Market)

			var marketCount int8
			var runnerCount int8

			row := conn.QueryRow("select count(*) from market")

			if err := row.Scan(&marketCount); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			row = conn.QueryRow("select count(*) from market_runner")

			if err := row.Scan(&runnerCount); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			assert.Equal(t, tc.MarketCount, marketCount)
			assert.Equal(t, tc.RunnerCount, runnerCount)
		}
	})
}

func newMarket(marketID, name, side string, t time.Time) *market.Market {
	over := market.Runner{
		ID:    423721,
		Name:  "Over 2.5 Goals",
		Price: 1.95,
		Size:  1591.01,
	}

	under := market.Runner{
		ID:    423721,
		Name:  "Under 2.5 Goals",
		Price: 2.05,
		Size:  11.55,
	}

	return &market.Market{
		ID:            marketID,
		Name:          name,
		EventID:       1827711,
		CompetitionID: 8,
		SeasonID:      17420,
		EventDate:     time.Now(),
		Side:          side,
		Exchange:      "betfair",
		Runners: []*market.Runner{
			&over,
			&under,
		},
		Timestamp: t.Unix(),
	}
}

func insertOverUnderMarket(t *testing.T, repo *postgres.MarketRepository, m *market.Market) {
	if err := repo.InsertMarket(m); err != nil {
		t.Errorf("Error when inserting market into the database: %s", err.Error())
	}
}
