package grpc

import (
	"github.com/statistico/statistico-odds-warehouse/internal/market"
	"github.com/statistico/statistico-proto/statistico-odds-warehouse/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func convertMarketSelectionRequest(r *statisticoproto.MarketRunnerRequest) (*market.RunnerQuery, error) {
	q := market.RunnerQuery{
		MarketName:     r.Name,
		RunnerName:     r.RunnerFilter.Name,
		Line:           r.RunnerFilter.Line.String(),
		CompetitionIDs: r.GetCompetitionIds(),
		SeasonIDs:      r.GetSeasonIds(),
	}

	if r.GetDateFrom() != nil {
		date, err := time.Parse(time.RFC3339, r.GetDateFrom().GetValue())

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Date provided is not a valid RFC3339 date")
		}

		q.DateFrom = &date
	}

	if r.GetDateTo() != nil {
		date, err := time.Parse(time.RFC3339, r.GetDateTo().GetValue())

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Date provided is not a valid RFC3339 date")
		}

		q.DateTo = &date
	}

	filters := r.GetRunnerFilter().GetOperators()

	for _, f := range filters {
		if f.GetOperator().String() == "GTE" {
			q.GreaterThan = &f.Value
		}

		if f.GetOperator().String() == "LTE" {
			q.LessThan = &f.Value
		}
	}

	return &q, nil
}
