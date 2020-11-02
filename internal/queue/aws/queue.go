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

type Message struct {
	Type   string   `json:"type"`
	MessageID string `json:"messageId"`
	TopicArn string `json:"topicArn"`
	Message string `json:"message"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
	SigningCertURL string `json:"signingCertUrl"`
	UnsubscribeURL string `json:"unsubscribeUrl"`
}

type Queue struct {
	client   sqsiface.SQSAPI
	logger   *logrus.Logger
	queueUrl string
	timeout  int64
}

func (q *Queue) ReceiveMarkets() <-chan *market.Market {
	ch := make(chan *market.Market, 100)

	go q.receiveMessages(ch)

	return ch
}

func (q *Queue) receiveMessages(ch chan<- *market.Market) {
	defer close(ch)

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
		return
	}

	for _, message := range result.Messages {
		q.parseMessage(message, ch)
	}
}

func (q *Queue) parseMessage(ms *sqs.Message, ch chan<- *market.Market) {
	var message Message
	err := json.Unmarshal([]byte(*ms.Body), &message)

	if err != nil {
		q.logger.Errorf("Unable to marshal message into message struct, %v.", err)
		return
	}

	var mk *market.Market
	err = json.Unmarshal([]byte(message.Message), &mk)

	if err != nil {
		q.logger.Errorf("Unable to marshal message into market struct, %v.", err)
		return
	}

	ch <- mk

	go q.deleteMessage(ms.ReceiptHandle)
}

func (q *Queue) deleteMessage(handle *string) {
	_, err := q.client.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(q.queueUrl),
		ReceiptHandle: handle,
	})

	if err != nil {
		q.logger.Errorf("Error deleting message from queue %q", err)
	}
}

func NewQueue(c sqsiface.SQSAPI, l *logrus.Logger, queue string, timeout int64) queue.Queue {
	return &Queue{
		client:   c,
		logger:   l,
		queueUrl: queue,
		timeout:  timeout,
	}
}
