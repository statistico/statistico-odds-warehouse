package market

import (
	"fmt"
	"github.com/statistico/statistico-odds-warehouse/internal/queue"
)

type Handler struct {
	repository Repository
}

func (m *Handler) Handle(q *queue.Market) error {
	if isSupportedOverUnderMarket(q.Name) {
		return m.repository.InsertOverUnderMarket(createOverUnderMarket(q))
	}

	return fmt.Errorf("market %s is not supported", q.Name)
}

func createOverUnderMarket(m *queue.Market) *OverUnderMarket {
	var over PriceSize
	var under PriceSize

	for _, r := range m.Runners {
		if r.Name[0:4] == "Over" {
			price := r.Prices[0]

			over.Price = price.Price
			over.Size = price.Size
		}

		if r.Name[0:5] == "Under" {
			price := r.Prices[0]

			under.Price = price.Price
			under.Size = price.Size
		}
	}

	return &OverUnderMarket{
		ID:        m.ID,
		EventID:   m.EventID,
		Name:      m.Name,
		Side:      m.Side,
		Exchange:  m.Exchange,
		Over:      over,
		Under:     under,
		Timestamp: m.Timestamp,
	}
}

func NewHandler(r Repository) *Handler {
	return &Handler{repository: r}
}
