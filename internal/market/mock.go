package market

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) InsertMarket(o *Market) error {
	args := m.Called(o)
	return args.Error(0)
}

func (m *MockRepository) GetByRunner(q *RunnerQuery) ([]*Market, error) {
	args := m.Called(q)
	return args.Get(0).([]*Market), args.Error(1)
}
