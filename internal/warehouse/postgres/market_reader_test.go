package postgres_test

import (
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse/postgres"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMarketReader_MarketsByEventID(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, []string{"market", "market_runner"})
	writer := postgres.NewMarketWriter(conn)
	reader := postgres.NewMarketReader(conn)

	t.Run("returns a list of markets for an event filtered by market name", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		insertMultipleMarketsAndRunner(t, writer)
		
		q := warehouse.MarketReaderQuery{Market: []string{"OVER_UNDER_25"}}

		markets, err := reader.MarketsByEventID(1827711, &q)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		market := markets[0]
		r1 := market.Runners[0]
		r2 := market.Runners[1]

		a.Equal(2, len(markets))
		a.Equal("1.234", market.ID)
		a.Equal("OVER_UNDER_25", market.Name)
		a.Equal(uint64(1827711), market.EventID)
		a.Equal(uint64(8), market.CompetitionID)
		a.Equal(uint64(17420), market.SeasonID)
		a.Equal(time.Unix(1584014400, 0), market.EventDate)
		a.Equal("BETFAIR", market.Exchange)
		a.Equal(2, len(market.Runners))

		a.Equal(uint64(423721), r1.ID)
		a.Equal("1.234", r1.MarketID)
		a.Equal("OVER", r1.Name)
		a.Equal(warehouse.Odds{
			Value:     1.95,
			Size:      1591.01,
			Side:      "BACK",
			Timestamp: time.Unix(1606824714, 0),
		}, *r1.BackPrice)

		a.Equal(uint64(423722), r2.ID)
		a.Equal("1.234", r2.MarketID)
		a.Equal("UNDER", r2.Name)
		a.Equal(warehouse.Odds{
			Value:     2.10,
			Size:      11.55,
			Side:      "BACK",
			Timestamp: time.Unix(1606824715, 0),
		}, *r2.BackPrice)
	})

	t.Run("returns a list of markets for an event filtered by exchange", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		insertMultipleMarketsAndRunner(t, writer)

		q := warehouse.MarketReaderQuery{Exchange: []string{"PINNACLE"}}

		markets, err := reader.MarketsByEventID(1827711, &q)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		market := markets[0]
		r1 := market.Runners[0]
		r2 := market.Runners[1]

		a.Equal(1, len(markets))
		a.Equal("1.999", market.ID)
		a.Equal("OVER_UNDER_25", market.Name)
		a.Equal(uint64(1827711), market.EventID)
		a.Equal(uint64(8), market.CompetitionID)
		a.Equal(uint64(17420), market.SeasonID)
		a.Equal(time.Unix(1584015400, 0), market.EventDate)
		a.Equal("PINNACLE", market.Exchange)
		a.Equal(2, len(market.Runners))

		a.Equal(uint64(423721), r1.ID)
		a.Equal("1.999", r1.MarketID)
		a.Equal("OVER", r1.Name)
		a.Equal(warehouse.Odds{
			Value:     1.94,
			Size:      592.61,
			Side:      "BACK",
			Timestamp: time.Unix(1606824712, 0),
		}, *r1.BackPrice)

		a.Equal(uint64(423722), r2.ID)
		a.Equal("1.999", r2.MarketID)
		a.Equal("UNDER", r2.Name)
		a.Equal(warehouse.Odds{
			Value:     2.15,
			Size:      11.55,
			Side:      "BACK",
			Timestamp: time.Unix(1606824718, 0),
		}, *r2.BackPrice)
	})
}

func TestMarketReader_ExchangeMarketRunnerOdds(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, []string{"market", "market_runner"})
	writer := postgres.NewMarketWriter(conn)
	reader := postgres.NewMarketReader(conn)

	t.Run("return a slice of odds struct", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		insertMultipleMarketsAndRunner(t, writer)

		odds, err := reader.ExchangeMarketRunnerOdds(1827711, "OVER_UNDER_25", "OVER", "BETFAIR", 2)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		a.Equal(2, len(odds))

		a.Equal(float32(1.95), odds[0].Value)
		a.Equal(float32(1591.01), odds[0].Size)
		a.Equal("BACK", odds[0].Side)
		a.Equal(time.Unix(1606824714, 0), odds[0].Timestamp)

		a.Equal(float32(1.94), odds[1].Value)
		a.Equal(float32(592.61), odds[1].Size)
		a.Equal("BACK", odds[1].Side)
		a.Equal(time.Unix(1606824712, 0), odds[1].Timestamp)
	})

	t.Run("returns an empty slice of odds struct", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		insertMultipleMarketsAndRunner(t, writer)

		odds, err := reader.ExchangeMarketRunnerOdds(1827754, "OVER_UNDER_25", "OVER", "BETFAIR", 2)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		a.Equal(0, len(odds))
	})
}
