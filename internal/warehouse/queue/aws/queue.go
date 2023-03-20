package aws

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse/queue"
)

type Message struct {
	Type           string `json:"type"`
	MessageID      string `json:"messageId"`
	TopicArn       string `json:"topicArn"`
	Message        string `json:"message"`
	Signature      string `json:"signature"`
	SigningCertURL string `json:"signingCertUrl"`
	UnsubscribeURL string `json:"unsubscribeUrl"`
}

type Queue struct {
	client   sqsiface.SQSAPI
	logger   *logrus.Logger
	queueUrl string
	timeout  int64
}

func (q *Queue) ReceiveMarkets() []*queue.EventMarket {
	markets := []*queue.EventMarket{}

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
		return markets
	}

	for _, message := range result.Messages {
		mk := q.parseMessage(message)

		if mk == nil {
			continue
		}

		markets = append(markets, mk)
	}

	return markets
}

func (q *Queue) parseMessage(ms *sqs.Message) *queue.EventMarket {
	var mk *queue.EventMarket
	err := json.Unmarshal([]byte(*ms.Body), &mk)

	if err != nil {
		q.logger.Errorf("Unable to marshal message into market struct, %v.", err)
		return nil
	}

	q.deleteMessage(ms.ReceiptHandle)

	return mk
}

func (q *Queue) deleteMessage(handle *string) {
	input := &sqs.DeleteMessageInput{
		QueueUrl:      &q.queueUrl,
		ReceiptHandle: handle,
	}

	_, err := q.client.DeleteMessage(input)

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
