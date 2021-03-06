package market

import "time"

type Repository interface {
	InsertMarket(market *Market) error
	InsertRunners(runners []*Runner) error
	MarketRunners(q *RunnerQuery) ([]*MarketRunner, error)
}

type RunnerQuery struct {
	MarketName     string
	RunnerName     string
	Line           string
	Side           string
	GreaterThan    *float32
	LessThan       *float32
	CompetitionIDs []uint64
	SeasonIDs      []uint64
	DateFrom       *time.Time
	DateTo         *time.Time
}
