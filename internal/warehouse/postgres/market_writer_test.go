package postgres_test

import (
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse/postgres"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMarketRepository_InsertMarket(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, []string{"market"})
	writer := postgres.NewMarketWriter(conn)

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		marketCounts := []struct {
			Market      *warehouse.Market
			MarketCount int8
		}{
			{newMarket("1.2729821", "OVER_UNDER_25", "BETFAIR", time.Unix(1584014400, 0)), 1},
			{newMarket("1.2729822", "OVER_UNDER_25", "BETFAIR", time.Unix(1584014400, 0)), 2},
			{newMarket("1.2729823", "OVER_UNDER_25", "BETFAIR", time.Unix(1584014400, 0)), 3},
		}

		for _, tc := range marketCounts {
			insertMarket(t, writer, tc.Market)

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
	writer := postgres.NewMarketWriter(conn)

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		runners := []*warehouse.Runner{
			{
				ID:   423721,
				Name: "Over 2.5 Goals",
				BackPrice: &warehouse.Odds{
					Value:     1.95,
					Size:      1591.01,
					Timestamp: time.Unix(1606824710, 0),
				},
			},
			{
				ID:   423721,
				Name: "Under 2.5 Goals",
				LayPrice: &warehouse.Odds{
					Value:     2.05,
					Size:      11.55,
					Timestamp: time.Unix(1606824710, 0),
				},
			},
		}

		runnerCounts := []struct {
			Runners     []*warehouse.Runner
			RunnerCount int8
		}{
			{runners, 2},
			{runners, 4},
			{runners, 6},
		}

		for _, tc := range runnerCounts {
			insertRunners(t, writer, tc.Runners)

			var runnerCount int8

			row := conn.QueryRow("select count(*) from market_runner")

			if err := row.Scan(&runnerCount); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			assert.Equal(t, tc.RunnerCount, runnerCount)
		}
	})
}

func newMarket(marketID, name, exchange string, date time.Time) *warehouse.Market {
	return &warehouse.Market{
		ID:            marketID,
		Name:          name,
		EventID:       1827711,
		CompetitionID: 8,
		SeasonID:      17420,
		EventDate:     date,
		Exchange:      exchange,
	}
}

func insertMarket(t *testing.T, repo warehouse.MarketWriter, m *warehouse.Market) {
	if err := repo.InsertMarket(m); err != nil {
		t.Errorf("Error when inserting market into the database: %s", err.Error())
	}
}

func insertRunners(t *testing.T, repo warehouse.MarketWriter, r []*warehouse.Runner) {
	if err := repo.InsertRunners(r); err != nil {
		t.Errorf("Error when inserting market into the database: %s", err.Error())
	}
}

