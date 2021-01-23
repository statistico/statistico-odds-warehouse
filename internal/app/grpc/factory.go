package grpc

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/statistico/statistico-odds-warehouse/internal/app/market"
	"github.com/statistico/statistico-proto/go"
)

func createMarketRunner(m *market.MarketRunner) (*statistico.MarketRunner, error) {
	date, err := ptypes.TimestampProto(m.EventDate)

	if err != nil {
		return nil, err
	}

	mk := statistico.MarketRunner{
		MarketId:      m.MarketID,
		MarketName:    m.MarketName,
		RunnerName:    m.RunnerName,
		EventId:       m.EventID,
		CompetitionId: m.CompetitionID,
		SeasonId:      m.SeasonID,
		EventDate:     date,
		Side:          m.Side,
		Exchange:      m.Exchange,
	}

	if len(m.Prices) == 0 {
		return nil, fmt.Errorf("market %s and runner %d does not contain prices", m.MarketID, m.RunnerID)
	}

	price := statistico.Price{
		Value:     m.Prices[0].Value,
		Size:      m.Prices[0].Size,
		Timestamp: m.Prices[0].Timestamp.Unix(),
	}

	mk.Prices = append(mk.Prices, &price)

	return &mk, nil
}
