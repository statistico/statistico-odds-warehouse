package bootstrap

import (
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse/queue"
)

func (c Container) QueueMarketHandler() *queue.MarketHandler {
	return queue.NewMarketHandler(c.PostgresMarketWriter())
}
