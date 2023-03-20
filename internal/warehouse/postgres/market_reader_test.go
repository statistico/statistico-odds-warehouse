package postgres_test

import (
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse/postgres"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

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
