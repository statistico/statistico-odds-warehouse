package grpc

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/statistico/statistico-odds-warehouse/internal/app/market"
	"github.com/statistico/statistico-proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func convertMarketSelectionRequest(r *statistico.MarketRunnerRequest) (*market.RunnerQuery, error) {
	q := market.RunnerQuery{
		MarketName:     r.GetMarket(),
		RunnerName:     r.GetRunner(),
		Line:           r.GetLine(),
		Side:           r.GetSide().String(),
		CompetitionIDs: r.GetCompetitionIds(),
		SeasonIDs:      r.GetSeasonIds(),
	}

	if r.GetMinOdds() != nil {
		q.GreaterThan = &r.GetMinOdds().Value
	}

	if r.GetMaxOdds() != nil {
		q.LessThan = &r.GetMaxOdds().Value
	}

	if r.GetDateFrom() != nil {
		date, err := ptypes.Timestamp(r.GetDateFrom())

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Date provided is not a valid RFC3339 date")
		}

		q.DateFrom = &date
	}

	if r.GetDateTo() != nil {
		date, err := ptypes.Timestamp(r.GetDateTo())

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Date provided is not a valid RFC3339 date")
		}

		q.DateTo = &date
	}

	return &q, nil
}
