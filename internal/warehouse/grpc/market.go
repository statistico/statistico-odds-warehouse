package grpc

import (
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse"
	"github.com/statistico/statistico-proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
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

func (m *MarketService) GetEventMarkets(r *statistico.EventMarketRequest, stream statistico.OddsWarehouseService_GetEventMarketsServer) error {
	q := warehouse.MarketReaderQuery{
		Market:   r.Market,
		Exchange: r.Exchange,
	}

	m.logger.Infof("Receiving a request for event %d: %+v\n", r.EventId, q.Market)

	markets, err := m.reader.MarketsByEventID(r.EventId, &q)

	if err != nil {
		m.logger.Errorf("error fetching markets from reader: %s", err.Error())
		return status.Error(codes.Internal, "internal server error")
	}

	for _, mk := range markets {
		ms := statistico.Market{
			Id:            mk.ID,
			Name:          mk.Name,
			EventId:       mk.EventID,
			CompetitionId: mk.CompetitionID,
			SeasonId:      mk.SeasonID,
			Exchange:      mk.Exchange,
			DateTime: &statistico.Date{
				Utc: mk.EventDate.Unix(),
				Rfc: mk.EventDate.Format(time.RFC3339),
			},
		}

		var runners []*statistico.Runner

		for _, r := range mk.Runners {
			runners = append(runners, &statistico.Runner{
				Id:   r.ID,
				Name: r.Name,
				BackOdds: &statistico.ExchangeOdds{
					Price:     r.BackPrice.Value,
					Size:      r.BackPrice.Size,
					Side:      r.BackPrice.Side,
					Timestamp: uint64(r.BackPrice.Timestamp.Unix()),
				},
			})
		}

		ms.Runners = runners

		if err := stream.Send(&ms); err != nil {
			m.logger.Errorf("error streaming market back to client: %s", err.Error())
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
