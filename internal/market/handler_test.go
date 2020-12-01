package market_test

import (
	"errors"
	"github.com/statistico/statistico-odds-warehouse/internal/market"
	"github.com/statistico/statistico-odds-warehouse/internal/queue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestHandler_Handle(t *testing.T) {
	t.Run("parses over under market and persist via the repository", func(t *testing.T) {
		t.Helper()

		repo := new(market.MockRepository)
		handler := market.NewHandler(repo)

		mk := &queue.Market{
			ID:            "1.2818721",
			EventID:       148192,
			CompetitionID: 8,
			SeasonID:      17420,
			EventDate:     "2020-11-28T12:00:00+00:00",
			Name:          "OVER_UNDER_25",
			Side:          "BACK",
			Exchange:      "betfair",
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
					ID:   472672,
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

		mkt := mock.MatchedBy(func(m *market.Market) bool {
			date, _ := time.Parse(time.RFC3339, "2020-11-28T12:00:00+00:00")

			assert.Equal(t, "1.2818721", m.ID)
			assert.Equal(t, uint64(148192), m.EventID)
			assert.Equal(t, uint64(8), m.CompetitionID)
			assert.Equal(t, uint64(17420), m.SeasonID)
			assert.Equal(t, "OVER_UNDER_25", m.Name)
			assert.Equal(t, "BACK", m.Side)
			assert.Equal(t, "betfair", m.Exchange)
			assert.Equal(t, date, m.EventDate)
			return true
		})

		run := mock.MatchedBy(func(r []*market.Runner) bool {
			assert.Equal(t, uint64(472671), r[0].ID)
			assert.Equal(t, "Over 2.5 Goals", r[0].Name)
			assert.Equal(t, float32(1.95), r[0].Price.Value)
			assert.Equal(t, float32(156.91), r[0].Price.Size)
			assert.Equal(t, time.Unix(1583971200, 0), r[0].Price.Timestamp)
			assert.Equal(t, uint64(472672), r[1].ID)
			assert.Equal(t, "Under 2.5 Goals", r[1].Name)
			assert.Equal(t, float32(2.05), r[1].Price.Value)
			assert.Equal(t, float32(1.92), r[1].Price.Size)
			assert.Equal(t, time.Unix(1583971200, 0), r[1].Price.Timestamp)
			return true
		})

		repo.On("InsertMarket", mkt).Return(nil)
		repo.On("InsertRunners", run).Return(nil)

		err := handler.Handle(mk)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		repo.AssertExpectations(t)
	})

	t.Run("returns error if returned by repository", func(t *testing.T) {
		t.Helper()

		repo := new(market.MockRepository)
		handler := market.NewHandler(repo)

		mk := &queue.Market{
			ID:            "1.2818721",
			EventID:       148192,
			CompetitionID: 8,
			SeasonID:      17420,
			EventDate:     "2020-11-28T12:00:00+00:00",
			Name:          "OVER_UNDER_25",
			Side:          "BACK",
			Exchange:      "betfair",
			Runners: []*queue.Runner{
				{
					ID:   472671,
					Name: "Over 2.5 Goals",
					Prices: []queue.PriceSize{
						{
							Price: 1.951827161651,
							Size:  156.91671761,
						},
					},
				},
				{
					ID:   472672,
					Name: "Under 2.5 Goals",
					Prices: []queue.PriceSize{
						{
							Price: 2.0555411311,
							Size:  1.92141241,
						},
					},
				},
			},
			Timestamp: 1583971200,
		}

		mkt := mock.MatchedBy(func(m *market.Market) bool {
			date, _ := time.Parse(time.RFC3339, "2020-11-28T12:00:00+00:00")

			assert.Equal(t, "1.2818721", m.ID)
			assert.Equal(t, uint64(148192), m.EventID)
			assert.Equal(t, uint64(8), m.CompetitionID)
			assert.Equal(t, uint64(17420), m.SeasonID)
			assert.Equal(t, "OVER_UNDER_25", m.Name)
			assert.Equal(t, "BACK", m.Side)
			assert.Equal(t, "betfair", m.Exchange)
			assert.Equal(t, date, m.EventDate)
			return true
		})

		repo.On("InsertMarket", mkt).Return(errors.New("oh no"))
		repo.AssertNotCalled(t, "InsertRunners")

		err := handler.Handle(mk)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "oh no", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("returns error if event date provided cannot be parsed", func(t *testing.T) {
		t.Helper()

		repo := new(market.MockRepository)
		handler := market.NewHandler(repo)

		mk := &queue.Market{
			ID:            "1.2818721",
			EventID:       148192,
			CompetitionID: 8,
			SeasonID:      17420,
			EventDate:     "Wrong",
			Name:          "OVER_UNDER_25",
			Side:          "BACK",
			Exchange:      "betfair",
			Runners: []*queue.Runner{
				{
					ID:   472671,
					Name: "Over 2.5 Goals",
					Prices: []queue.PriceSize{
						{
							Price: 1.951827161651,
							Size:  156.91671761,
						},
					},
				},
				{
					ID:   472672,
					Name: "Under 2.5 Goals",
					Prices: []queue.PriceSize{
						{
							Price: 2.0555411311,
							Size:  1.92141241,
						},
					},
				},
			},
			Timestamp: 1583971200,
		}

		repo.AssertNotCalled(t, "Persist")

		err := handler.Handle(mk)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(
			t,
			"parsing time \"Wrong\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"Wrong\" as \"2006\"",
			err.Error(),
		)

		repo.AssertExpectations(t)
	})

	t.Run("parse runners name by sort priority for MATCH_ODDS market", func(t *testing.T) {
		t.Helper()

		repo := new(market.MockRepository)
		handler := market.NewHandler(repo)

		mk := &queue.Market{
			ID:            "1.2818721",
			EventID:       148192,
			CompetitionID: 8,
			SeasonID:      17420,
			EventDate:     "2020-11-28T12:00:00+00:00",
			Name:          "MATCH_ODDS",
			Side:          "BACK",
			Exchange:      "betfair",
			Runners: []*queue.Runner{
				{
					ID:   472671,
					Name: "West Ham",
					Sort: 1,
					Prices: []queue.PriceSize{
						{
							Price: 1.95581981,
							Size:  156.9198171,
						},
					},
				},
				{
					ID:   472672,
					Name: "Arsenal",
					Sort: 2,
					Prices: []queue.PriceSize{
						{
							Price: 2.05091981,
							Size:  1.92719817,
						},
					},
				},
				{
					ID:   472673,
					Name: "The Draw",
					Sort: 3,
					Prices: []queue.PriceSize{
						{
							Price: 3.051111,
							Size:  0.989819,
						},
					},
				},
			},
			Timestamp: 1583971200,
		}

		mkt := mock.MatchedBy(func(m *market.Market) bool {
			date, _ := time.Parse(time.RFC3339, "2020-11-28T12:00:00+00:00")

			assert.Equal(t, "1.2818721", m.ID)
			assert.Equal(t, uint64(148192), m.EventID)
			assert.Equal(t, uint64(8), m.CompetitionID)
			assert.Equal(t, uint64(17420), m.SeasonID)
			assert.Equal(t, "MATCH_ODDS", m.Name)
			assert.Equal(t, "BACK", m.Side)
			assert.Equal(t, "betfair", m.Exchange)
			assert.Equal(t, date, m.EventDate)
			return true
		})

		run := mock.MatchedBy(func(r []*market.Runner) bool {
			assert.Equal(t, uint64(472671), r[0].ID)
			assert.Equal(t, "Home", r[0].Name)
			assert.Equal(t, float32(1.96), r[0].Price.Value)
			assert.Equal(t, float32(156.92), r[0].Price.Size)
			assert.Equal(t, time.Unix(1583971200, 0), r[0].Price.Timestamp)
			assert.Equal(t, uint64(472672), r[1].ID)
			assert.Equal(t, "Away", r[1].Name)
			assert.Equal(t, float32(2.05), r[1].Price.Value)
			assert.Equal(t, float32(1.93), r[1].Price.Size)
			assert.Equal(t, time.Unix(1583971200, 0), r[1].Price.Timestamp)
			assert.Equal(t, "Draw", r[2].Name)
			assert.Equal(t, float32(3.05), r[2].Price.Value)
			assert.Equal(t, float32(0.99), r[2].Price.Size)
			assert.Equal(t, time.Unix(1583971200, 0), r[2].Price.Timestamp)
			return true
		})

		repo.On("InsertMarket", mkt).Return(nil)
		repo.On("InsertRunners", run).Return(nil)

		err := handler.Handle(mk)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		repo.AssertExpectations(t)
	})
}
