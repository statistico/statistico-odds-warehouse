package grpc_test

import (
	"errors"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-odds-warehouse/internal/grpc"
	"github.com/statistico/statistico-odds-warehouse/internal/market"
	"github.com/statistico/statistico-odds-warehouse/internal/mock"
	"github.com/statistico/statistico-proto/statistico-odds-warehouse/go"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestMarketService_MarketSelectionSearch(t *testing.T) {
	t.Run("parses market runners and streams market selection structs", func(t *testing.T) {
		t.Helper()

		repo := new(market.MockRepository)
		logger, hook := test.NewNullLogger()

		service := grpc.NewMarketService(repo, logger)

		server := new(mock.MarketSelectionServer)

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

		req := statisticoproto.MarketRunnerRequest{
			Name:           "MATCH_ODDS",
			RunnerFilter:   &f,
			CompetitionIds: []uint64{1, 2, 3},
			SeasonIds:      []uint64{4, 5, 6},
			DateFrom:       &wrappers.StringValue{Value: "2020-03-12T12:00:00+00:00"},
			DateTo:         &wrappers.StringValue{Value: "2020-03-12T20:00:00+00:00"},
		}

		q := mock2.MatchedBy(func(query *market.RunnerQuery) bool {
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
			return true
		})

		mRunners := []*market.MarketRunner{
			newMarketRunner(),
			newMarketRunner(),
		}

		repo.On("MarketRunners", q).Return(mRunners, nil)

		server.On("Send", mock2.AnythingOfType("*statisticoproto.MarketRunner")).
			Times(2).
			Return(nil)

		err := service.MarketRunnerSearch(&req, server)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, 0, len(hook.Entries))
		repo.AssertExpectations(t)
		server.AssertExpectations(t)
	})

	t.Run("logs error if error return from market repository", func(t *testing.T) {
		t.Helper()

		repo := new(market.MockRepository)
		logger, hook := test.NewNullLogger()

		service := grpc.NewMarketService(repo, logger)

		server := new(mock.MarketSelectionServer)

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

		req := statisticoproto.MarketRunnerRequest{
			Name:           "MATCH_ODDS",
			RunnerFilter:   &f,
			CompetitionIds: []uint64{1, 2, 3},
			SeasonIds:      []uint64{4, 5, 6},
			DateFrom:       &wrappers.StringValue{Value: "2020-03-12T12:00:00+00:00"},
			DateTo:         &wrappers.StringValue{Value: "2020-03-12T20:00:00+00:00"},
		}

		q := mock2.MatchedBy(func(query *market.RunnerQuery) bool {
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
			return true
		})

		repo.On("MarketRunners", q).Return([]*market.MarketRunner{}, errors.New("oh no"))

		server.AssertNotCalled(t, "Send", mock2.AnythingOfType("*statisticoproto.MarketRunner"))

		err := service.MarketRunnerSearch(&req, server)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "rpc error: code = Internal desc = Internal server error", err.Error())
		assert.Equal(t, "Error fetching market runners in market service. oh no", hook.LastEntry().Message)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		repo.AssertExpectations(t)
	})

	t.Run("logs error if error streaming MarketRunner", func(t *testing.T) {
		t.Helper()

		repo := new(market.MockRepository)
		logger, hook := test.NewNullLogger()

		service := grpc.NewMarketService(repo, logger)

		server := new(mock.MarketSelectionServer)

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

		req := statisticoproto.MarketRunnerRequest{
			Name:           "MATCH_ODDS",
			RunnerFilter:   &f,
			CompetitionIds: []uint64{1, 2, 3},
			SeasonIds:      []uint64{4, 5, 6},
			DateFrom:       &wrappers.StringValue{Value: "2020-03-12T12:00:00+00:00"},
			DateTo:         &wrappers.StringValue{Value: "2020-03-12T20:00:00+00:00"},
		}

		q := mock2.MatchedBy(func(query *market.RunnerQuery) bool {
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
			return true
		})

		mRunners := []*market.MarketRunner{
			newMarketRunner(),
			newMarketRunner(),
		}

		repo.On("MarketRunners", q).Return(mRunners, nil)

		server.On("Send", mock2.AnythingOfType("*statisticoproto.MarketRunner")).
			Times(2).
			Return(errors.New("oh no"))

		err := service.MarketRunnerSearch(&req, server)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, "Error streaming market runner back to client. oh no", hook.LastEntry().Message)
		assert.Equal(t, 2, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		repo.AssertExpectations(t)
		server.AssertExpectations(t)
	})
}

func newMarketRunner() *market.MarketRunner {
	return &market.MarketRunner{}
}
