package aws_test

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-odds-warehouse/internal/market"
	"github.com/statistico/statistico-odds-warehouse/internal/queue/aws"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestMarketReceiver_Receive(t *testing.T) {
	t.Run("calls client and pushes messages into channel provided", func(t *testing.T) {
		t.Helper()

		client := new(aws.MockSqsClient)
		logger, _ := test.NewNullLogger()

		r := aws.NewMarketReceiver(client, logger, "messages", 3600)

		input := mock.MatchedBy(func(i *sqs.ReceiveMessageInput) bool {
			assert.Equal(t, "messages", *i.QueueUrl)
			assert.Equal(t, int64(3600), *i.WaitTimeSeconds)
			return true
		})

		messages := []*sqs.Message{
			{
				Body: &messageBody,
			},
		}

		client.On("ReceiveMessage", input).Return(&sqs.ReceiveMessageOutput{Messages: messages}, nil)

		mk := &market.Market{
			EventID:  148192,
			Name:     "OVER_UNDER_25",
			Side:     "BACK",
			Exchange: "betfair",
			ExchangeMarket: market.ExchangeMarket{
				ID: "1.28910191",
				Runners: []market.Runner{
					{
						ID:   472671,
						Name: "Over 2.5 Goals",
						Prices: []market.PriceSize{
							{
								Price: 1.95,
								Size:  156.91,
							},
						},
					},
				},
			},
			StatisticoOdds: []*market.StatisticoOdds{
				{
					Price:     1.56,
					Selection: "over",
				},
			},
			Timestamp: 1583971200,
		}

		ch := make(chan *market.Market, 10)

		err := r.Receive(ch)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err)
		}

		mt := <-ch

		assert.Equal(t, mk, mt)
	})

	t.Run("logs and returns error if error returned by SQS client", func(t *testing.T) {
		t.Helper()

		client := new(aws.MockSqsClient)
		logger, hook := test.NewNullLogger()

		r := aws.NewMarketReceiver(client, logger, "messages", 3600)

		input := mock.MatchedBy(func(i *sqs.ReceiveMessageInput) bool {
			assert.Equal(t, "messages", *i.QueueUrl)
			assert.Equal(t, int64(3600), *i.WaitTimeSeconds)
			return true
		})

		e := errors.New("error happened")

		client.On("ReceiveMessage", input).Return(&sqs.ReceiveMessageOutput{}, e)

		ch := make(chan *market.Market, 10)

		err := r.Receive(ch)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, e, err)
		assert.Equal(t, 0, len(ch))
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Unable to receive messages from queue \"messages\", error happened.", hook.LastEntry().Message)
	})

	t.Run("logs error if unable to parse message body in market struct", func(t *testing.T) {
		t.Helper()

		client := new(aws.MockSqsClient)
		logger, hook := test.NewNullLogger()

		r := aws.NewMarketReceiver(client, logger, "messages", 3600)

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

		ch := make(chan *market.Market, 10)

		err := r.Receive(ch)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err)
		}

		assert.Equal(t, 0, len(ch))
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Unable to marshal message into market struct, invalid character 'i' looking for beginning of value.", hook.LastEntry().Message)
	})
}

var messageBody = `
	{
		"eventId": 148192,
		"name": "OVER_UNDER_25",
		"side": "BACK",
		"exchange": "betfair",
		"exchangeMarket": {
			"id": "1.28910191",
			"runners": [
				{
					"id": 472671,
					"name": "Over 2.5 Goals",
					"prices": [
						{
							"price": 1.95,
							"size": 156.91
						}
					]
				}
			]
		},
		"statisticoOdds": [
			{
				"price": 1.56,
				"selection": "over"
			}
		],
		"timestamp": 1583971200
	}
`
