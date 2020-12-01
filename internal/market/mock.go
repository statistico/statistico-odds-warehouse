package market

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) InsertMarket(o *Market) error {
	args := m.Called(o)
	return args.Error(0)
}

func (m *MockRepository) InsertRunners(r []*Runner) error {
	args := m.Called(r)
	return args.Error(0)
}
