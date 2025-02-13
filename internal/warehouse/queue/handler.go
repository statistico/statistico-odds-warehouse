package queue

import (
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse"
	"math"
	"time"
)

type MarketHandler struct {
	writer warehouse.MarketWriter
}

func (m *MarketHandler) Handle(q *EventMarket) error {
	var runners []*warehouse.Runner

	for _, r := range q.Runners {
		if len(r.BackPrices) == 0 && len(r.LayPrices) == 0 {
			continue
		}

		run := warehouse.Runner{
			ID:       r.ID,
			MarketID: q.ID,
			Name:     r.Name,
			Label:    r.Label,
		}

		if len(r.BackPrices) != 0 {
			price := warehouse.Odds{
				Value:     float32(math.Round(float64(r.BackPrices[0].Price*100)) / 100),
				Size:      float32(math.Round(float64(r.BackPrices[0].Size*100)) / 100),
				Timestamp: time.Unix(q.Timestamp, 0),
			}

			run.BackPrice = &price
		}

		if len(r.LayPrices) != 0 {
			price := warehouse.Odds{
				Value:     float32(math.Round(float64(r.LayPrices[0].Price*100)) / 100),
				Size:      float32(math.Round(float64(r.LayPrices[0].Size*100)) / 100),
				Timestamp: time.Unix(q.Timestamp, 0),
			}

			run.LayPrice = &price
		}

		runners = append(runners, &run)
	}

	market := warehouse.Market{
		ID:            q.ID,
		Name:          q.Name,
		EventID:       q.EventID,
		CompetitionID: q.CompetitionID,
		SeasonID:      q.SeasonID,
		EventDate:     time.Unix(q.EventDate, 0),
		Exchange:      q.Exchange,
	}

	if err := m.writer.InsertMarket(&market); err != nil {
		return err
	}

	if err := m.writer.InsertRunners(runners); err != nil {
		return err
	}

	return nil
}

func NewMarketHandler(w warehouse.MarketWriter) *MarketHandler {
	return &MarketHandler{writer: w}
}
