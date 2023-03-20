package market

import (
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse"
	"github.com/stretchr/testify/mock"
)

type MockMarketWriter struct {
	mock.Mock
}

func (m *MockMarketWriter) InsertMarket(o *warehouse.Market) error {
	args := m.Called(o)
	return args.Error(0)
}

func (m *MockMarketWriter) InsertRunners(r []*warehouse.Runner) error {
	args := m.Called(r)
	return args.Error(0)
}
