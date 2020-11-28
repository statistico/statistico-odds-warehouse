package market

import (
	"github.com/statistico/statistico-odds-warehouse/internal/queue"
	"time"
)

type Handler struct {
	repository Repository
}

func (m *Handler) Handle(q *queue.Market) error {
	var runners []*Runner

	for _, r := range q.Runners {
		run := Runner{
			ID:    r.ID,
			Name:  r.Name,
			Price: r.Prices[0].Price,
			Size:  r.Prices[0].Size,
		}

		runners = append(runners, &run)
	}

	date, err := time.Parse(time.RFC3339, q.EventDate)

	if err != nil {
		return err
	}

	market := Market{
		ID:            q.ID,
		Name:          q.Name,
		EventID:       q.EventID,
		CompetitionID: q.CompetitionID,
		SeasonID:      q.SeasonID,
		EventDate:     date,
		Side:          q.Side,
		Exchange:      q.Exchange,
		Runners:       runners,
		Timestamp:     q.Timestamp,
	}

	return m.repository.InsertMarket(&market)
}

func NewHandler(r Repository) *Handler {
	return &Handler{repository: r}
}
