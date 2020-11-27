package market

type RepositoryQuery struct {
	EventID    *uint64
	MarketName *string
	Side       *string
	SortBy     *string
}

type Repository interface {
	InsertOverUnderMarket(m *OverUnderMarket) error
}
