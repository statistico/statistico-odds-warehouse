package grpc

import (
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-warehouse/internal/app/market"
	"github.com/statistico/statistico-proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MarketService struct {
	repository market.Repository
	logger     *logrus.Logger
}

func (m *MarketService) MarketRunnerSearch(r *statisticoproto.MarketRunnerRequest, stream statisticoproto.MarketService_MarketRunnerSearchServer) error {
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
		mr, err := createMarketRunner(mk)

		if err != nil {
			m.logger.Errorf("Error converting market runner in market service. %s", err.Error())
			continue
		}

		if err := stream.Send(mr); err != nil {
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
