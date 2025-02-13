package queue

type Queue interface {
	ReceiveMarkets() []*EventMarket
}

type EventMarket struct {
	ID            string    `json:"id"`
	EventID       uint64    `json:"eventId"`
	Name          string    `json:"name"`
	CompetitionID uint64    `json:"competitionId"`
	SeasonID      uint64    `json:"seasonId"`
	EventDate     int64     `json:"eventDate"`
	Exchange      string    `json:"exchange"`
	Runners       []*Runner `json:"runners"`
	Timestamp     int64     `json:"timestamp"`
}

type Runner struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Label      *string     `json:"label"`
	BackPrices []PriceSize `json:"backPrices"`
	LayPrices  []PriceSize `json:"layPrices"`
}

type PriceSize struct {
	Price float32 `json:"price"`
	Size  float32 `json:"size"`
}
