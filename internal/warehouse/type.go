package warehouse

import "time"

type MarketWriter interface {
	InsertMarket(market *Market) error
	InsertRunners(runners []*Runner) error
}

type MarketReader interface {
	ExchangeMarketRunnerOdds(eventID uint64, market, runner, exchange string, limit uint32) ([]*Odds, error)
	MarketsByEventID(eventID uint64, q *MarketReaderQuery) ([]*Market, error)
}

type MarketReaderQuery struct {
	Market   []string
	Exchange []string
}

type Market struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	EventID       uint64    `json:"eventId"`
	CompetitionID uint64    `json:"competitionId"`
	SeasonID      uint64    `json:"seasonId"`
	EventDate     time.Time `json:"date"`
	Exchange      string    `json:"exchange"`
	Runners       []*Runner `json:"runners"`
}

type Runner struct {
	ID        uint64 `json:"id"`
	MarketID  string `json:"marketId"`
	Name      string `json:"name"`
	BackPrice *Odds  `json:"backPrice"`
	LayPrice  *Odds  `json:"layPrice"`
}

type Odds struct {
	Value     float32   `json:"price"`
	Size      float32   `json:"size"`
	Side      string    `json:"side"`
	Timestamp time.Time `json:"timestamp"`
}
