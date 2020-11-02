package bootstrap

import "github.com/statistico/statistico-odds-warehouse/internal/market/postgres"

func (c Container) MarketRepository() *postgres.MarketRepository {
	return postgres.NewMarketRepository(c.DatabaseConnection)
}
