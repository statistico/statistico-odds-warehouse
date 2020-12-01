package market

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Persist(o *Market) error {
	args := m.Called(o)
	return args.Error(0)
}
