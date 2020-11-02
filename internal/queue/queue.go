package queue

import "github.com/statistico/statistico-odds-warehouse/internal/market"

type Queue interface {
	ReceiveMarkets() <-chan *market.Market
}
