package postgres_test

import (
	"github.com/statistico/statistico-odds-warehouse/internal/app/market"
	"github.com/statistico/statistico-odds-warehouse/internal/app/market/postgres"
	"github.com/statistico/statistico-odds-warehouse/internal/app/test"
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
			{newMarket("1.2729821", "OVER_UNDER_25", time.Unix(1584014400, 0)), 1},
			{newMarket("1.2729822", "OVER_UNDER_25", time.Unix(1584014400, 0)), 2},
			{newMarket("1.2729823", "OVER_UNDER_25", time.Unix(1584014400, 0)), 3},
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
				ID:   423721,
				Name: "Over 2.5 Goals",
				BackPrice: &market.Price{
					Value:     1.95,
					Size:      1591.01,
					Timestamp: time.Unix(1606824710, 0),
				},
			},
			{
				ID:   423721,
				Name: "Under 2.5 Goals",
				LayPrice: &market.Price{
					Value:     2.05,
					Size:      11.55,
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

func TestMarketRepository_MarketRunners(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, []string{"market", "market_runner"})
	repo := postgres.NewMarketRepository(conn)

	t.Run("returns MarketRunner struct filtered by market and runner name", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		insertMultipleMarketsAndRunner(t, repo)

		q := market.RunnerQuery{
			MarketName: "BOTH_TEAMS_TO_SCORE",
			RunnerName: "No",
			Line:       "MAX",
			Side:       "LAY",
		}

		fetched, err := repo.MarketRunners(&q)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, 1, len(fetched))
		assert.Equal(t, "1.567", fetched[0].MarketID)
		assert.Equal(t, "BOTH_TEAMS_TO_SCORE", fetched[0].MarketName)
		assert.Equal(t, uint64(1827711), fetched[0].EventID)
		assert.Equal(t, uint64(8), fetched[0].CompetitionID)
		assert.Equal(t, uint64(17420), fetched[0].SeasonID)
		assert.Equal(t, time.Unix(1583020800, 0), fetched[0].EventDate)
		assert.Equal(t, "betfair", fetched[0].Exchange)
		assert.Equal(t, "1.567", fetched[0].MarketID)
		assert.Equal(t, uint64(423722), fetched[0].RunnerID)
		assert.Equal(t, "No", fetched[0].RunnerName)
		assert.Equal(t, float32(3.05), fetched[0].Price.Value)
		assert.Equal(t, float32(11.55), fetched[0].Price.Size)
		assert.Equal(t, "LAY", fetched[0].Price.Side)
		assert.Equal(t, int64(1605139200), fetched[0].Price.Timestamp.Unix())
	})

	t.Run("returns MarketRunner struct filtered by event date", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		insertMultipleMarketsAndRunner(t, repo)

		from := time.Unix(1584014000, 0)

		q := market.RunnerQuery{
			MarketName: "OVER_UNDER_25",
			RunnerName: "Over 2.5 Goals",
			DateFrom:   &from,
			Side:       "BACK",
		}

		fetched, err := repo.MarketRunners(&q)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, 1, len(fetched))
		assert.Equal(t, "1.234", fetched[0].MarketID)
		assert.Equal(t, "OVER_UNDER_25", fetched[0].MarketName)
		assert.Equal(t, uint64(1827711), fetched[0].EventID)
		assert.Equal(t, uint64(8), fetched[0].CompetitionID)
		assert.Equal(t, uint64(17420), fetched[0].SeasonID)
		assert.Equal(t, time.Unix(1584014400, 0), fetched[0].EventDate)
		assert.Equal(t, "betfair", fetched[0].Exchange)
		assert.Equal(t, "1.234", fetched[0].MarketID)
		assert.Equal(t, uint64(423721), fetched[0].RunnerID)
		assert.Equal(t, "Over 2.5 Goals", fetched[0].RunnerName)
		assert.Equal(t, float32(1.95), fetched[0].Price.Value)
		assert.Equal(t, float32(1591.01), fetched[0].Price.Size)
		assert.Equal(t, "BACK", fetched[0].Price.Side)
		assert.Equal(t, int64(1606824710), fetched[0].Price.Timestamp.Unix())
	})

	t.Run("returns MarketRunner struct filtered price", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		insertMultipleMarketsAndRunner(t, repo)

		q := market.RunnerQuery{
			MarketName: "BOTH_TEAMS_TO_SCORE",
			RunnerName: "Yes",
			Line:       "CLOSING",
			Side:       "LAY",
		}

		fetched, err := repo.MarketRunners(&q)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, 1, len(fetched))
		assert.Equal(t, "1.567", fetched[0].MarketID)
		assert.Equal(t, "BOTH_TEAMS_TO_SCORE", fetched[0].MarketName)
		assert.Equal(t, uint64(1827711), fetched[0].EventID)
		assert.Equal(t, uint64(8), fetched[0].CompetitionID)
		assert.Equal(t, uint64(17420), fetched[0].SeasonID)
		assert.Equal(t, time.Unix(1583020800, 0), fetched[0].EventDate)
		assert.Equal(t, "betfair", fetched[0].Exchange)
		assert.Equal(t, "1.567", fetched[0].MarketID)
		assert.Equal(t, uint64(423721), fetched[0].RunnerID)
		assert.Equal(t, "Yes", fetched[0].RunnerName)
		assert.Equal(t, float32(1.95), fetched[0].Price.Value)
		assert.Equal(t, float32(1591.45), fetched[0].Price.Size)
		assert.Equal(t, "LAY", fetched[0].Price.Side)
		assert.Equal(t, int64(1606839427), fetched[0].Price.Timestamp.Unix())
	})
}

