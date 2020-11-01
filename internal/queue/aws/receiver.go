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

type MarketReceiver struct {
	client   sqsiface.SQSAPI
	logger   *logrus.Logger
	queueUrl string
	timeout  int64
}

func (m *MarketReceiver) Receive(ch chan<- *market.Market) error {
	input := &sqs.ReceiveMessageInput{
		QueueUrl: &m.queueUrl,
		MessageAttributeNames: aws.StringSlice([]string{
			"All",
		}),
		WaitTimeSeconds: &m.timeout,
	}

	result, err := m.client.ReceiveMessage(input)

	if err != nil {
		m.logger.Errorf("Unable to receive messages from queue %q, %v.", m.queueUrl, err)
		return err
	}

	for _, message := range result.Messages {
		m.parseMessage(message, ch)
	}

	return nil
}

func (m *MarketReceiver) parseMessage(ms *sqs.Message, ch chan<- *market.Market) {
	var mk *market.Market
	err := json.Unmarshal([]byte(*ms.Body), &mk)

	if err != nil {
		m.logger.Errorf("Unable to marshal message into market struct, %v.", err)
		return
	}

	ch <- mk
}

func NewMarketReceiver(c sqsiface.SQSAPI, l *logrus.Logger, queue string, timeout int64) queue.MarketReceiver {
	return &MarketReceiver{
		client:   c,
		logger:   l,
		queueUrl: queue,
		timeout:  timeout,
	}
}
