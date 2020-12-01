package mock

import (
	"github.com/statistico/statistico-odds-warehouse/internal/grpc/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type MarketSelectionServer struct {
	mock.Mock
	grpc.ServerStream
}

func (m *MarketSelectionServer) Send(mk *proto.MarketSelection) error {
	args := m.Called(mk)
	return args.Error(0)
}
