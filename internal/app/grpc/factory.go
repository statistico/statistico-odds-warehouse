package grpc

import (
	"fmt"
	"github.com/statistico/statistico-odds-warehouse/internal/app/market"
	"github.com/statistico/statistico-proto/go"
	"time"
)

func createMarketRunner(m *market.MarketRunner) (*statisticoproto.MarketRunner, error) {
	mk := statisticoproto.MarketRunner{
		MarketId:      m.MarketID,
		MarketName:    m.MarketName,
		RunnerName:    m.RunnerName,
		EventId:       m.EventID,
		CompetitionId: m.CompetitionID,
		SeasonId:      m.SeasonID,
		EventDate:     m.EventDate.Format(time.RFC3339),
		Side:          m.Side,
		Exchange:      m.Exchange,
	}

	if len(m.Prices) == 0 {
		return nil, fmt.Errorf("market %s and runner %d does not contain prices", m.MarketID, m.RunnerID)
	}

	price := statisticoproto.Price{
		Value:     m.Prices[0].Value,
		Size:      m.Prices[0].Size,
		Timestamp: m.Prices[0].Timestamp.Unix(),
	}

	mk.Prices = append(mk.Prices, &price)

	return &mk, nil
}
