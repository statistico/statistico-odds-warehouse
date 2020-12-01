package market

import (
	"github.com/statistico/statistico-odds-warehouse/internal/queue"
	"math"
	"time"
)

const Away = "Away"
const Draw = "Draw"
const Home = "Home"
const MatchOdds = "MATCH_ODDS"

type Handler struct {
	repository Repository
}

func (m *Handler) Handle(q *queue.Market) error {
	var runners []*Runner

	for _, r := range q.Runners {
		run := Runner{
			ID:        r.ID,
			Name:      parseRunner(q.Name, r),
			Price:     float32(math.Round(float64(r.Prices[0].Price*100)) / 100),
			Size:      float32(math.Round(float64(r.Prices[0].Size*100)) / 100),
			Timestamp: q.Timestamp,
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
	}

	return m.repository.Persist(&market)
}

func parseRunner(market string, runner *queue.Runner) string {
	if market != MatchOdds {
		return runner.Name
	}

	if runner.Sort == 1 {
		return Home
	}

	if runner.Sort == 2 {
		return Away
	}

	return Draw
}

func NewHandler(r Repository) *Handler {
	return &Handler{repository: r}
}
