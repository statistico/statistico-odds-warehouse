package grpc_test

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse"
	g "github.com/statistico/statistico-odds-warehouse/internal/warehouse/grpc"
	statistico "github.com/statistico/statistico-proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestMarketService_GetExchangeOdds(t *testing.T) {
	t.Run("parses odds from reader and streams exchange odds structs", func(t *testing.T) {
		t.Helper()

		reader := new(MockMarketReader)
		logger, hook := test.NewNullLogger()

		service := g.NewMarketService(reader, logger)

		server := new(MockExchangeOddsServer)

		req := statistico.ExchangeOddsRequest{
			EventId:  15691981,
			Market:   "MATCH_ODDS",
			Runner:   "HOME",
			Exchange: "BETFAIR",
			Limit:    1,
		}

		odds := []*warehouse.Odds{
			{
				Value:     1.65,
				Size:      12301.55,
				Side:      "BACK",
				Timestamp: time.Unix(1606824714, 0),
			},
		}

		reader.On("ExchangeMarketRunnerOdds", req.EventId, req.Market, req.Runner, req.Exchange, req.Limit).Return(odds, nil)

		oddsResponse := mock.MatchedBy(func(o *statistico.ExchangeOdds) bool {
			assert.Equal(t, float32(1.65), o.Price)
			assert.Equal(t, uint64(1606824714), o.Timestamp)
			return true
		})

		server.On("Send", oddsResponse).
			Times(1).
			Return(nil)

		err := service.GetExchangeOdds(&req, server)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, 0, len(hook.Entries))
		reader.AssertExpectations(t)
		server.AssertExpectations(t)
	})

	t.Run("logs error if error return from market reader", func(t *testing.T) {
		t.Helper()

		reader := new(MockMarketReader)
		logger, hook := test.NewNullLogger()

		service := g.NewMarketService(reader, logger)

		server := new(MockExchangeOddsServer)

		req := statistico.ExchangeOddsRequest{
			EventId:  15691981,
			Market:   "MATCH_ODDS",
			Runner:   "HOME",
			Exchange: "BETFAIR",
			Limit:    1,
		}

		reader.On("ExchangeMarketRunnerOdds", req.EventId, req.Market, req.Runner, req.Exchange, req.Limit).Return([]*warehouse.Odds{}, errors.New("oh no"))

		server.AssertNotCalled(t, "Send", mock.AnythingOfType("*statistico.ExchangeOdds"))

		err := service.GetExchangeOdds(&req, server)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "rpc error: code = Internal desc = internal server error", err.Error())
		assert.Equal(t, "error fetching odds from reader: oh no", hook.LastEntry().Message)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		reader.AssertExpectations(t)
	})

	t.Run("logs error if error streaming MarketRunner", func(t *testing.T) {
		t.Helper()

		reader := new(MockMarketReader)
		logger, hook := test.NewNullLogger()

		service := g.NewMarketService(reader, logger)

		server := new(MockExchangeOddsServer)

		req := statistico.ExchangeOddsRequest{
			EventId:  15691981,
			Market:   "MATCH_ODDS",
			Runner:   "HOME",
			Exchange: "BETFAIR",
			Limit:    1,
		}

		odds := []*warehouse.Odds{
			{
				Value:     1.65,
				Size:      12301.55,
				Side:      "BACK",
				Timestamp: time.Unix(1606824714, 0),
			},
		}

		reader.On("ExchangeMarketRunnerOdds", req.EventId, req.Market, req.Runner, req.Exchange, req.Limit).Return(odds, nil)

		oddsResponse := mock.MatchedBy(func(o *statistico.ExchangeOdds) bool {
			assert.Equal(t, float32(1.65), o.Price)
			assert.Equal(t, uint64(1606824714), o.Timestamp)
			return true
		})

		server.On("Send", oddsResponse).
			Times(1).
			Return(errors.New("oh no"))

		err := service.GetExchangeOdds(&req, server)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, "error streaming exchange odds back to client: oh no", hook.LastEntry().Message)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		reader.AssertExpectations(t)
		server.AssertExpectations(t)
	})
}

type MockMarketReader struct {
	mock.Mock
}

func (m *MockMarketReader) ExchangeMarketRunnerOdds(eventID uint64, market, runner, exchange string, limit uint32) ([]*warehouse.Odds, error) {
	args := m.Called(eventID, market, runner, exchange, limit)
	return args.Get(0).([]*warehouse.Odds), args.Error(1)
}

func (m *MockMarketReader) MarketsByEventID(eventID uint64, q *warehouse.MarketReaderQuery) ([]*warehouse.Market, error) {
	args := m.Called(eventID, q)
	return args.Get(0).([]*warehouse.Market), args.Error(1)
}

type MockExchangeOddsServer struct {
	mock.Mock
	grpc.ServerStream
}

func (m *MockExchangeOddsServer) Send(e *statistico.ExchangeOdds) error {
	args := m.Called(e)
	return args.Error(0)
}