func newMarket(marketID, name string, date time.Time) *market.Market {
	return &market.Market{
		ID:            marketID,
		Name:          name,
		EventID:       1827711,
		CompetitionID: 8,
		SeasonID:      17420,
		EventDate:     date,
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

func insertMultipleMarketsAndRunner(t *testing.T, repo *postgres.MarketRepository) {
	// Event Date: 2020-03-12T12:00:00+00:00
	mk1 := newMarket("1.234", "OVER_UNDER_25", time.Unix(1584014400, 0))

	run1 := []*market.Runner{
		{
			MarketID: mk1.ID,
			ID:       423721,
			Name:     "Over 2.5 Goals",
			BackPrice: &market.Price{
				Value:     1.95,
				Size:      1591.01,
				Timestamp: time.Unix(1606824710, 0),
			},
		},
		{
			MarketID: mk1.ID,
			ID:       423721,
			Name:     "Under 2.5 Goals",
			BackPrice: &market.Price{
				Value:     2.05,
				Size:      11.55,
				Timestamp: time.Unix(1606824710, 0),
			},
		},
	}

	insertMarket(t, repo, mk1)
	insertRunners(t, repo, run1)

	// Event Date: 2020-03-01T00:00:00+00:00
	mk2 := newMarket("1.567", "BOTH_TEAMS_TO_SCORE", time.Unix(1583020800, 0))

	run2 := []*market.Runner{
		{
			MarketID: mk2.ID,
			ID:       423721,
			Name:     "Yes",
			LayPrice: &market.Price{
				Value:     1.95,
				Size:      1591.45,
				Timestamp: time.Unix(1606839427, 0),
			},
		},
		{
			MarketID: mk2.ID,
			ID:       423721,
			Name:     "No",
			LayPrice: &market.Price{
				Value:     2.05,
				Size:      11.55,
				Timestamp: time.Unix(1606824710, 0),
			},
		},
	}

	insertMarket(t, repo, mk2)
	insertRunners(t, repo, run2)

	// Event Date: 2020-03-01T00:00:00+00:00
	mk3 := newMarket("1.567", "BOTH_TEAMS_TO_SCORE", time.Unix(1583020800, 0))

	run3 := []*market.Runner{
		{
			MarketID: mk3.ID,
			ID:       423721,
			Name:     "Yes",
			BackPrice: &market.Price{
				Value:     1.95,
				Size:      1591.01,
				Timestamp: time.Unix(1606824710, 0),
			},
		},
		{
			MarketID: mk3.ID,
			ID:       423722,
			Name:     "No",
			LayPrice: &market.Price{
				Value:     3.05,
				Size:      11.55,
				Timestamp: time.Unix(1605139200, 0),
			},
		},
	}

	insertMarket(t, repo, mk3)
	insertRunners(t, repo, run3)
}
