package bootstrap

import (
	"github.com/statistico/statistico-odds-warehouse/internal/app"
	"github.com/statistico/statistico-odds-warehouse/internal/app/postgres"
)

func (c Container) PostgresMarketReader() app.MarketReader {
	return postgres.NewMarketReader(c.DatabaseConnection)
}

func (c Container) PostgresMarketWriter() app.MarketWriter {
	return postgres.NewMarketWriter(c.DatabaseConnection)
}
