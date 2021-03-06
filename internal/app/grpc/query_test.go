package grpc

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/statistico/statistico-proto/go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_convertMarketSelectionRequest(t *testing.T) {
	t.Run("converts MarketRunnerRequest in RunnerQuery", func(t *testing.T) {
		t.Helper()

		req := statistico.MarketRunnerRequest{
			Market:         "MATCH_ODDS",
			Runner:         "Home",
			Line:           "MAX",
			MinOdds:        &wrappers.FloatValue{Value: 1.95},
			MaxOdds:        &wrappers.FloatValue{Value: 3.55},
			CompetitionIds: []uint64{1, 2, 3},
			SeasonIds:      []uint64{4, 5, 6},
			DateFrom:       &timestamp.Timestamp{Seconds: 1584014400},
			DateTo:         &timestamp.Timestamp{Seconds: 1584014400},
		}

		query, err := convertMarketSelectionRequest(&req)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		a.Equal("MATCH_ODDS", query.MarketName)
		a.Equal("Home", query.RunnerName)
		a.Equal("MAX", query.Line)
		a.Equal([]uint64{1, 2, 3}, query.CompetitionIDs)
		a.Equal([]uint64{4, 5, 6}, query.SeasonIDs)
		a.Equal("2020-03-12T12:00:00Z", query.DateFrom.Format(time.RFC3339))
		a.Equal("2020-03-12T12:00:00Z", query.DateTo.Format(time.RFC3339))
		a.Equal(float32(1.95), *query.GreaterThan)
		a.Equal(float32(3.55), *query.LessThan)
	})

	t.Run("converts MarketRunnerRequest in RunnerQuery handling nullable fields", func(t *testing.T) {
		t.Helper()

		req := statistico.MarketRunnerRequest{
			Market: "MATCH_ODDS",
			Runner: "Home",
			Line:   "MAX",
		}

		query, err := convertMarketSelectionRequest(&req)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		a.Equal("MATCH_ODDS", query.MarketName)
		a.Equal("Home", query.RunnerName)
		a.Equal("MAX", query.Line)
		a.Equal([]uint64(nil), query.CompetitionIDs)
		a.Equal([]uint64(nil), query.SeasonIDs)
		a.Nil(query.DateFrom)
		a.Nil(query.DateTo)
		a.Nil(query.GreaterThan)
		a.Nil(query.LessThan)
	})
}
