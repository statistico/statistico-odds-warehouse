package log

import (
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse/queue"
	"time"
)

type Queue struct {
	logger *logrus.Logger
}

func (q *Queue) ReceiveMarkets() []*queue.EventMarket {
	markets := []*queue.EventMarket{}

	q.logger.Infof("Pretending to poll for messages from queue...")

	time.Sleep(10 * time.Second)

	q.logger.Infof("...polling complete.")

	return markets
}

func NewQueue(l *logrus.Logger) queue.Queue {
	return &Queue{logger: l}
}
