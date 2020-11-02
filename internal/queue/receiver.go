package queue

import "github.com/statistico/statistico-odds-warehouse/internal/market"

type Queue interface {
	ReceiveMarkets(ch chan<- *market.Market) error
}
