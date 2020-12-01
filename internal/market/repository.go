package market

type RepositoryQuery struct {
	EventID    *uint64
	MarketName *string
	Side       *string
	SortBy     *string
}

type Repository interface {
	Persist(m *Market) error
}
