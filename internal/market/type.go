package market

import "time"

type Market struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	EventID       uint64    `json:"eventId"`
	CompetitionID uint64    `json:"competitionId"`
	SeasonID      uint64    `json:"seasonId"`
	EventDate     time.Time `json:"date"`
	Side          string    `json:"side"`
	Exchange      string    `json:"exchange"`
}

type Runner struct {
	ID        uint64     `json:"id"`
	MarketID  string     `json:"marketId"`
	Name      string     `json:"name"`
	Price     Price      `json:"price"`
}

type Price struct {
	Value     float32      `json:"price"`
	Size      float32     `json:"size"`
	Timestamp time.Time   `json:"timestamp"`
}
