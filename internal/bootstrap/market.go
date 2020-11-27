package bootstrap

import (
	"github.com/statistico/statistico-odds-warehouse/internal/market"
	"github.com/statistico/statistico-odds-warehouse/internal/market/postgres"
)

func (c Container) MarketRepository() *postgres.MarketRepository {
	return postgres.NewMarketRepository(c.DatabaseConnection)
}

func (c Container) MarketHandler() *market.Handler {
	return market.NewHandler(c.MarketRepository())
}
