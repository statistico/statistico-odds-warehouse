package grpc

import (
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-warehouse/internal/grpc/proto"
	"github.com/statistico/statistico-odds-warehouse/internal/market"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MarketService struct {
	repository market.Repository
	logger *logrus.Logger
}

func (m *MarketService) MarketSelectionSearch(r *proto.MarketSelectionRequest, stream proto.MarketService_MarketSelectionSearchServer) error {
	query, err := convertMarketSelectionRequest(r)

	if err != nil {
		return err
	}

	markets, err := m.repository.MarketRunners(query)

	if err != nil {
		m.logger.Errorf("Error fetching market runners in market service. %s", err.Error())
		return status.Error(codes.Internal, "Internal server error")
	}

	for _, mk := range markets {
		if err := stream.Send(createMarketSelection(mk)); err != nil {
			m.logger.Errorf("Error streaming market runner back to client. %s", err.Error())
			continue
		}
	}

	return nil
}

func NewMarketService(r market.Repository, l *logrus.Logger) *MarketService {
	return &MarketService{
		repository: r,
		logger:     l,
	}
}