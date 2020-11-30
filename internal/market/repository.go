package market

import "time"

type Repository interface {
	InsertMarket(m *Market) error
	GetByRunner(q *RunnerQuery) ([]*Market, error)
}

type RepositoryQuery struct {
	MarketName   *string
	DateFrom     *time.Time
	DateTo       *time.Time
	CompetitionIDs []uint64
	SeasonIDs    []uint64
}

type RunnerQuery struct {
	Name    string
	Line  string
	GreaterThan *float32
	LessThan    *float32
	DateFrom     *time.Time
	DateTo       *time.Time
}
