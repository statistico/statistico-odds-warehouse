package market

type BTTSMarket struct {
	ID        string    `json:"id"`
	EventID   uint64    `json:"eventId"`
	Name      string    `json:"name"`
	Side      string    `json:"side"`
	Exchange  string    `json:"exchange"`
	Yes       PriceSize `json:"yes"`
	No        PriceSize `json:"no"`
	Timestamp int64     `json:"timestamp"`
}

type MatchOddsMarket struct {
	ID        string    `json:"id"`
	EventID   uint64    `json:"eventId"`
	Name      string    `json:"name"`
	Side      string    `json:"side"`
	Exchange  string    `json:"exchange"`
	Home      PriceSize `json:"home"`
	Away      PriceSize `json:"away"`
	Draw      PriceSize `json:"draw"`
	Timestamp int64     `json:"timestamp"`
}

type OverUnderMarket struct {
	ID        string    `json:"id"`
	EventID   uint64    `json:"eventId"`
	Name      string    `json:"name"`
	Side      string    `json:"side"`
	Exchange  string    `json:"exchange"`
	Over      PriceSize `json:"over"`
	Under     PriceSize `json:"under"`
	Timestamp int64     `json:"timestamp"`
}

type PriceSize struct {
	Price float32 `json:"price"`
	Size  float32 `json:"size"`
}
