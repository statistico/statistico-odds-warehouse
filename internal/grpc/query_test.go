package grpc

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/statistico/statistico-proto/statistico-odds-warehouse/go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_convertMarketSelectionRequest(t *testing.T) {
	t.Run("converts MarketRunnerRequest in RunnerQuery", func(t *testing.T) {
		t.Helper()

		f := statisticoproto.RunnerFilter{
			Name:      "Home",
			Line:      statisticoproto.RunnerFilter_MAX,
			Operators: []*statisticoproto.FilterOperator{
				{
					Operator: statisticoproto.FilterOperator_GTE,
					Value: 1.95,
				},
				{
					Operator: statisticoproto.FilterOperator_LTE,
					Value: 3.55,
				},
			},
		}

		r := statisticoproto.MarketRunnerRequest{
			Name:           "MATCH_ODDS",
			RunnerFilter:   &f,
			CompetitionIds: []uint64{1, 2, 3},
			SeasonIds:      []uint64{4, 5, 6},
			DateFrom:       &wrappers.StringValue{Value: "2020-03-12T12:00:00+00:00"},
			DateTo:         &wrappers.StringValue{Value: "2020-03-12T20:00:00+00:00"},
		}

		query, err := convertMarketSelectionRequest(&r)

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
		a.Equal("2020-03-12T20:00:00Z", query.DateTo.Format(time.RFC3339))
		a.Equal(float32(1.95), *query.GreaterThan)
		a.Equal(float32(3.55), *query.LessThan)
	})

	t.Run("converts MarketRunnerRequest in RunnerQuery handling nullable fields", func(t *testing.T) {
		t.Helper()

		f := statisticoproto.RunnerFilter{
			Name:      "Home",
			Line:      statisticoproto.RunnerFilter_CLOSING,
		}

		r := statisticoproto.MarketRunnerRequest{
			Name:           "MATCH_ODDS",
			RunnerFilter:   &f,
		}

		query, err := convertMarketSelectionRequest(&r)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		a.Equal("MATCH_ODDS", query.MarketName)
		a.Equal("Home", query.RunnerName)
		a.Equal("CLOSING", query.Line)
		a.Equal([]uint64(nil), query.CompetitionIDs)
		a.Equal([]uint64(nil), query.SeasonIDs)
		a.Nil(query.DateFrom)
		a.Nil(query.DateTo)
		a.Nil(query.GreaterThan)
		a.Nil(query.LessThan)
	})
}
