package market_test

import (
	"errors"
	"github.com/statistico/statistico-odds-warehouse/internal/market"
	"github.com/statistico/statistico-odds-warehouse/internal/queue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	t.Run("parses over under market and persist via the repository", func(t *testing.T) {
		t.Helper()

		repo := new(market.MockRepository)
		handler := market.NewHandler(repo)

		mk := &queue.Market{
			ID:       "1.2818721",
			EventID:  148192,
			Name:     "OVER_UNDER_25",
			Side:     "BACK",
			Exchange: "betfair",
			Runners: []*queue.Runner{
				{
					ID:   472671,
					Name: "Over 2.5 Goals",
					Prices: []queue.PriceSize{
						{
							Price: 1.95,
							Size:  156.91,
						},
					},
				},
				{
					ID:   472671,
					Name: "Under 2.5 Goals",
					Prices: []queue.PriceSize{
						{
							Price: 2.05,
							Size:  1.92,
						},
					},
				},
			},
			Timestamp: 1583971200,
		}

		mkt := mock.MatchedBy(func(m *market.OverUnderMarket) bool {
			assert.Equal(t, "1.2818721", m.ID)
			assert.Equal(t, uint64(148192), m.EventID)
			assert.Equal(t, "OVER_UNDER_25", m.Name)
			assert.Equal(t, "BACK", m.Side)
			assert.Equal(t, "betfair", m.Exchange)
			assert.Equal(t, "betfair", m.Exchange)
			assert.Equal(t, float32(1.95), m.Over.Price)
			assert.Equal(t, float32(156.91), m.Over.Size)
			assert.Equal(t, float32(2.05), m.Under.Price)
			assert.Equal(t, float32(1.92), m.Under.Size)
			assert.Equal(t, int64(1583971200), m.Timestamp)
			return true
		})

		repo.On("InsertOverUnderMarket", mkt).Return(nil)

		err := handler.Handle(mk)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		repo.AssertExpectations(t)
	})

	t.Run("parses both teams to score market and persist via the repository", func(t *testing.T) {
		t.Helper()

		repo := new(market.MockRepository)
		handler := market.NewHandler(repo)

		mk := &queue.Market{
			ID:       "1.2818721",
			EventID:  148192,
			Name:     "BOTH_TEAMS_TO_SCORE",
			Side:     "BACK",
			Exchange: "betfair",
			Runners: []*queue.Runner{
				{
					ID:   472671,
					Name: "Yes",
					Prices: []queue.PriceSize{
						{
							Price: 1.95,
							Size:  156.91,
						},
					},
				},
				{
					ID:   472671,
					Name: "No",
					Prices: []queue.PriceSize{
						{
							Price: 2.05,
							Size:  1.92,
						},
					},
				},
			},
			Timestamp: 1583971200,
		}

		mkt := mock.MatchedBy(func(m *market.BTTSMarket) bool {
			assert.Equal(t, "1.2818721", m.ID)
			assert.Equal(t, uint64(148192), m.EventID)
			assert.Equal(t, "BOTH_TEAMS_TO_SCORE", m.Name)
			assert.Equal(t, "BACK", m.Side)
			assert.Equal(t, "betfair", m.Exchange)
			assert.Equal(t, "betfair", m.Exchange)
			assert.Equal(t, float32(1.95), m.Yes.Price)
			assert.Equal(t, float32(156.91), m.Yes.Size)
			assert.Equal(t, float32(2.05), m.No.Price)
			assert.Equal(t, float32(1.92), m.No.Size)
			assert.Equal(t, int64(1583971200), m.Timestamp)
			return true
		})

		repo.On("InsertBTTSMarket", mkt).Return(nil)

		err := handler.Handle(mk)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		repo.AssertExpectations(t)
	})

	t.Run("parses match odds market and persist via the repository", func(t *testing.T) {
		t.Helper()

		repo := new(market.MockRepository)
		handler := market.NewHandler(repo)

		mk := &queue.Market{
			ID:       "1.2818721",
			EventID:  148192,
			Name:     "MATCH_ODDS",
			Side:     "BACK",
			Exchange: "betfair",
			Runners: []*queue.Runner{
				{
					ID:   472671,
					Name: "West Ham United",
					Prices: []queue.PriceSize{
						{
							Price: 1.95,
							Size:  156.91,
						},
					},
				},
				{
					ID:   472671,
					Name: "Chelsea",
					Prices: []queue.PriceSize{
						{
							Price: 2.05,
							Size:  1.92,
						},
					},
				},
				{
					ID:   472671,
					Name: "The Draw",
					Prices: []queue.PriceSize{
						{
							Price: 5.51,
							Size:  224.12,
						},
					},
				},
			},
			Timestamp: 1583971200,
		}

		mkt := mock.MatchedBy(func(m *market.MatchOddsMarket) bool {
			assert.Equal(t, "1.2818721", m.ID)
			assert.Equal(t, uint64(148192), m.EventID)
			assert.Equal(t, "MATCH_ODDS", m.Name)
			assert.Equal(t, "BACK", m.Side)
			assert.Equal(t, "betfair", m.Exchange)
			assert.Equal(t, "betfair", m.Exchange)
			assert.Equal(t, float32(1.95), m.Home.Price)
			assert.Equal(t, float32(156.91), m.Home.Size)
			assert.Equal(t, float32(2.05), m.Away.Price)
			assert.Equal(t, float32(1.92), m.Away.Size)
			assert.Equal(t, float32(5.51), m.Draw.Price)
			assert.Equal(t, float32(224.12), m.Draw.Size)
			assert.Equal(t, int64(1583971200), m.Timestamp)
			return true
		})

		repo.On("InsertMatchOddsMarket", mkt).Return(nil)

		err := handler.Handle(mk)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		repo.AssertExpectations(t)
	})

	t.Run("returns error if market is not supported", func(t *testing.T) {
		t.Helper()

		repo := new(market.MockRepository)
		handler := market.NewHandler(repo)

		mk := &queue.Market{
			ID:        "1.2818721",
			EventID:   148192,
			Name:      "OVER_UNDER_111625",
			Side:      "BACK",
			Exchange:  "betfair",
			Timestamp: 1583971200,
		}

		repo.AssertNotCalled(t, "InsertOverUnderMarket")

		err := handler.Handle(mk)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "market OVER_UNDER_111625 is not supported", err.Error())
	})

	t.Run("returns error if returned by repository", func(t *testing.T) {
		t.Helper()

		repo := new(market.MockRepository)
		handler := market.NewHandler(repo)

		mk := &queue.Market{
			ID:       "1.2818721",
			EventID:  148192,
			Name:     "OVER_UNDER_135_CORNR",
			Side:     "BACK",
			Exchange: "betfair",
			Runners: []*queue.Runner{
				{
					ID:   472671,
					Name: "Over 13.5 Corners",
					Prices: []queue.PriceSize{
						{
							Price: 1.95,
							Size:  156.91,
						},
					},
				},
				{
					ID:   472671,
					Name: "Under 13.5 Corners",
					Prices: []queue.PriceSize{
						{
							Price: 2.05,
							Size:  1.92,
						},
					},
				},
			},
			Timestamp: 1583971200,
		}

		mkt := mock.MatchedBy(func(m *market.OverUnderMarket) bool {
			assert.Equal(t, "1.2818721", m.ID)
			assert.Equal(t, uint64(148192), m.EventID)
			assert.Equal(t, "OVER_UNDER_135_CORNR", m.Name)
			assert.Equal(t, "BACK", m.Side)
			assert.Equal(t, "betfair", m.Exchange)
			assert.Equal(t, "betfair", m.Exchange)
			assert.Equal(t, float32(1.95), m.Over.Price)
			assert.Equal(t, float32(156.91), m.Over.Size)
			assert.Equal(t, float32(2.05), m.Under.Price)
			assert.Equal(t, float32(1.92), m.Under.Size)
			assert.Equal(t, int64(1583971200), m.Timestamp)
			return true
		})

		repo.On("InsertOverUnderMarket", mkt).Return(errors.New("oh no"))

		err := handler.Handle(mk)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "oh no", err.Error())

		repo.AssertExpectations(t)
	})
}
