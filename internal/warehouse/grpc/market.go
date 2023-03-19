package grpc

import (
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse"
	statistico "github.com/statistico/statistico-proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MarketService struct {
	reader warehouse.MarketReader
	logger *logrus.Logger
	statistico.UnimplementedOddsWarehouseServiceServer
}

func (m *MarketService) GetExchangeOdds(r *statistico.ExchangeOddsRequest, stream statistico.OddsWarehouseService_GetExchangeOddsClient) error {
	odds, err := m.reader.MarketRunnerOdds(r.EventId, r.Market, r.Runner, r.Exchange, r.Limit)

	if err != nil {
		m.logger.Errorf("Error fetching exchange odds in market service. %s", err.Error())
		return status.Error(codes.Internal, "Internal server error")
	}

	for _, o := range odds {
		x := &statistico.ExchangeOdds{
			Price:     o.Price,
			Timestamp: uint64(o.Timestamp.Unix()),
		}

		if err := stream.SendMsg(x); err != nil {
			m.logger.Errorf("Error streaming exchange odds back to client. %s", err.Error())
			continue
		}
	}

	return nil
}

func NewMarketService(r warehouse.MarketReader, l *logrus.Logger) *MarketService {
	return &MarketService{
		reader: r,
		logger: l,
	}
}
