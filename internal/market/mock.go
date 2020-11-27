package market

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) InsertBTTSMarket(o *BTTSMarket) error {
	args := m.Called(o)
	return args.Error(0)
}

func (m *MockRepository) InsertMatchOddsMarket(o *MatchOddsMarket) error {
	args := m.Called(o)
	return args.Error(0)
}

func (m *MockRepository) InsertOverUnderMarket(o *OverUnderMarket) error {
	args := m.Called(o)
	return args.Error(0)
}
