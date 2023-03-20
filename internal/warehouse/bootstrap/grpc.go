package bootstrap

import "github.com/statistico/statistico-odds-warehouse/internal/warehouse/grpc"

func (c Container) MarketService() *grpc.MarketService {
	return grpc.NewMarketService(c.PostgresMarketReader(), c.Logger)
}
