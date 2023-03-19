package warehouse

import "time"

type EventMarket struct {
	ID            string          `json:"id"`
	EventID       uint64          `json:"eventId"`
	Market        string          `json:"market"`
	Exchange      string          `json:"exchange"`
	Date          time.Time       `json:"date"`
	CompetitionID uint64          `json:"competitionId"`
	SeasonID      uint64          `json:"seasonId"`
	Runners       []*MarketRunner `json:"runners"`
}

type MarketReader interface {
	MarketRunnerOdds(eventID uint64, market, runner, exchange string, limit uint32) ([]*Odds, error)
}

type MarketRunner struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Odds []*Odds `json:"odds"`
}

type Odds struct {
	Price     float32   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}
