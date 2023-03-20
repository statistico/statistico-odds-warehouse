package bootstrap

import (
	"github.com/statistico/statistico-odds-warehouse/internal/app/market"
)

func (c Container) MarketHandler() *market.Handler {
	return market.NewHandler(c.PostgresMarketWriter())
}
