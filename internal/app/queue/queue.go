package queue

type Queue interface {
	ReceiveMarkets() <-chan *Market
}

type Market struct {
	ID            string    `json:"id"`
	EventID       uint64    `json:"eventId"`
	Name          string    `json:"name"`
	CompetitionID uint64    `json:"competitionId"`
	SeasonID      uint64    `json:"seasonId"`
	EventDate     string    `json:"date"`
	Side          string    `json:"side"`
	Exchange      string    `json:"exchange"`
	Runners       []*Runner `json:"runners"`
	Timestamp     int64     `json:"timestamp"`
}

type Runner struct {
	ID     uint64      `json:"id"`
	Name   string      `json:"name"`
	Sort   int8        `json:"sort"`
	Prices []PriceSize `json:"prices"`
}

type PriceSize struct {
	Price float32 `json:"price"`
	Size  float32 `json:"size"`
}
