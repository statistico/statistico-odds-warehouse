package bootstrap

import (
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse/postgres"
)

func (c Container) PostgresMarketReader() warehouse.MarketReader {
	return postgres.NewMarketReader(c.DatabaseConnection)
}

func (c Container) PostgresMarketWriter() warehouse.MarketWriter {
	return postgres.NewMarketWriter(c.DatabaseConnection)
}
