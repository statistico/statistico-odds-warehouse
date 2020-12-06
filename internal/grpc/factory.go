package grpc

import (
	"github.com/statistico/statistico-odds-warehouse/internal/market"
	"github.com/statistico/statistico-proto/statistico-odds-warehouse/go"
	"time"
)

func createMarketSelection(m *market.MarketRunner) *statisticoproto.MarketRunner {
	return &statisticoproto.MarketRunner{
		MarketId:             m.MarketID,
		MarketName:           m.Market.Name,
		RunnerName:           m.Runner.Name,
		RunnerPrice:          m.Price.Value,
		EventId:              m.EventID,
		CompetitionId:        m.CompetitionID,
		SeasonId:             m.SeasonID,
		EventDate:            m.EventDate.Format(time.RFC3339),
		Side:                 m.Side,
		Exchange:             m.Exchange,
	}
}
