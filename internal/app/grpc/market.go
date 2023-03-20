package grpc

import (
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-warehouse/internal/app"
	"github.com/statistico/statistico-proto/go"
)

type MarketService struct {
	repository app.MarketWriter
	logger     *logrus.Logger
	statistico.UnimplementedOddsWarehouseServiceServer
}

func (m *MarketService) MarketRunnerSearch(r *statistico.MarketRunnerRequest, stream statistico.OddsWarehouseService_MarketRunnerSearchServer) error {
	return nil
}

func NewMarketService(r app.MarketWriter, l *logrus.Logger) *MarketService {
	return &MarketService{
		repository: r,
		logger:     l,
	}
}
