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

		runners := []*market.Runner{
			{
				ID:    423721,
				Name:  "Over 2.5 Goals",
				Price: 1.95,
				Size:  1591.01,
			},
			{
				ID:    423721,
				Name:  "Under 2.5 Goals",
				Price: 2.05,
				Size:  11.55,
			},
		}

		marketCounts := []struct {
			Market      *market.Market
			MarketCount int8
			RunnerCount int8
		}{
			{newMarket("1.2729821", "OVER_UNDER_25", "BACK", time.Now(), runners), 1, 2},
			{newMarket("1.2729822", "OVER_UNDER_25", "BACK", time.Now(), runners), 2, 4},
			{newMarket("1.2729823", "OVER_UNDER_25", "BACK", time.Now(), runners), 3, 6},
		}

		for _, tc := range marketCounts {
			insertMarket(t, repo, tc.Market)

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

func TestMarketRepository_GetByRunner(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, []string{"market", "market_runner"})
	repo := postgres.NewMarketRepository(conn)

	t.Run("returns markets filtered by runner name", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		insertMultipleMarkets(t, repo)

		q := market.RunnerQuery{
			Name:        "",
			Line:        "",
			GreaterThan: nil,
			LessThan:    nil,
			DateFrom:    nil,
			DateTo:      nil,
		}
	})
}

func newMarket(marketID, name, side string, t time.Time, r []*market.Runner) *market.Market {
	return &market.Market{
		ID:            marketID,
		Name:          name,
		EventID:       1827711,
		CompetitionID: 8,
		SeasonID:      17420,
		EventDate:     time.Now(),
		Side:          side,
		Exchange:      "betfair",
		Runners:       r,
		Timestamp:     t.Unix(),
	}
}

func insertMarket(t *testing.T, repo *postgres.MarketRepository, m *market.Market) {
	if err := repo.InsertMarket(m); err != nil {
		t.Errorf("Error when inserting market into the database: %s", err.Error())
	}
}

func insertMultipleMarkets(t *testing.T, repo *postgres.MarketRepository) {
	runnersOne := []*market.Runner{
		{
			ID:    423721,
			Name:  "Over 2.5 Goals",
			Price: 1.95,
			Size:  1591.01,
		},
		{
			ID:    423722,
			Name:  "Under 2.5 Goals",
			Price: 2.05,
			Size:  11.55,
		},
	}

	marketOne := newMarket(
		"1.2345",
		"OVER_UNDER_25",
		"BACK",
		time.Unix(1606762634, 0),
		runnersOne,
	)

	insertMarket(t, repo, marketOne)

	runnersTwo := []*market.Runner{
		{
			ID:    423721,
			Name:  "Home",
			Price: 1.95,
			Size:  1591.01,
		},
		{
			ID:    423722,
			Name:  "Away",
			Price: 2.05,
			Size:  11.55,
		},
		{
			ID:    423723,
			Name:  "Draw",
			Price: 5.15,
			Size:  1111.55,
		},
	}

	marketTwo := newMarket(
		"1.6789",
		"MATCH_ODDS",
		"BACK",
		time.Unix(1606762640, 0),
		runnersTwo,
	)

	insertMarket(t, repo, marketTwo)

	runnersThree := []*market.Runner{
		{
			ID:    423730,
			Name:  "Yes",
			Price: 1.23,
			Size:  1591.01,
		},
		{
			ID:    423731,
			Name:  "No",
			Price: 3.05,
			Size:  11.55,
		},
	}

	marketThree := newMarket(
		"1.6789",
		"MATCH_ODDS",
		"BACK",
		time.Unix(1606762640, 0),
		runnersThree,
	)

	insertMarket(t, repo, marketThree)

	runnersFour := []*market.Runner{
		{
			ID:    423721,
			Name:  "Home",
			Price: 2.95,
			Size:  1.51,
		},
		{
			ID:    423722,
			Name:  "Away",
			Price: 2.05,
			Size:  11.55,
		},
		{
			ID:    423723,
			Name:  "Draw",
			Price: 5.15,
			Size:  1111.55,
		},
	}

	marketFour := newMarket(
		"1.6789",
		"MATCH_ODDS",
		"BACK",
		time.Unix(1606762632, 0),
		runnersFour,
	)

	insertMarket(t, repo, marketFour)
}
