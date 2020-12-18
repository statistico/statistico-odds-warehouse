package mock

import (
	"github.com/statistico/statistico-proto/go"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type MarketSelectionServer struct {
	mock.Mock
	grpc.ServerStream
}

func (m *MarketSelectionServer) Send(mk *statisticoproto.MarketRunner) error {
	args := m.Called(mk)
	return args.Error(0)
}
