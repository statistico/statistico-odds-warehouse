package log

import (
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-warehouse/internal/market"
	"github.com/statistico/statistico-odds-warehouse/internal/queue"
	"time"
)

type Queue struct {
	logger *logrus.Logger
}

func (q *Queue) ReceiveMarkets(ch chan<- *market.Market) error {
	q.logger.Infof("Pretending to poll for messages from queue...")

	time.Sleep(3 * time.Second)

	q.logger.Infof("..polling complete.")

	return nil
}

func NewQueue(l *logrus.Logger) queue.Queue {
	return &Queue{logger: l}
}
