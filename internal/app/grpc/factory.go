package grpc

import (
	"github.com/statistico/statistico-odds-warehouse/internal/app/market"
	"github.com/statistico/statistico-proto/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func createMarketRunner(m *market.MarketRunner) (*statistico.MarketRunner, error) {
	mk := statistico.MarketRunner{
		MarketId:      m.MarketID,
		MarketName:    m.MarketName,
		RunnerId:      m.RunnerID,
		RunnerName:    m.RunnerName,
		EventId:       m.EventID,
		CompetitionId: m.CompetitionID,
		SeasonId:      m.SeasonID,
		EventDate:     timestamppb.New(m.EventDate),
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
