package grpc

import (
	"github.com/statistico/statistico-odds-warehouse/internal/grpc/proto"
	"github.com/statistico/statistico-odds-warehouse/internal/market"
	"time"
)

func createMarketSelection(m *market.MarketRunner) *proto.MarketSelection {
	mk := proto.Market{
		MarketId:      m.Market.ID,
		MarketName:    m.Market.Name,
		EventId:       m.EventID,
		CompetitionId: m.CompetitionID,
		SeasonId:      m.SeasonID,
		EventDate:     m.EventDate.Format(time.RFC3339),
		Side:          m.Side,
		Exchange:      m.Exchange,
	}

	rn := proto.Runner{
		Id:    m.Runner.ID,
		Name:  m.Runner.Name,
		Price: m.Price.Value,
		Size:  m.Price.Size,
	}

	return &proto.MarketSelection{
		Market: &mk,
		Runner: &rn,
	}
}
