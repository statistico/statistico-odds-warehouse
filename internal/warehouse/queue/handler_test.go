package queue_test

import (
	"errors"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse/queue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestHandler_Handle(t *testing.T) {
	t.Run("parses over under market and persist via the repository", func(t *testing.T) {
		t.Helper()

		writer := new(queue.MockMarketWriter)
		handler := queue.NewMarketHandler(writer)

		mk := &queue.EventMarket{
			ID:            "1.2818721",
			EventID:       148192,
			CompetitionID: 8,
			SeasonID:      17420,
			EventDate:     1583971200,
			Name:          "MATCH_GOALS",
			Exchange:      "betfair",
			Runners: []*queue.Runner{
				{
					ID:   "472671",
					Name: "Over 2.5 Goals",
					BackPrices: []queue.PriceSize{
						{
							Price: 1.95,
							Size:  156.91,
						},
					},
					LayPrices: []queue.PriceSize{
						{
							Price: 1.95,
							Size:  156.91,
						},
					},
				},
				{
					ID:   "472672",
					Name: "Under 2.5 Goals",
					BackPrices: []queue.PriceSize{
						{
							Price: 2.05,
							Size:  1.92,
						},
					},
					LayPrices: []queue.PriceSize{
						{
							Price: 2.05,
							Size:  1.92,
						},
					},
				},
			},
			Timestamp: 1583971200,
		}

		mkt := mock.MatchedBy(func(m *warehouse.Market) bool {
			assert.Equal(t, "1.2818721", m.ID)
			assert.Equal(t, uint64(148192), m.EventID)
			assert.Equal(t, uint64(8), m.CompetitionID)
			assert.Equal(t, uint64(17420), m.SeasonID)
			assert.Equal(t, "MATCH_GOALS", m.Name)
			assert.Equal(t, "betfair", m.Exchange)
			assert.Equal(t, time.Unix(1583971200, 0), m.EventDate)
			return true
		})

		run := mock.MatchedBy(func(r []*warehouse.Runner) bool {
			assert.Equal(t, "472671", r[0].ID)
			assert.Equal(t, "Over 2.5 Goals", r[0].Name)
			assert.Equal(t, float32(1.95), r[0].BackPrice.Value)
			assert.Equal(t, float32(156.91), r[0].BackPrice.Size)
			assert.Equal(t, time.Unix(1583971200, 0), r[0].BackPrice.Timestamp)
			assert.Equal(t, float32(1.95), r[0].LayPrice.Value)
			assert.Equal(t, float32(156.91), r[0].LayPrice.Size)
			assert.Equal(t, time.Unix(1583971200, 0), r[0].LayPrice.Timestamp)
			assert.Equal(t, "472672", r[1].ID)
			assert.Equal(t, "Under 2.5 Goals", r[1].Name)
			assert.Equal(t, float32(2.05), r[1].BackPrice.Value)
			assert.Equal(t, float32(1.92), r[1].BackPrice.Size)
			assert.Equal(t, time.Unix(1583971200, 0), r[1].BackPrice.Timestamp)
			assert.Equal(t, float32(2.05), r[1].LayPrice.Value)
			assert.Equal(t, float32(1.92), r[1].LayPrice.Size)
			assert.Equal(t, time.Unix(1583971200, 0), r[1].LayPrice.Timestamp)
			return true
		})

		writer.On("InsertMarket", mkt).Return(nil)
		writer.On("InsertRunners", run).Return(nil)

		err := handler.Handle(mk)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		writer.AssertExpectations(t)
	})

	t.Run("returns error if returned by repository", func(t *testing.T) {
		t.Helper()

		writer := new(queue.MockMarketWriter)
		handler := queue.NewMarketHandler(writer)

		mk := &queue.EventMarket{
			ID:            "1.2818721",
			EventID:       148192,
			CompetitionID: 8,
			SeasonID:      17420,
			EventDate:     1583971200,
			Name:          "OVER_UNDER_25",
			Exchange:      "betfair",
			Runners: []*queue.Runner{
				{
					ID:   "472671",
					Name: "Over 2.5 Goals",
					BackPrices: []queue.PriceSize{
						{
							Price: 1.951827161651,
							Size:  156.91671761,
						},
					},
				},
				{
					ID:   "472672",
					Name: "Under 2.5 Goals",
					BackPrices: []queue.PriceSize{
						{
							Price: 2.0555411311,
							Size:  1.92141241,
						},
					},
				},
			},
			Timestamp: 1583971200,
		}

		mkt := mock.MatchedBy(func(m *warehouse.Market) bool {
			assert.Equal(t, "1.2818721", m.ID)
			assert.Equal(t, uint64(148192), m.EventID)
			assert.Equal(t, uint64(8), m.CompetitionID)
			assert.Equal(t, uint64(17420), m.SeasonID)
			assert.Equal(t, "OVER_UNDER_25", m.Name)
			assert.Equal(t, "betfair", m.Exchange)
			assert.Equal(t, time.Unix(1583971200, 0), m.EventDate)
			return true
		})

		writer.On("InsertMarket", mkt).Return(errors.New("oh no"))
		writer.AssertNotCalled(t, "InsertRunners")

		err := handler.Handle(mk)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "oh no", err.Error())

		writer.AssertExpectations(t)
	})

	t.Run("parse runners name by sort priority for MATCH_ODDS market", func(t *testing.T) {
		t.Helper()

		writer := new(queue.MockMarketWriter)
		handler := queue.NewMarketHandler(writer)

		mk := &queue.EventMarket{
			ID:            "1.2818721",
			EventID:       148192,
			CompetitionID: 8,
			SeasonID:      17420,
			EventDate:     1583971200,
			Name:          "MATCH_ODDS",
			Exchange:      "betfair",
			Runners: []*queue.Runner{
				{
					ID:   "472671",
					Name: "Home",
					BackPrices: []queue.PriceSize{
						{
							Price: 1.95581981,
							Size:  156.9198171,
						},
					},
				},
				{
					ID:   "472672",
					Name: "Away",
					BackPrices: []queue.PriceSize{
						{
							Price: 2.05091981,
							Size:  1.92719817,
						},
					},
				},
				{
					ID:   "472673",
					Name: "Draw",
					BackPrices: []queue.PriceSize{
						{
							Price: 3.051111,
							Size:  0.989819,
						},
					},
				},
			},
			Timestamp: 1583971200,
		}

		mkt := mock.MatchedBy(func(m *warehouse.Market) bool {
			assert.Equal(t, "1.2818721", m.ID)
			assert.Equal(t, uint64(148192), m.EventID)
			assert.Equal(t, uint64(8), m.CompetitionID)
			assert.Equal(t, uint64(17420), m.SeasonID)
			assert.Equal(t, "MATCH_ODDS", m.Name)
			assert.Equal(t, "betfair", m.Exchange)
			assert.Equal(t, time.Unix(1583971200, 0), m.EventDate)
			return true
		})

		run := mock.MatchedBy(func(r []*warehouse.Runner) bool {
			assert.Equal(t, "472671", r[0].ID)
			assert.Equal(t, "Home", r[0].Name)
			assert.Equal(t, float32(1.96), r[0].BackPrice.Value)
			assert.Equal(t, float32(156.92), r[0].BackPrice.Size)
			assert.Equal(t, time.Unix(1583971200, 0), r[0].BackPrice.Timestamp)
			assert.Equal(t, "472672", r[1].ID)
			assert.Equal(t, "Away", r[1].Name)
			assert.Equal(t, float32(2.05), r[1].BackPrice.Value)
			assert.Equal(t, float32(1.93), r[1].BackPrice.Size)
			assert.Equal(t, time.Unix(1583971200, 0), r[1].BackPrice.Timestamp)
			assert.Equal(t, "Draw", r[2].Name)
			assert.Equal(t, float32(3.05), r[2].BackPrice.Value)
			assert.Equal(t, float32(0.99), r[2].BackPrice.Size)
			assert.Equal(t, time.Unix(1583971200, 0), r[2].BackPrice.Timestamp)
			return true
		})

		writer.On("InsertMarket", mkt).Return(nil)
		writer.On("InsertRunners", run).Return(nil)

		err := handler.Handle(mk)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		writer.AssertExpectations(t)
	})

	t.Run("runner is not persisted if runner prices slice is empty", func(t *testing.T) {
		t.Helper()

		writer := new(queue.MockMarketWriter)
		handler := queue.NewMarketHandler(writer)

		mk := &queue.EventMarket{
			ID:            "1.2818721",
			EventID:       148192,
			CompetitionID: 8,
			SeasonID:      17420,
			EventDate:     1583971200,
			Name:          "MATCH_ODDS",
			Exchange:      "betfair",
			Runners: []*queue.Runner{
				{
					ID:        "472671",
					Name:      "Home",
					LayPrices: []queue.PriceSize{},
				},
				{
					ID:   "472672",
					Name: "Away",
					LayPrices: []queue.PriceSize{
						{
							Price: 2.05091981,
							Size:  1.92719817,
						},
					},
				},
				{
					ID:   "472673",
					Name: "Draw",
					LayPrices: []queue.PriceSize{
						{
							Price: 3.051111,
							Size:  0.989819,
						},
					},
				},
			},
			Timestamp: 1583971200,
		}

		mkt := mock.MatchedBy(func(m *warehouse.Market) bool {
			assert.Equal(t, "1.2818721", m.ID)
			assert.Equal(t, uint64(148192), m.EventID)
			assert.Equal(t, uint64(8), m.CompetitionID)
			assert.Equal(t, uint64(17420), m.SeasonID)
			assert.Equal(t, "MATCH_ODDS", m.Name)
			assert.Equal(t, "betfair", m.Exchange)
			assert.Equal(t, time.Unix(1583971200, 0), m.EventDate)
			return true
		})

		run := mock.MatchedBy(func(r []*warehouse.Runner) bool {
			assert.Equal(t, "472672", r[0].ID)
			assert.Equal(t, "Away", r[0].Name)
			assert.Equal(t, float32(2.05), r[0].LayPrice.Value)
			assert.Equal(t, float32(1.93), r[0].LayPrice.Size)
			assert.Equal(t, time.Unix(1583971200, 0), r[0].LayPrice.Timestamp)
			assert.Equal(t, "Draw", r[1].Name)
			assert.Equal(t, float32(3.05), r[1].LayPrice.Value)
			assert.Equal(t, float32(0.99), r[1].LayPrice.Size)
			assert.Equal(t, time.Unix(1583971200, 0), r[1].LayPrice.Timestamp)
			return true
		})

		writer.On("InsertMarket", mkt).Return(nil)
		writer.On("InsertRunners", run).Return(nil)

		err := handler.Handle(mk)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		writer.AssertExpectations(t)
	})
}
