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
	Runners       []*Runner `json:"runners"`
}

type Runner struct {
	ID        uint64  `json:"id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	Size      float32 `json:"size"`
	Timestamp int64   `json:"timestamp"`
}
