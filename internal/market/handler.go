package market

import (
	"fmt"
	"github.com/statistico/statistico-odds-warehouse/internal/queue"
)

const BTTS = "BOTH_TEAMS_TO_SCORE"

type Handler struct {
	repository Repository
}

func (m *Handler) Handle(q *queue.Market) error {
	if isSupportedOverUnderMarket(q.Name) {
		return m.repository.InsertOverUnderMarket(createOverUnderMarket(q))
	}

	if q.Name == BTTS {
		return m.repository.InsertBTTSMarket(createBTTSMarket(q))
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

func createBTTSMarket(m *queue.Market) *BTTSMarket {
	var yes PriceSize
	var no PriceSize

	for _, r := range m.Runners {
		if r.Name == "Yes" {
			price := r.Prices[0]

			yes.Price = price.Price
			yes.Size = price.Size
		}

		if r.Name == "No" {
			price := r.Prices[0]

			no.Price = price.Price
			no.Size = price.Size
		}
	}

	return &BTTSMarket{
		ID:        m.ID,
		EventID:   m.EventID,
		Name:      m.Name,
		Side:      m.Side,
		Exchange:  m.Exchange,
		Yes:       yes,
		No:        no,
		Timestamp: m.Timestamp,
	}
}

func NewHandler(r Repository) *Handler {
	return &Handler{repository: r}
}
