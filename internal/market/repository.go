package market

type RepositoryQuery struct {
	EventID    *uint64
	MarketName *string
	Side       *string
	SortBy     *string
}

type Repository interface {
	InsertBTTSMarket(m *BTTSMarket) error
	InsertMatchOddsMarket(m *MatchOddsMarket) error
	InsertOverUnderMarket(m *OverUnderMarket) error
}
