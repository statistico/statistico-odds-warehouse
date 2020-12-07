package bootstrap

import "github.com/statistico/statistico-odds-warehouse/internal/app/grpc"

func (c Container) MarketService() *grpc.MarketService {
	return grpc.NewMarketService(c.MarketRepository(), c.Logger)
}
