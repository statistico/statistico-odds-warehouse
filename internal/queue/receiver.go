package queue

import "github.com/statistico/statistico-odds-warehouse/internal/market"

type MarketReceiver interface {
	Receive(ch chan<- *market.Market) error
}
