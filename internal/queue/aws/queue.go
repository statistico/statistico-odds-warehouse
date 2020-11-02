package aws

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-warehouse/internal/market"
	"github.com/statistico/statistico-odds-warehouse/internal/queue"
)

type Queue struct {
	client   sqsiface.SQSAPI
	logger   *logrus.Logger
	queueUrl string
	timeout  int64
}

func (q *Queue) ReceiveMarkets(ch chan<- *market.Market) error {
	input := &sqs.ReceiveMessageInput{
		QueueUrl: &q.queueUrl,
		MessageAttributeNames: aws.StringSlice([]string{
			"All",
		}),
		WaitTimeSeconds: &q.timeout,
	}

	result, err := q.client.ReceiveMessage(input)

	if err != nil {
		q.logger.Errorf("Unable to receive messages from queue %q, %v.", q.queueUrl, err)
		return err
	}

	for _, message := range result.Messages {
		q.parseMessage(message, ch)
	}

	return nil
}

func (q *Queue) parseMessage(ms *sqs.Message, ch chan<- *market.Market) {
	var mk *market.Market
	err := json.Unmarshal([]byte(*ms.Body), &mk)

	if err != nil {
		q.logger.Errorf("Unable to marshal message into market struct, %v.", err)
		return
	}

	ch <- mk
}

func NewQueue(c sqsiface.SQSAPI, l *logrus.Logger, queue string, timeout int64) queue.Queue {
	return &Queue{
		client:   c,
		logger:   l,
		queueUrl: queue,
		timeout:  timeout,
	}
}
