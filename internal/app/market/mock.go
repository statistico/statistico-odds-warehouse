package market

import (
	"github.com/statistico/statistico-odds-warehouse/internal/app"
	"github.com/stretchr/testify/mock"
)

type MockMarketWriter struct {
	mock.Mock
}

func (m *MockMarketWriter) InsertMarket(o *app.Market) error {
	args := m.Called(o)
	return args.Error(0)
}

func (m *MockMarketWriter) InsertRunners(r []*app.Runner) error {
	args := m.Called(r)
	return args.Error(0)
}
