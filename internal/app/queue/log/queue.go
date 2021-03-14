package log

import (
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-warehouse/internal/app/queue"
	"time"
)

type Queue struct {
	logger *logrus.Logger
}

func (q *Queue) ReceiveMarkets() <-chan *queue.EventMarket {
	ch := make(chan *queue.EventMarket, 100)

	q.logger.Infof("Pretending to poll for messages from queue...")

	go q.simulate(ch)

	return ch
}

func (q *Queue) simulate(ch chan<- *queue.EventMarket) {
	time.Sleep(10 * time.Second)

	q.logger.Infof("..polling complete.")

	close(ch)
}

func NewQueue(l *logrus.Logger) queue.Queue {
	return &Queue{logger: l}
}
