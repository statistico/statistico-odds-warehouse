package grpc

import (
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
		Exchange:      m.Exchange,
		Price:         &statistico.Price{
			Value:     m.Price.Value,
			Size:      m.Price.Size,
			Side:      statistico.SideEnum(statistico.SideEnum_value[m.Price.Side]),
			Timestamp: m.Price.Timestamp.Unix(),
		},
	}

	return &mk, nil
}
