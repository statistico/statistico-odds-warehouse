package grpc

import (
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse"
	"github.com/statistico/statistico-proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MarketService struct {
	reader warehouse.MarketReader
	logger *logrus.Logger
	statistico.UnimplementedOddsWarehouseServiceServer
}

func (m *MarketService) GetExchangeOdds(r *statistico.ExchangeOddsRequest, stream statistico.OddsWarehouseService_GetExchangeOddsServer) error {
	odds, err := m.reader.ExchangeMarketRunnerOdds(r.EventId, r.Market, r.Runner, r.Exchange, r.Limit)

	if err != nil {
		m.logger.Errorf("error fetching odds from reader: %s", err.Error())
		return status.Error(codes.Internal, "internal server error")
	}

	for _, o := range odds {
		eo := statistico.ExchangeOdds{
			Price:     o.Value,
			Timestamp: uint64(o.Timestamp.Unix()),
		}

		if err := stream.Send(&eo); err != nil {
			m.logger.Errorf("error streaming exchange odds back to client: %s", err.Error())
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
