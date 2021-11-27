package aws_test

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-odds-warehouse/internal/app/queue"
	saws "github.com/statistico/statistico-odds-warehouse/internal/app/queue/aws"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestQueue_ReceiveMarkets(t *testing.T) {
	t.Run("calls client and pushes messages into channel provided", func(t *testing.T) {
		t.Helper()

		client := new(saws.MockSqsClient)
		logger, _ := test.NewNullLogger()

		r := saws.NewQueue(client, logger, "messages", 3600)

		input := mock.MatchedBy(func(i *sqs.ReceiveMessageInput) bool {
			assert.Equal(t, "messages", *i.QueueUrl)
			assert.Equal(t, int64(3600), *i.WaitTimeSeconds)
			return true
		})

		messages := []*sqs.Message{
			{
				ReceiptHandle: aws.String("1234"),
				Body:          &messageBody,
			},
		}

		deleteInput := mock.MatchedBy(func(i *sqs.DeleteMessageInput) bool {
			assert.Equal(t, "1234", *i.ReceiptHandle)
			return true
		})

		client.On("ReceiveMessage", input).Return(&sqs.ReceiveMessageOutput{Messages: messages}, nil)
		client.On("DeleteMessage", deleteInput).Return(&sqs.DeleteMessageOutput{}, nil)

		mk := &queue.EventMarket{
			ID:       "1.2818721",
			EventID:  148192,
			Name:     "OVER_UNDER_25",
			Exchange: "betfair",
			Runners: []*queue.Runner{
				{
					ID:   472671,
					Name: "Over 2.5 Goals",
					BackPrices: []queue.PriceSize{
						{
							Price: 1.95,
							Size:  1461,
						},
					},
					LayPrices: []queue.PriceSize{
						{
							Price: 1.95,
							Size:  1461,
						},
					},
				},
			},
			Timestamp: 1583971200,
		}

		ch := r.ReceiveMarkets()

		mt := <-ch

		assert.Equal(t, mk, mt)
		client.AssertExpectations(t)
	})

	t.Run("logs and returns error if error returned by SQS client", func(t *testing.T) {
		t.Helper()

		client := new(saws.MockSqsClient)
		logger, hook := test.NewNullLogger()

		r := saws.NewQueue(client, logger, "messages", 3600)

		input := mock.MatchedBy(func(i *sqs.ReceiveMessageInput) bool {
			assert.Equal(t, "messages", *i.QueueUrl)
			assert.Equal(t, int64(3600), *i.WaitTimeSeconds)
			return true
		})

		e := errors.New("error happened")

		client.On("ReceiveMessage", input).Return(&sqs.ReceiveMessageOutput{}, e)

		ch := r.ReceiveMarkets()

		<-ch

		assert.Equal(t, 0, len(ch))
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Unable to receive messages from queue \"messages\", error happened.", hook.LastEntry().Message)

		client.AssertExpectations(t)
	})

	t.Run("logs error if unable to parse message body in market struct", func(t *testing.T) {
		t.Helper()

		client := new(saws.MockSqsClient)
		logger, hook := test.NewNullLogger()

		r := saws.NewQueue(client, logger, "messages", 3600)

		input := mock.MatchedBy(func(i *sqs.ReceiveMessageInput) bool {
			assert.Equal(t, "messages", *i.QueueUrl)
			assert.Equal(t, int64(3600), *i.WaitTimeSeconds)
			return true
		})

		body := "invalid body"

		messages := []*sqs.Message{
			{
				Body: &body,
			},
		}

		client.On("ReceiveMessage", input).Return(&sqs.ReceiveMessageOutput{Messages: messages}, nil)

		ch := r.ReceiveMarkets()

		<-ch

		assert.Equal(t, 0, len(ch))
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Unable to marshal message into market struct, invalid character 'i' looking for beginning of value.", hook.LastEntry().Message)

		client.AssertExpectations(t)
	})
}

var messageBody = `
	{
	  "id": "1.2818721",
	  "eventId": 148192,
	  "name": "OVER_UNDER_25",
	  "exchange": "betfair",
	  "runners": [
		{
		  "id": 472671,
		  "name": "Over 2.5 Goals",
		  "backPrices": [
			{
			  "price": 1.95,
			  "size": 1461
			}
		  ],
		  "layPrices": [
			{
			  "price": 1.95,
			  "size": 1461
			}
		  ]
		}
	  ],
	  "timestamp": 1583971200
	}
`