func insertMultipleMarketsAndRunner(t *testing.T, repo warehouse.MarketWriter) {
	// Event Date: 2020-03-12T12:00:00+00:00
	mk1 := newMarket("1.234", "OVER_UNDER_25", "BETFAIR", time.Unix(1584014400, 0))

	run1 := []*warehouse.Runner{
		{
			MarketID: mk1.ID,
			ID:       423721,
			Name:     "OVER",
			BackPrice: &warehouse.Odds{
				Value:     1.95,
				Size:      1591.01,
				Timestamp: time.Unix(1606824710, 0),
			},
		},
		{
			MarketID: mk1.ID,
			ID:       423721,
			Name:     "OVER",
			BackPrice: &warehouse.Odds{
				Value:     1.94,
				Size:      592.61,
				Timestamp: time.Unix(1606824712, 0),
			},
		},
		{
			MarketID: mk1.ID,
			ID:       423721,
			Name:     "OVER",
			BackPrice: &warehouse.Odds{
				Value:     1.95,
				Size:      1591.01,
				Timestamp: time.Unix(1606824714, 0),
			},
		},
		{
			MarketID: mk1.ID,
			ID:       423722,
			Name:     "UNDER",
			BackPrice: &warehouse.Odds{
				Value:     2.05,
				Size:      11.55,
				Timestamp: time.Unix(1606824710, 0),
			},
		},
		{
			MarketID: mk1.ID,
			ID:       423722,
			Name:     "UNDER",
			BackPrice: &warehouse.Odds{
				Value:     2.10,
				Size:      11.55,
				Timestamp: time.Unix(1606824715, 0),
			},
		},
	}

	insertMarket(t, repo, mk1)
	insertRunners(t, repo, run1)

	// Event Date: 2020-03-01T00:00:00+00:00
	mk2 := newMarket("1.567", "BOTH_TEAMS_TO_SCORE", "BETFAIR", time.Unix(1583020800, 0))

	run2 := []*warehouse.Runner{
		{
			MarketID: mk2.ID,
			ID:       423721,
			Name:     "YES",
			LayPrice: &warehouse.Odds{
				Value:     1.95,
				Size:      1591.45,
				Timestamp: time.Unix(1606839427, 0),
			},
		},
		{
			MarketID: mk2.ID,
			ID:       423721,
			Name:     "NO",
			LayPrice: &warehouse.Odds{
				Value:     2.05,
				Size:      11.55,
				Timestamp: time.Unix(1606824710, 0),
			},
		},
	}

	insertMarket(t, repo, mk2)
	insertRunners(t, repo, run2)

	// Event Date: 2020-03-01T00:00:00+00:00
	mk3 := newMarket("1.555", "MATCH_ODDS", "BETFAIR", time.Unix(1583020800, 0))

	run3 := []*warehouse.Runner{
		{
			MarketID: mk3.ID,
			ID:       423721,
			Name:     "YES",
			BackPrice: &warehouse.Odds{
				Value:     1.95,
				Size:      1591.01,
				Timestamp: time.Unix(1606824710, 0),
			},
		},
		{
			MarketID: mk3.ID,
			ID:       423722,
			Name:     "NO",
			LayPrice: &warehouse.Odds{
				Value:     3.05,
				Size:      11.55,
				Timestamp: time.Unix(1605139200, 0),
			},
		},
	}

	insertMarket(t, repo, mk3)
	insertRunners(t, repo, run3)

	// Event Date: 2020-03-12T12:00:00+00:00
	mk4 := newMarket("1.999", "OVER_UNDER_25", "PINNACLE", time.Unix(1584015400, 0))

	run4 := []*warehouse.Runner{
		{
			MarketID: mk4.ID,
			ID:       423721,
			Name:     "OVER",
			BackPrice: &warehouse.Odds{
				Value:     1.95,
				Size:      1591.01,
				Timestamp: time.Unix(1606824710, 0),
			},
		},
		{
			MarketID: mk4.ID,
			ID:       423721,
			Name:     "OVER",
			BackPrice: &warehouse.Odds{
				Value:     1.94,
				Size:      592.61,
				Timestamp: time.Unix(1606824712, 0),
			},
		},
		{
			MarketID: mk4.ID,
			ID:       423721,
			Name:     "OVER",
			BackPrice: &warehouse.Odds{
				Value:     1.90,
				Size:      1591.01,
				Timestamp: time.Unix(1606824704, 0),
			},
		},
		{
			MarketID: mk4.ID,
			ID:       423722,
			Name:     "UNDER",
			BackPrice: &warehouse.Odds{
				Value:     2.12,
				Size:      11.55,
				Timestamp: time.Unix(1606824712, 0),
			},
		},
		{
			MarketID: mk4.ID,
			ID:       423722,
			Name:     "UNDER",
			BackPrice: &warehouse.Odds{
				Value:     2.15,
				Size:      11.55,
				Timestamp: time.Unix(1606824718, 0),
			},
		},
	}

	insertMarket(t, repo, mk4)
	insertRunners(t, repo, run4)
}
